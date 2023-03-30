package main

import (
	"fmt"
)

/* 2 arrays with 4 rows and 2 columns, to act as 2 pages of codebooks*/
var codebook1 = [4][2]int{{0b0000, 0b100000}, {0b0001, 0b110000}, {0b0010, 0b110011}, {0b0011, 0b111111}}
var codebook2 = [4][2]int{{0b0100, 0b101010}, {0b0101, 0b111110}, {0b0110, 0b100111}, {0b0111, 0b101011}}
var message = [4]int{0b0000, 0b0100, 0b0011, 0b0111} //4 row msg
var encrypted_message []int //empty array for encrypted message
var decrypted_message []int //empty array for decrypted message

func codebookLookup(lookupValue int)(returnLookupVal int) { //to lookup encrypted message
	var i, j int = 0, 0 //to get in the indices 
	for i = 0; i < 4; i++ {
	if codebook1[i][j] == lookupValue{ //check first elem of first page in codebook intially
      		j++ //access encrypted value
		returnLookupVal = codebook1[i][j]
	} else if codebook2[i][j] == lookupValue{ //then checks codebook 2
      		j++
		returnLookupVal = codebook2[i][j]
      		break
      		}
	}
	return returnLookupVal
}

func reverseCodebookLookup(lookupValue int)(returnLookupVal int) { //to lookup decrypted message
	var i, j int = 0, 1 //to get in the indices 
	for i = 0; i < 4; i++ {
	if codebook1[i][j] == lookupValue{
      		j--
		returnLookupVal = codebook1[i][j]
	} else if codebook2[i][j] == lookupValue{ //codebook2 check
      		j--
		returnLookupVal = codebook2[i][j]
      		break
      		}		
	}
	return returnLookupVal
}

func main() {
	var x, i int = 0, 0 //int vars to loop through codelookup
  	var y, z int = 0, 0 //int vars to loop through reverseCodeLookup
	var lookupValue int = 0 //used for intial lookup
  	var returnLookupVal int = 0 //used to store looked up value
  	fmt.Println(message) //print original message
  
	for i = 0; i < 4; i++ {
      		lookupValue = message[x] //start with original message
      		returnLookupVal = codebookLookup(lookupValue)
      		encrypted_message = append(encrypted_message, returnLookupVal)
		x++
		fmt.Printf("The ciphered value of a is %b\n", returnLookupVal)
	}
	
  fmt.Println(encrypted_message) //print encrypted message
  
  for z = 0; z < 4; z++{
  	lookupValue = encrypted_message[y] //start with encrypted message
    	returnLookupVal = reverseCodebookLookup(lookupValue)
    	decrypted_message = append(decrypted_message, returnLookupVal)
	y++
	fmt.Printf("The deciphered value of a is %b\n", returnLookupVal)
  }
  fmt.Println(decrypted_message) //print decrypted message
}
