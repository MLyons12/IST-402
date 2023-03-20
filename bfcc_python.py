def caesar_encrypt(text, shift):
    """
    Encrypts a given text using a Caesar Cipher encryption with the given shift.
    """
    result = ""
    # iterate over each character in the text
    for char in text:
        # check if the character is a letter
        if char.isalpha():
            # get the ASCII code of the character
            ascii_code = ord(char)
            # calculate the new ASCII code by applying the shift
            new_ascii_code = (ascii_code - 65 + shift) % 26 + 65
            # convert the new ASCII code back to a character and append it to the result
            result += chr(new_ascii_code)
        else:
            # if the character is not a letter, just append it to the result
            result += char
    return result


def caesar_decrypt(text):
    """
    Decrypts a given text encrypted with a Caesar Cipher by trying all possible shifts.
    """
    # try all possible shifts
    for shift in range(26):
        result = ""
        # iterate over each character in the text
        for char in text:
            # check if the character is a letter
            if char.isalpha():
                # get the ASCII code of the character
                ascii_code = ord(char)
                # calculate the new ASCII code by applying the shift
                new_ascii_code = (ascii_code - 65 - shift) % 26 + 65
                # convert the new ASCII code back to a character and append it to the result
                result += chr(new_ascii_code)
            else:
                # if the character is not a letter, just append it to the result
                result += char
        # print the decrypted text for the current shift
        print(f"Shift {shift}: {result}")


text = "HELLO WORLD"
shift = 3
encrypted_text = caesar_encrypt(text, shift)
print(encrypted_text)

caesar_decrypt(encrypted_text)