/*
Matthew Lyons and Avinash Sookram
ChaCha20 Go Implementation
*/

func main() {
	// Create a new buffered reader to read user input from the console.
	reader := bufio.NewReader(os.Stdin)

	// Get the user input and remove the newline character.
	fmt.Print("Enter desired plaintext: ")
	input, _ := reader.ReadString('\n')
	plaintext := strings.TrimSuffix(input, "\n")

	// Generate a random 256-bit key and nonce for ChaCha20 encryption.
	key := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Convert the key and nonce to hexadecimal strings for storage.
	keyHex := hex.EncodeToString(key)
	nonceHex := hex.EncodeToString(nonce)

	// Encrypt the plaintext using ChaCha20.
	ciphertext := encrypt(key, nonce, plaintext)

	// Print the encrypted ciphertext and the key and nonce.
	fmt.Printf("Encrypted text: %x\n", ciphertext)
	fmt.Printf("Key: %s\n", keyHex)
	fmt.Printf("Nonce: %s\n", nonceHex)

	// Decrypt the ciphertext using ChaCha20.
	decryptedText := decrypt(key, nonce, ciphertext)

	// Print the decrypted plaintext.
	fmt.Printf("Decrypted text: %s\n", decryptedText)
}

func encrypt(key, nonce []byte, plaintext string) []byte {
	// Convert the plaintext to a byte slice.
	plaintextBytes := []byte(plaintext)
	//fmt.Printf("%s\n", plaintext)
	// Create a new ChaCha20 cipher with the given key and nonce.
	block, err := chacha20.NewUnauthenticatedCipher(key, nonce)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Encrypt the plaintext using the ChaCha20 cipher.
	ciphertext := make([]byte, len(plaintextBytes))
	block.XORKeyStream(ciphertext, plaintextBytes)

	return ciphertext
}

func decrypt(key, nonce, ciphertext []byte) string {
	// Create a new ChaCha20 cipher with the given key and nonce.
	block, err := chacha20.NewUnauthenticatedCipher(key, nonce)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Decrypt the ciphertext using the ChaCha20 cipher.
	plaintext := make([]byte, len(ciphertext))
	block.XORKeyStream(plaintext, ciphertext)

	// Convert the plaintext byte slice to a string.
	decryptedText := string(plaintext)

	return decryptedText
}
