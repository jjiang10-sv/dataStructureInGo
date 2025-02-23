import hashlib

class MerkleTree:
    def __init__(self, transactions):
        self.transactions = transactions
        self.tree = self.build_merkle_tree()

    def hash_function(self, data):
        return hashlib.sha256(data.encode()).hexdigest()

    def build_merkle_tree(self):
        nodes = [self.hash_function(tx) for tx in self.transactions]

        while len(nodes) > 1:
            if len(nodes) % 2 != 0:  # If odd, duplicate last node
                nodes.append(nodes[-1])

            temp = []
            for i in range(0, len(nodes), 2):
                parent_hash = self.hash_function(nodes[i] + nodes[i+1])
                temp.append(parent_hash)
            nodes = temp

        return nodes[0]  # Merkle Root

# Example transactions
transactions = ["tx1", "tx2", "tx3", "tx4", "tx5"]
merkle_tree = MerkleTree(transactions)

print("Merkle Root:", merkle_tree.tree)
