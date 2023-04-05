package main

import (
	"fmt"
)

/* 2 arrays with 4 rows and 2 columns, to act as 2 pages of codebooks*/
var codebook1 = [4][2]int{{0b00, 0b01}, {0b01, 0b11}, {0b10, 0b00}, {0b11, 0b10}}
var message = [4]int{0b00, 0b10, 0b01, 0b11} //4 row msg
var iv int = 0b10 //initialization vector
var key int = 0b11 //key for keystream, usually is random but hardcoded for testing sake
var encrypted_message []int // encrypted message array
var decrypted_message []int // decrypted message array
var finalDecrypt []int

func codebookLookup(xor int)(lookupValue int) { //to lookup encrypted message
	var i, j int = 0, 0 //to get in the indices 
	for i = 0; i < 4; i++ {
		if codebook1[i][j] == xor{ 
			j++ //access encrypted value
			lookupValue = codebook1[i][j]
			j = 0 //reset access variable
		} 
  	}
	return lookupValue
}

func reverseCodebookLookup(xor int)(lookupValue int) { //to lookup decrypted message
	var i, j int = 0, 1 //to get in the indices 
	for i = 0; i < 4; i++ {
		if codebook1[i][j] == xor{
			j--
			lookupValue = codebook1[i][j]
			j = 1
			break
		}
	}
	return lookupValue
}

func reverseArray(arr []int) []int{
	for i, j := 0, len(arr)-1; i<j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
		}
	return arr
}

func main() {
var xor int = 0
var keystream int = 0
var x, i int = 0, 0 //int vars to loop through codelookup
var y, z int = 0, 0 //int vars to loop through reverseCodeLookup
var lookupValue int = 0 //used for intial lookup

fmt.Println(message) //print original message
  
for i = 0; i < 4; i++ {
	if x == 0{
		keystream = iv ^ key //create keystream
		xor = message[x] ^ keystream //use keystream to XOR plaintext to create ciphertext
	} else{
		xor = message[x] ^ keystream
	}
	lookupValue = codebookLookup(xor)
	encrypted_message = append(encrypted_message, lookupValue)
	x++
	fmt.Printf("The ciphered value of a is %b\n", lookupValue)
	}
  
fmt.Println(encrypted_message) //print encrypted message
  
for z = 0; z < 4; z++{
	if z == 0{
		keystream = iv ^ key //create keystream
		xor = encrypted_message[y] ^ keystream //use keystream to XOR ciphertext to create plaintext
	} else{
		xor = encrypted_message[y] ^ keystream
	}
	lookupValue = reverseCodebookLookup(xor)
	decrypted_message = append(decrypted_message, lookupValue)
	y++
	fmt.Printf("The deciphered value of a is %b\n", lookupValue)
}
finalDecrypt = reverseArray(decrypted_message)
fmt.Println(finalDecrypt) //print decrypted message
}
