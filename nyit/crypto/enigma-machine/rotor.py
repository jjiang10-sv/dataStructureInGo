class Rotor:
    def __init__(self, rotor_wiring, rotor_notch):
        self.left = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
        self.right = rotor_wiring
        self.notch = rotor_notch

    def map_right_left(self, cipher_index):
        xter_at = self.right[cipher_index]

        index = self.left.find(xter_at)
        return index

    def map_left_right(self, cipher_index):
        xter_at = self.left[cipher_index]
        index = self.right.find(xter_at)
        return index
    
    def rotate(self):
        self.left = self.left[1:] + self.left[0]
        self.right = self.right[1:] + self.right[0]
    
    def show_state(self):
        return self.left[0]

    def display(self):
        print(self.left)
        print(self.right)

# left_rotor = Rotor("DMTWSILRUYQNKFEJCAZBPGXOHV", "Q")
# right_rotor = Rotor("HQZGPJTMOBLNCIFDYAWVEUSRKX", "R")

# right_rotor.rotate()
# nextindex = right_rotor.map_right_left(17)

# print(nextindex)