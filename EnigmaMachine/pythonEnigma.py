from string import ascii_lowercase
import random
import string

alphabet = []
plugboard = []
#Fill alphabet and generate random plugboard configuration
for c in ascii_lowercase:
    alphabet.append(c)
    plugboard.append(c)
random.shuffle(plugboard)

print("Alphabet: ",alphabet,"\n\nPlugboard:",plugboard)

#Rotate encryption
def rotate(list):
    end = list[-1]
    del list[-1]
    list.insert(0, end)

#unroatate for encryption
def unRotate(list):
    start = list[0]
    del list[0]
    list.append(start)

#Xor two strings together
def xorStrings(s1, s2):
    b1 = bytes.fromhex(s1) if all(c in string.hexdigits for c in s1) else s1.encode()
    b2 = bytes.fromhex(s2) if all(c in string.hexdigits for c in s2) else s2.encode()
    result = bytearray(x ^ y for x, y in zip(b1, b2))
    return result.hex() if all(c in string.hexdigits for c in s1) else result.decode()

#Undo the Xor on two strings
def unxorStrings(s1, s2):
    key = xorStrings(s1, s2)
    result = xorStrings(s2, key)
    return result

#Generate a key to Xor message with
def keyGenerator(length):
    letters = string.ascii_letters
    return ''.join(random.choice(letters) for _ in range(length))

class rotor:
    #set rotor and fill with alphabet, starting letter determined by start argument
    def __init__(self,start):
        self.rotor = [alphabet[start]]
        for i in range (start+1,26):
            self.rotor.append(alphabet[i])
        if len(self.rotor) < 26:
            for i in range(0,start):
                self.rotor.append(alphabet[i])

    #Rotates the rotor by one
    def rotate(self):
        end = self.rotor[-1]
        del self.rotor[-1]
        self.rotor.insert(0,end)

    #Unrotates the rotor by one
    def unRotate(self):
        start = self.rotor[0]
        del self.rotor[0]
        self.rotor.append(start)

    #Prints a rotor
    def printer(self):
        print(self.rotor)

    #Encodes through rotor
    def encode(self,data,start):
        encryptedString = ""
        for i in data:
            if i in start:
                index = start.index(i)
                encryptedString = encryptedString + self.rotor[index]
                self.rotate()
        #print("\nEncrypted Message: ", encryptedString)
        return encryptedString

    #decodes through rotor
    def decode(self,data,end):
        decryptedString = ""
        for i in data:
            if i in self.rotor:
                index = self.rotor.index(i)
                decryptedString = decryptedString + end[index]
                unRotate(end)
        #print("\nDecrypted Message: ", decryptedString)
        return decryptedString

#Creates three rotors at different positions
first = rotor(14)
second = rotor(3)
third = rotor(4)
print("\n\n")

#user input
input = input("Enter String: ")
input = input.lower()

#Encrypt from initial alphabet through plugboard
def encrypt(input,start,end):
    encryptedString = ""
    for i in input:
        if i in start:
            index = start.index(i)
            encryptedString = encryptedString + end[index]
    #print("\nEncrypted Message: ", encryptedString)
    return encryptedString

#Decrypt from plugboard back to alphabet
def decrypt(encryptedString,start,end):
    decryptedString = ""
    for i in encryptedString:
        if i in start:
            index = start.index(i)
            decryptedString = decryptedString + end[index]
    #print("\nDecrypted Message: ", decryptedString)
    return decryptedString

#Start of encryption from alphabet to plugboard
secret = encrypt(input,alphabet,plugboard)

#First pass through rotors
Im = first.encode(secret,plugboard)
am = second.encode(Im,first.__getattribute__('rotor'))
output = third.encode(am,second.__getattribute__('rotor'))

#Pass back through rotors (reflector)
maybe = second.encode(output,third.__getattribute__('rotor'))
this = first.encode(maybe,second.__getattribute__('rotor'))
fullyEncodedOutput = encrypt(this,first.__getattribute__('rotor'),plugboard)


#creates a random key and then performs XOR to further encrypt message
key = keyGenerator(len(input))
encryptedMessagefull = fullyEncodedOutput
#Messaged fully encrypted, no longer appearing as a string of characters
encrypted = xorStrings(encryptedMessagefull, key)
hidden = b'encrypted'
print("\nHidden Encrypted Message: ",hidden.hex())
decrypted = unxorStrings(encryptedMessagefull, encrypted)

#Starts decryption through rotors
step1 = decrypt(decrypted,plugboard,first.__getattribute__('rotor'))
step2 = first.decode(step1,second.__getattribute__('rotor'))
step3 = second.decode(step2,third.__getattribute__('rotor'))

#Continues back through rotors again(reflector)
maybe = third.decode(step3,second.__getattribute__('rotor'))
this = second.decode(maybe,first.__getattribute__('rotor'))
works = first.decode(this,plugboard)

#Finishes decryption from plugboard back to alphabet
final = decrypt(works,plugboard,alphabet)

print("")
print("Decrypted Message: ",final)
print("")






