package main 

import (
	"encoding/hex"
	"fmt"
	"math/rand"
)

var alphabet []string  //used for standard alphabet
var plugboard []string //use for the plugboard implementation of random alphabet

func xorStrings(s1 string, s2 string) string {
	b1, err := hex.DecodeString(s1) //converts string s1 to hexadecimal
	if err != nil {
		b1 = []byte(s1) //checks if there's no errors, continues
	}
	b2, err := hex.DecodeString(s2) //converts string s2 to hexadecimal
	if err != nil {
		b2 = []byte(s2) //checks if there's no errors, continues
	}
	result := make([]byte, len(b1)) //allocates memory for the slices
	for i := range b1 {
		result[i] = b1[i] ^ b2[i] //performs xor of the two bytes in indices ranging from lengths of one of the bytemaps
	}
	return string(result) //return xor in string
}

func unxorStrings(s1 string, s2 string) string {
	key := xorStrings(s1, s2)     //xor the s1 and s2 string
	result := xorStrings(s2, key) //xor the s2 string and the key
	return result
}

func keyGenerator(length int) string {
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, length) //makes memory based on inputted length
	for i := range b {        //creates random integers based on range of length of valid letters
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func rotate(list []string) {
	end := list[len(list)-1]     //gets end of list
	list = append(list[0:], end) //appends end of list to the end... again..
}

func unRotate(list []string) {
	start := list[0]          //get beginning index
	copy(list, list[1:])      //saves list as its being sliced up
	list[len(list)-1] = start //add the starting index to the end of list
}

type rotor struct {
	rotor []string //rotor is a string type
}

func NewRotor(start int) *rotor {
	r := &rotor{}                                  //initialize new rotor
	r.rotor = append(r.rotor, alphabet[start:]...) //appends alphabets to rotor
	r.rotor = append(r.rotor, alphabet[:start]...)
	return r
}

func (r *rotor) rotate() {
	end := r.rotor[len(r.rotor)-1]                               //retrieve end of rotor
	r.rotor = append([]string{end}, r.rotor[:len(r.rotor)-1]...) //add end of rotor to the beginning, rotate like a wheel
}

func (r *rotor) unRotate() {
	start := r.rotor[0] //similar to rotate but with the starting index and adding that start index to the end
	r.rotor = append(r.rotor[1:], start)
}

func (r *rotor) encode(data string, start []string) string {
	var encryptedString string //create string var
	for _, c := range data {   //looop through data string
		if idx := findIndex(start, string(c)); idx >= 0 { //get index based on start string
			encryptedString += r.rotor[idx] //add index of the rotor to the encrypted string
			r.rotate()
		}
	}
	return encryptedString
}

func (r *rotor) decode(data string, end []string) string {
	var decryptedString string //create string var
	for _, c := range data {   //loop through data string
		if idx := findIndex(r.rotor, string(c)); idx >= 0 { //finds index of rotor
			decryptedString = decryptedString + end[idx] //uses index of rotor to find end string and add to decrypted string
			unRotate(end)
		}
	}
	return decryptedString
}

func encrypt(input string, start []string, end []string) string {
	var encryptedString string //create string var
	for _, c := range input {  //loop through input
		if idx := findIndex(start, string(c)); idx >= 0 { //find index of the start string
			encryptedString += end[idx] //use start index to index end string, add to encrypted string
		}
	}
	return encryptedString
}

func decrypt(data string, start []string, end []string) string {
	var decryptedString string //create string var
	for _, c := range data {   //loop through data string
		if idx := findIndex(start, string(c)); idx >= 0 { //after getting start index
			decryptedString += end[idx] //use start index to index end string and add to decrypted string
		}
	}
	return decryptedString
}

func findIndex(s []string, val string) int {
	for i := range s { //loop through range of s
		if s[i] == val { //if index at s = val, return that index
			return i
		}
	}
	return -1 //if not, return -1
}

func main() {
	var encryptedMessagefull string //create string for placeholder encryptedmessage

	for _, c := range "abcdefghijklmnopqrstuvwxyz" {
		alphabet = append(alphabet, string(c))   //add to alphabet list
		plugboard = append(plugboard, string(c)) //add to plugboard list
	}

	rand.Shuffle(len(plugboard), func(i, j int) {
		plugboard[i], plugboard[j] = plugboard[j], plugboard[i] //randomize the order of plugboard
	})

	fmt.Println("Alphabet:", alphabet, "\n\nPlugboard:", plugboard) //print alphabet and plugboard

	first := NewRotor(14) //initialize the three rotors
	second := NewRotor(3)
	third := NewRotor(4)

	input := "rawr"
	fmt.Printf("\nInitial input message: %s\n", input) //print input message

	secret := encrypt(input, alphabet, plugboard) //initial encryption with input, alphabet, and plugboard

	Im := first.encode(secret, plugboard)    //encode first rotor with secret and plugboard
	am := second.encode(Im, first.rotor)     //encode second rotor with first encoding output and first rotor
	output := third.encode(am, second.rotor) //encode third rotor with second encoding output and second rotor

	//pass back through rotors, reflector logic

	maybe := second.encode(output, third.rotor)                 //go back to second rotor
	this := first.encode(maybe, second.rotor)                   //go back to first rotor
	fullyEncodedOutput := encrypt(this, first.rotor, plugboard) //encrypt back through plugboard

	key := keyGenerator(len(input))                                            //generate a key using length of the input
	encryptedMessagefull = fullyEncodedOutput                                  //reassign the encrypted output to another variable
	encrypted := xorStrings(encryptedMessagefull, key)                         //perform xor on the encrypted message and the key
	hidden := []byte("encrypted")                                              //create bytemap for encrypted message
	fmt.Printf("\nHidden Encrypted Message: %s\n", hex.EncodeToString(hidden)) //convert hidden to hexadecimal
	decrypted := unxorStrings(encryptedMessagefull, encrypted)                 //to decrypt, unxor the encrypted message and the xor'd result of encrypted msg + key

	step1 := decrypt(decrypted, plugboard, first.rotor) //decrypt from plugboard to first rotor
	step2 := first.decode(step1, second.rotor)          // decode from first to second rotor
	step3 := second.decode(step2, third.rotor)          // decode from second to third rotor

	//Pass back through rotors (reflector)

	maybe = third.decode(step3, second.rotor) // decode from third to second rotor
	this = second.decode(maybe, first.rotor)  // decode from second to first rotor
	works := first.decode(this, plugboard)    // decode from first rotor to plugboard

	wooo := decrypt(works, plugboard, alphabet) // decrypt from plugboard back to alphabet
	fmt.Println("\nDecrypted Message:", wooo)   //prints decrypted response
}

/*
Resources that assisted us in programming the added quirks:
https://pkg.go.dev/encoding/hex#EncodeToString - hexadecimal documentation in GoLang
https://pkg.go.dev/math/rand@go1.20#Seed - Random documentation for Go (like random numbers etc)
https://www.geeksforgeeks.org/how-to-copy-one-slice-into-another-slice-in-golang/# - Assisted in saving lists and splicing properly
https://askgolang.com/how-to-create-a-byte-array-in-golang/ - make() function and byte array assistance
*/
