package main

import (
    "fmt"
)

func caesarEncrypt(text string, shift int) string {
    result := ""
    // iterate over each character in the text
    for _, char := range text {
        // check if the character is a letter
        if char >= 'A' && char <= 'Z' {
            // calculate the new ASCII code by applying the shift
            newChar := (char-'A'+rune(shift))%26 + 'A'
            // append the new character to the result
            result += string(newChar)
        } else if char >= 'a' && char <= 'z' {
            // calculate the new ASCII code by applying the shift
            newChar := (char-'a'+rune(shift))%26 + 'a'
            // append the new character to the result
            result += string(newChar)
        } else {
            // if the character is not a letter, just append it to the result
            result += string(char)
        }
    }
    return result
}

func caesarDecrypt(text string) {
    // try all possible shifts
    for shift := 0; shift < 26; shift++ {
        result := ""
        // iterate over each character in the text
        for _, char := range text {
            // check if the character is a letter
            if char >= 'A' && char <= 'Z' {
                // calculate the new ASCII code by applying the shift
                newChar := (char-'A'-rune(shift)+26)%26 + 'A'
                // append the new character to the result
                result += string(newChar)
            } else if char >= 'a' && char <= 'z' {
                // calculate the new ASCII code by applying the shift
                newChar := (char-'a'-rune(shift)+26)%26 + 'a'
                // append the new character to the result
                result += string(newChar)
            } else {
                // if the character is not a letter, just append it to the result
                result += string(char)
            }
        }
        // print the decrypted text for the current shift
        fmt.Printf("Shift %d: %s\n", shift, result)
    }
}

func main() {
    // Example usage
    text := "HELLO WORLD"
    shift := 3
    encryptedText := caesarEncrypt(text, shift)
    fmt.Println(encryptedText)

    caesarDecrypt(encryptedText)
}
