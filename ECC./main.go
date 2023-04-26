package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// Generate a private/public key pair using an elliptic curve
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}

	// Get user input using bufio library
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter message: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	input = strings.TrimSpace(input)

	// AEC-GCM is used to encrypt the user input (Galois/Counter Mode)
	block, err := aes.NewCipher(privateKey.D.Bytes())
	if err != nil {
		panic(err)
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}
	//Produce nonce
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err)
	}

	//Perform encryption using gcm.Seal()
	ciphertext := gcm.Seal(nil, nonce, []byte(input), nil)
	fmt.Printf("Encrypted Message: %s\n", ciphertext)

	// Decrypt the ciphertext back to plaintext
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err)
	}

	// Print decrypted message
	fmt.Printf("Decrypted Message: %s\n", string(plaintext))
}
