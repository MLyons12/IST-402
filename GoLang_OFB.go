package main

import (
	"fmt"
)

/* 2 arrays with 4 rows and 2 columns, to act as 2 pages of codebooks*/
var codebook1 = [4][2]int{{0b0000, 0b100000}, {0b0001, 0b110000}, {0b0010, 0b110011}, {0b0011, 0b111111}}
var codebook2 = [4][2]int{{0b0100, 0b101010}, {0b0101, 0b111110}, {0b0110, 0b100111}, {0b0111, 0b101011}}
var message = [4]int{0b0000, 0b0100, 0b0011, 0b0111} //4 row msg
var iv int = 0b10 //initialization vector
var encrypted_message []int // encrypted message array
var decrypted_message []int // decrypted message array

func codebookLookup(xor int)(lookupValue int) { //to lookup encrypted message
	var i, j int = 0, 0 //to get in the indices 
	for i = 0; i < 4; i++ {
	if codebook1[i][j] == xor{ //check first elem of first page in codebook intially
      		j++ //access encrypted value
		lookupValue = codebook1[i][j]
	} else if codebook2[i][j] == xor{ //then checks codebook 2
      		j++
		lookupValue = codebook2[i][j]
      		break
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
	} else if codebook2[i][j] == xor{ //codebook2 check
      		j--
		lookupValue = codebook2[i][j]
      		break
      		}
	}
	return lookupValue
}

func main() {
	var xor int = 0
 	var x, i int = 0, 0 //int vars to loop through codelookup
 	var y, z int = 0, 0 //int vars to loop through reverseCodeLookup
	var lookupValue int = 0 //used for intial lookup

  	fmt.Println(message) //print original message
  
  	for i = 0; i < 4; i++ {
  		if x == 0{
        		lookupValue = message[x] ^ iv
      		} else{
        		xor = message[x] ^ lookupValue
      		}
      		lookupValue = codebookLookup(xor)
      		encrypted_message = append(encrypted_message, lookupValue)
		x++
		fmt.Printf("The ciphered value of a is %b\n", lookupValue)
	}
  
  	fmt.Println(encrypted_message) //print encrypted message
  
  	for z = 0; z < 4; z++{
    		if z == 3{
      			xor = encrypted_message[y] ^ iv
    		} else{
      			xor = encrypted_message[y] ^ lookupValue
    		}
    		lookupValue = reverseCodebookLookup(xor)
    		decrypted_message = append(decrypted_message, lookupValue)
	  	y++
	  	fmt.Printf("The deciphered value of a is %b\n", lookupValue)
  	}
  	fmt.Println(decrypted_message) //print decrypted message
}
