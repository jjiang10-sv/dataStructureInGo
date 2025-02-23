class Reflector:
	def __init__(self):
		self.left = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		self.right = "EJMZALYXVBWFCRQUONTSPIKHGD" #reflector A - https://en.wikipedia.org/wiki/Enigma_rotor_details

	def reflect(self, index):
		xter_at = self.right[index]
		return self.left.find(xter_at)

# r = Reflector()
# print(r.reflect(6))
