package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

// Ramdom bytes for encryption and decryption
var bytes = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

// This should be in an env file in production
const MySecret string = "abc&1*~#^2^#s0^=)^^7%b34"

//How to encode with Base64 for Binary 
func Encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

// Base64 decoding
func Decode(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}

// Encrypt method is to encrypt or hide any classified text
func Encrypt(text, MySecret string) (string, error) {
	//Create the AES block
	block, err := aes.NewCipher([]byte(MySecret))
	if err != nil {
		return "", err
	}
	//plaintext must be a byte format
	plainText := []byte(text)
	ofb := cipher.NewOFB(block, bytes) //NewOFB is the cipher built-in instance for output feedback mode
	cipherText := make([]byte, len(plainText))
	ofb.XORKeyStream(cipherText, plainText)
	return Encode(cipherText), nil
}

// Decrypt method is to extract back the encrypted text
func Decrypt(text, MySecret string) (string, error) {
	block, err := aes.NewCipher([]byte(MySecret))
	if err != nil {
		return "", err
	}
	cipherText := Decode(text)
	ofb := cipher.NewOFB(block, bytes) //NewOFB acts as both encryption and decryption based off the documentation
	plainText := make([]byte, len(cipherText))
	ofb.XORKeyStream(plainText, cipherText)
	return string(plainText), nil
}

func main() {
	fmt.Println("Enter the string to encrypt: ")
	var StringToEncrypt string
	fmt.Scanln(&StringToEncrypt)
	// To encrypt the StringToEncrypt
	encText, err := Encrypt(StringToEncrypt, MySecret)
	if err != nil {
		fmt.Println("error encrypting your classified text: ", err)
	}
	fmt.Println(encText)
	decText, err := Decrypt(encText, MySecret)
	if err != nil {
		fmt.Println("error decrypting your encrypted text: ", err)
	}
	fmt.Println(decText)
}

//Credits to the original source code posted in Canvas
//Following are references:
//https://pkg.go.dev/crypto/cipher#example-NewOFB - documentation
//https://go.dev/src/crypto/cipher/ofb.go - NewOFB's implemetnation via the in-depth documentation
