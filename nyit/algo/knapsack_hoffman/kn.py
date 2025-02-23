def fractinal_knapsack(values, weights, W):
    n = len(weights)
    ratios = [(values[i]/weights[i], values[i], weights[i]) for i in range(n)]

    # sort the items based in decrasing order
    ratios.sort(key=lambda x: x[0], reverse=True)

    total_val = 0
    current_weight = 0
    for ratio, val, weight in ratios:
        if current_weight + weight <= W:
            total_val += val
            current_weight += weight
        else:
            fraction = (W-current_weight)/weight
            total_val += fraction*val
            break
    return total_val
import heapq

class node:
    def __init__(self, name, frequency):
        self.name = name
        self.frequency = frequency
        self.left = None
        self.right = None

def hoff_man(frequencies):
    nodes = []
    for name, frequency in frequencies.items():
        nodes.append(node(name, frequency))
    heapq.heapify(nodes)
    while len(nodes) > 1:
        left = heapq.heappop(nodes)
        right = heapq.heappop(nodes)
        new_node = node(None, left.frequency + right.frequency)
        new_node.left = left
        new_node.right = right
        heapq.heappush(nodes, new_node)
    return heapq.heappop(nodes)

if __name__ == "__main__":
    values = [40,50,20]
    weights = [2,5,4]
    W = 6
    print(fractinal_knapsack(values=values,weights=weights,W=W))
    frequencies = {"a": 5, "b": 9, "c": 12, "d": 13, "e": 16, "f": 45}
    
    test_str = "abc"
            
    # Build Huffman tree
    root = hoff_man(frequencies)
    
    # Generate codes by traversing tree
    codes = {}
    def generate_codes(node, code=""):
        if node.name:
            codes[node.name] = code
            return
        generate_codes(node.left, code + "0")
        generate_codes(node.right, code + "1")
    
    generate_codes(root)
    
    # Encode string
    encoded = ""
    for char in test_str:
        encoded += codes[char]
        
    print(f"Original string: {test_str}")
    print(f"Huffman codes: {codes}")
    print(f"Encoded string: {encoded}")
