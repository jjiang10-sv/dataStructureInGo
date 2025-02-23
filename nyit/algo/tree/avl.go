package tree

import (
	"fmt"
)

// TreeNode represents a node in the AVL tree
type TreeNode struct {
	Value  int
	Left   *TreeNode
	Right  *TreeNode
	Height int
}

// GetHeight returns the height of a node
func GetHeight(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return node.Height
}

// GetBalanceFactor returns the balance factor of a node
func GetBalanceFactor(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return GetHeight(node.Left) - GetHeight(node.Right)
}

// RightRotate performs a right rotation
func RightRotate(y *TreeNode) *TreeNode {
	x := y.Left
	T2 := x.Right

	// Perform rotation
	x.Right = y
	y.Left = T2

	// Update heights
	y.Height = max(GetHeight(y.Left), GetHeight(y.Right)) + 1
	x.Height = max(GetHeight(x.Left), GetHeight(x.Right)) + 1

	return x
}

// LeftRotate performs a left rotation
func LeftRotate(x *TreeNode) *TreeNode {
	y := x.Right
	T2 := y.Left

	// Perform rotation
	y.Left = x
	x.Right = T2

	// Update heights
	x.Height = max(GetHeight(x.Left), GetHeight(x.Right)) + 1
	y.Height = max(GetHeight(y.Left), GetHeight(y.Right)) + 1

	return y
}

// Insert inserts a value into the AVL tree (Recursive)
func Insert(root *TreeNode, value int) *TreeNode {
	if root == nil {
		return &TreeNode{Value: value, Height: 1}
	}

	if value < root.Value {
		root.Left = Insert(root.Left, value)
	} else if value > root.Value {
		root.Right = Insert(root.Right, value)
	} else {
		// Duplicate values are not allowed in AVL Tree
		return root
	}

	// Update height
	root.Height = 1 + max(GetHeight(root.Left), GetHeight(root.Right))

	// Get balance factor
	balance := GetBalanceFactor(root)

	// Perform rotations if necessary
	if balance > 1 && value < root.Left.Value {
		return RightRotate(root) // Left-Left (LL) Case
	}
	if balance < -1 && value > root.Right.Value {
		return LeftRotate(root) // Right-Right (RR) Case
	}
	if balance > 1 && value > root.Left.Value {
		root.Left = LeftRotate(root.Left)
		return RightRotate(root) // Left-Right (LR) Case
	}
	if balance < -1 && value < root.Right.Value {
		root.Right = RightRotate(root.Right)
		return LeftRotate(root) // Right-Left (RL) Case
	}

	return root
}

// Recursive Inorder Traversal (Left → Root → Right)
func InorderRecursive(root *TreeNode) {
	if root == nil {
		return
	}
	InorderRecursive(root.Left)
	fmt.Print(root.Value, " ")
	InorderRecursive(root.Right)
}

// Iterative Inorder Traversal (Using Stack)
func InorderIterative(root *TreeNode) {
	stack := []*TreeNode{}
	current := root

	for current != nil || len(stack) > 0 {
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}

		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Print(current.Value, " ")

		current = current.Right
	}
}

// Recursive Preorder Traversal (Root → Left → Right)
func PreorderRecursive(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Print(root.Value, " ")
	PreorderRecursive(root.Left)
	PreorderRecursive(root.Right)
}

// Iterative Preorder Traversal (Using Stack)
func PreorderIterative(root *TreeNode) {
	if root == nil {
		return
	}

	stack := []*TreeNode{root}

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Print(current.Value, " ")

		if current.Right != nil {
			stack = append(stack, current.Right)
		}
		if current.Left != nil {
			stack = append(stack, current.Left)
		}
	}
}

// Helper function to get the maximum of two numbers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var root *TreeNode

	// Insert values into the AVL tree
	values := []int{10, 20, 30, 40, 50, 25}
	for _, v := range values {
		root = Insert(root, v)
	}

	fmt.Print("Inorder Traversal (Recursive): ")
	InorderRecursive(root)
	fmt.Println()

	fmt.Print("Inorder Traversal (Iterative): ")
	InorderIterative(root)
	fmt.Println()

	fmt.Print("Preorder Traversal (Recursive): ")
	PreorderRecursive(root)
	fmt.Println()

	fmt.Print("Preorder Traversal (Iterative): ")
	PreorderIterative(root)
	fmt.Println()
}

// // GetHeight returns the height of a node
// func GetHeight(node *TreeNode) int {
// 	if node == nil {
// 		return 0
// 	}
// 	return node.Height
// }

// // GetBalanceFactor returns the balance factor of a node
// func GetBalanceFactor(node *TreeNode) int {
// 	if node == nil {
// 		return 0
// 	}
// 	return GetHeight(node.Left) - GetHeight(node.Right)
// }

// // RightRotate performs a right rotation
// func RightRotate(y *TreeNode) *TreeNode {
// 	x := y.Left
// 	T2 := x.Right

// 	x.Right = y
// 	y.Left = T2

// 	y.Height = max(GetHeight(y.Left), GetHeight(y.Right)) + 1
// 	x.Height = max(GetHeight(x.Left), GetHeight(x.Right)) + 1

// 	return x
// }

// // LeftRotate performs a left rotation
// func LeftRotate(x *TreeNode) *TreeNode {
// 	y := x.Right
// 	T2 := y.Left

// 	y.Left = x
// 	x.Right = T2

// 	x.Height = max(GetHeight(x.Left), GetHeight(x.Right)) + 1
// 	y.Height = max(GetHeight(y.Left), GetHeight(y.Right)) + 1

// 	return y
// }

// InsertIterative inserts a value into the AVL tree using iteration
func InsertIterative(root *TreeNode, value int) *TreeNode {
	if root == nil {
		return &TreeNode{Value: value, Height: 1}
	}

	stack := []*TreeNode{}
	node := root
	var parent *TreeNode

	// Traverse the tree iteratively to find the correct insertion point
	for node != nil {
		stack = append(stack, node)
		parent = node

		if value < node.Value {
			node = node.Left
		} else if value > node.Value {
			node = node.Right
		} else {
			// Duplicate values are not allowed in AVL Tree
			return root
		}
	}

	// Insert the new node
	newNode := &TreeNode{Value: value, Height: 1}
	if value < parent.Value {
		parent.Left = newNode
	} else {
		parent.Right = newNode
	}

	// Update heights and balance the tree
	for len(stack) > 0 {
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1] // Pop

		// Update height
		node.Height = 1 + max(GetHeight(node.Left), GetHeight(node.Right))

		// Get balance factor
		balance := GetBalanceFactor(node)

		// Perform rotations if necessary
		if balance > 1 && value < node.Left.Value {
			if len(stack) == 0 {
				return RightRotate(node) // Left-Left (LL) Case
			}
			parent := stack[len(stack)-1]
			if parent.Left == node {
				parent.Left = RightRotate(node)
			} else {
				parent.Right = RightRotate(node)
			}
		} else if balance < -1 && value > node.Right.Value {
			if len(stack) == 0 {
				return LeftRotate(node) // Right-Right (RR) Case
			}
			parent := stack[len(stack)-1]
			if parent.Left == node {
				parent.Left = LeftRotate(node)
			} else {
				parent.Right = LeftRotate(node)
			}
		} else if balance > 1 && value > node.Left.Value {
			node.Left = LeftRotate(node.Left)
			if len(stack) == 0 {
				return RightRotate(node) // Left-Right (LR) Case
			}
			parent := stack[len(stack)-1]
			if parent.Left == node {
				parent.Left = RightRotate(node)
			} else {
				parent.Right = RightRotate(node)
			}
		} else if balance < -1 && value < node.Right.Value {
			node.Right = RightRotate(node.Right)
			if len(stack) == 0 {
				return LeftRotate(node) // Right-Left (RL) Case
			}
			parent := stack[len(stack)-1]
			if parent.Left == node {
				parent.Left = LeftRotate(node)
			} else {
				parent.Right = LeftRotate(node)
			}
		}
	}

	return root
}

// // Inorder Traversal to print sorted values
// func InorderRecursive(root *TreeNode) {
// 	if root == nil {
// 		return
// 	}
// 	InorderRecursive(root.Left)
// 	fmt.Print(root.Value, " ")
// 	InorderRecursive(root.Right)
// }

// // Helper function to get the maximum of two numbers
// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// func main() {
// 	var root *TreeNode

// 	// Insert values into the AVL tree using iteration
// 	values := []int{10, 20, 30, 40, 50, 25}
// 	for _, v := range values {
// 		root = InsertIterative(root, v)
// 	}

// 	fmt.Print("Inorder Traversal (Recursive): ")
// 	InorderRecursive(root)
// 	fmt.Println()
// }
