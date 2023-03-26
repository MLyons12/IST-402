#Matthew Lyons and Avinash Sookram
#CBC Python Implementation
#IST 402 Spring 2023

codebook = [[0b00, 0b01], [0b01, 0b10], [0b10, 0b11], [0b11, 0b00]]
message_str = input("Enter message ")
iv = 10

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
        xor = message[x] ^ iv
    else:
        xor = message[x] ^ lookup_value
    encrypted = codebook_lookup(xor)
    print(f"Cipher Value of {message[i]} is {bin(encrypted)}")
    lookup_value = codebook_lookup_reverse(xor)
    x -= 1
    decrypted_message.append(xor ^ iv)
    iv = codebook_lookup_reverse(xor)
decrypted_message.reverse()
decryption_Result = binary_to_string(''.join([str(bit) for bit in decrypted_message]))
print(f"Decrypted message is:", decryption_Result)



# Reference https://www.youtube.com/watch?v=oVCCXZfpu-w&ab_channel=AleksanderEssex
# https://www.youtube.com/watch?v=soatRmpccPk&ab_channel=CISSPrep
#https://www.programiz.com/python-programming/methods/built-in/ord
#https://www.w3schools.com/python/ref_string_format.asp
#https://www.scaler.com/topics/xor-in-python/
