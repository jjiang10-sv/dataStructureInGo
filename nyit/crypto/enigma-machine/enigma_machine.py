import argparse

from keyboard import Keyboard
from rotor import Rotor
from reflector import Reflector
from switch_board import SwitchBoard

class EnigmaMachine:
    def __init__(self, leftrotor, rightrotor, switchboard, reflector, keyboard):
        self.left_rotor = leftrotor
        self.right_rotor = rightrotor
        self.switchboard = switchboard
        self.reflector = reflector
        self.keyboard = keyboard

    def encrypt(self, message):
        uppercaseMsg = message.upper()
        ciphertext = ""
        for char in uppercaseMsg:
            if not char.isalpha():
                ciphertext += char
                continue
            initialChar = char

            charIndex = self.keyboard.map_right_left(char)
            charIndex = self.switchboard.map_right_left(charIndex)
            charIndex = self.right_rotor.map_right_left(charIndex)
            charIndex = self.left_rotor.map_right_left(charIndex)
            charIndex = self.reflector.reflect(charIndex)
            charIndex = self.left_rotor.map_left_right(charIndex)
            charIndex = self.right_rotor.map_left_right(charIndex)
            charIndex = self.switchboard.map_left_right(charIndex)
            char = self.keyboard.map_left_right(charIndex)
            ciphertext += char

            # print("left rotor before rotate state" + " " + self.left_rotor.right)
            # print("left rotor before rotate state" + " " + self.left_rotor.left)
            # print("right rotor before rotate state" + " " + self.right_rotor.right)
            # print("right rotor before rotate state" + " " + self.right_rotor.left)
            

            self.right_rotor.rotate()
            # why fix it as 25, should it compare the position and notch to see
            # whether they are equal
            if self.right_rotor.right[25] == self.right_rotor.notch:
                self.left_rotor.rotate()
            
            # print(initialChar + ": " + char)
            # print("left rotor before rotate state" + " " + self.left_rotor.right)
            # print("left rotor before rotate state" + " " + self.left_rotor.left)
            # print("right rotor before rotate state" + " " + self.right_rotor.right)
            # print("right rotor before rotate state" + " " + self.right_rotor.left)
            # print("\n") 
        return ciphertext
   
    def showState(self):
        print("outer cylinder"+ "     " + "inner cylinder")
        for char1, char2, char3, char4 in zip(self.right_rotor.left, self.left_rotor.right, self.left_rotor.left, self.reflector.right):
            print(char3 + ":" + char4 + "            " +char1 + ":" + char2)
        


rotor1 = Rotor("DMTWSILRUYQNKFEJCAZBPGXOHV", "Q")
rotor2 = Rotor("HQZGPJTMOBLNCIFDYAWVEUSRKX", "Z")
switchboard = SwitchBoard()
reflector = Reflector()
keyboard = Keyboard()

em = EnigmaMachine(rotor1, rotor2, switchboard, reflector, keyboard)

while(1):
    message = input("Enter your text to be encrypted (Type exit to exit the program): ")

    if message == "exit":
        break
    elif message == "showstate":
        em.showState()
        continue
    
    result = em.encrypt(message)
    print("The Encrypted Text is: " + result)

# python enigma_machine.py "asds"
# if __name__ == "__main__":
#     parser = argparse.ArgumentParser(description="Encrypt a message using the Enigma machine.")
#     parser.add_argument("message", type=str, help="The message to encrypt")
#     args = parser.parse_args()
    
#     main(args.message)  # Call the main function with the provided message