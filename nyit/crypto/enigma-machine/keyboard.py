class Keyboard:
        
    def map_right_left(self, letter):
        index = "ABCDEFGHIJKLMNOPQRSTUVWXYZ".find(letter)
        return index

    def map_left_right(self, index):
        letter = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"[index]
        return letter
    
