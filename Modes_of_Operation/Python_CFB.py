#Matthew Lyons and Avinash Sookram
#CBC Python Implementation
#IST 402 Spring 2023

codebook = [[0b00, 0b01], [0b01, 0b10], [0b10, 0b11], [0b11, 0b00]]
message_str = "hi"
iv = input("Enter IV: ")
iv = int(iv)

print("Initial Message: ",message_str)
def string_to_binary(s):
    #convert string to binary
    binary = ""
    for char in s:
        binary += format(ord(char), '08b')
    return binary

def binary_to_string(binary):
    #convert binary back to string
    section = [binary[i:i+8] for i in range(0, len(binary), 8)]
    string = ''.join([chr(int(chunk, 2)) for chunk in section])
    return string

message = [int(bit) for bit in string_to_binary(message_str)]

def codebook_lookup(xor):
    for row in codebook:
        if row[0] == xor:
            return row[1]
    return 0

def codebook_lookup_reverse(xor):
    for row in codebook:
        if row[1] == xor:
            return row[0]
    return 0

encrypted = []
x, lookup_value = len(message)-1, 0
decrypted_message = []
for i in range(len(message)):
    if x == len(message)-1:
        xor = message[x] ^ lookup_value
    else:
        xor = message[x] ^ iv
    encrypted = codebook_lookup(xor)
    print(f"Cipher Value of {message[i]} is {bin(encrypted)}")
    lookup_value = codebook_lookup_reverse(xor)
    x -= 1
    decrypted_message.append(xor ^ iv)
    iv = codebook_lookup_reverse(xor)
decrypted_message.reverse()
decryption_Result = binary_to_string(''.join([str(bit) for bit in decrypted_message]))
print(f"Decrypted message is:", decryption_Result)


# Reference The code provided on the assignment page
# https://www.educative.io/answers/what-is-cfb
# https://www.youtube.com/watch?v=3kxqgc4YIbE&ab_channel=ChiragBhalodia
#https://www.tutorialspoint.com/cryptography/block_cipher_modes_of_operation.htm
#https://www.educative.io/answers/how-to-convert-string-to-binary-in-python
