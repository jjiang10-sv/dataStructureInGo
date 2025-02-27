class SwitchBoard:
	def __init__(self):
		self.left = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		self.right = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	def swap(self, swaps):
		for swap in swaps:
			from_ = swap[0]
			to_ = swap[1]
			from_index = self.left.find(swap[0])
			to_index = self.left.find(swap[1])

			self.left = self.left[:from_index] + to_ + self.left[from_index+1:]
			self.left = self.left[:to_index] + from_ + self.left[to_index+1:]

	def map_right_left(self, cipher_index):
		xter_at = self.right[cipher_index]
		index = self.left.find(xter_at)
		return index

	def map_left_right(self, cipher_index):
		xter_at = self.left[cipher_index]
		index = self.right.find(xter_at)
		return index

	def display(self):
		return "SwitchBoard \nLeft=" + self.left + "\nRight=" + self.right + "\n"

# sb = SwitchBoard()
# sb.swap(["AB", "OZ"])
# print(sb.display())
# print(sb.map_right_left(0))
# print(sb.map_left_right(23))
# print(sb.display())