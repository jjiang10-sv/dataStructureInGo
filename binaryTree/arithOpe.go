// To implement arithmetic operations using a tree in Go, you can use a binary expression tree. The idea is:
// 	1.	Each internal node represents an operator (+, -, *, /).
// 	2.	Each leaf node represents an operand (numbers).
// 	3.	In-order traversal (Left → Root → Right) reconstructs the expression in correct order.
// 	4.	Post-order traversal (Left → Right → Root) is used to evaluate the expression.

// Implementation in Go

package binaryTree

import (
	"fmt"
	"strconv"
	"strings"
)

// NodeOpe represents a node in the expression tree
type NodeOpe struct {
	Value string
	Left  *NodeOpe
	Right *NodeOpe
}

// InOrderTraversal prints the tree in infix notation
func InOrderTraversal(root *NodeOpe) {
	if root == nil {
		return
	}
	if root.Left != nil {
		fmt.Print("(")
	}
	InOrderTraversal(root.Left)
	fmt.Print(root.Value)
	InOrderTraversal(root.Right)
	if root.Right != nil {
		fmt.Print(")")
	}
}

// Evaluate recursively computes the value of the expression tree
func Evaluate(root *NodeOpe) float64 {
	if root == nil {
		return 0
	}

	// If it's a leaf node, return the numeric value
	if root.Left == nil && root.Right == nil {
		val, _ := strconv.ParseFloat(root.Value, 64)
		return val
	}

	// Recursively evaluate left and right subtrees
	leftVal := Evaluate(root.Left)
	rightVal := Evaluate(root.Right)

	// Apply the operator at the current node
	switch root.Value {
	case "+":
		return leftVal + rightVal
	case "-":
		return leftVal - rightVal
	case "*":
		return leftVal * rightVal
	case "/":
		if rightVal == 0 {
			panic("division by zero")
		}
		return leftVal / rightVal
	}

	return 0
}

// Example: (3 + 5) * (10 - 2)
func mainOp() {
	// Constructing the tree for "(3 + 5) * (10 - 2)"
	root := &NodeOpe{"*",
		&NodeOpe{"+",
			&NodeOpe{"3", nil, nil},
			&NodeOpe{"5", nil, nil},
		},
		&NodeOpe{"-",
			&NodeOpe{"10", nil, nil},
			&NodeOpe{"2", nil, nil},
		},
	}

	// Print the expression in infix notation
	fmt.Print("Expression: ")
	InOrderTraversal(root)
	fmt.Println()

	// Evaluate the expression
	result := Evaluate(root)
	fmt.Println("Result:", result)
}

// Stack data structure
type Stack struct {
	items []*NodeOpe
}

// Push adds a node to the stack
func (s *Stack) Push(n *NodeOpe) {
	s.items = append(s.items, n)
}

// Pop removes and returns the top node from the stack
func (s *Stack) Pop() *NodeOpe {
	if len(s.items) == 0 {
		return nil
	}
	n := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return n
}

// BuildExpressionTree constructs an expression tree from a postfix expression
func BuildExpressionTree(postfix string) *NodeOpe {
	stack := &Stack{}
	tokens := strings.Split(postfix, " ")

	for _, token := range tokens {
		if isOperator(token) {
			// Pop two nodes from stack
			right := stack.Pop()
			left := stack.Pop()

			// Create new tree node with operator
			node := &NodeOpe{Value: token, Left: left, Right: right}

			// Push new subtree back to stack
			stack.Push(node)
		} else {
			// Push operand as a tree node
			stack.Push(&NodeOpe{Value: token})
		}
	}

	// Final tree root is at the top of the stack
	return stack.Pop()
}

// isOperator checks if a token is an operator
func isOperator(token string) bool {
	return token == "+" || token == "-" || token == "*" || token == "/"
}

// InOrderTraversalSta prints the tree in infix notation
func InOrderTraversalSta(root *NodeOpe) {
	if root == nil {
		return
	}
	if isOperator(root.Value) {
		fmt.Print("(")
	}
	InOrderTraversalSta(root.Left)
	fmt.Print(root.Value)
	InOrderTraversalSta(root.Right)
	if isOperator(root.Value) {
		fmt.Print(")")
	}
}

// EvaluateSta recursively computes the value of the expression tree
func EvaluateSta(root *NodeOpe) float64 {
	if root == nil {
		return 0
	}

	// If it's a leaf node, return the numeric value
	if root.Left == nil && root.Right == nil {
		val, _ := strconv.ParseFloat(root.Value, 64)
		return val
	}

	// Recursively evaluate left and right subtrees
	leftVal := EvaluateSta(root.Left)
	rightVal := EvaluateSta(root.Right)

	// Apply the operator at the current node
	switch root.Value {
	case "+":
		return leftVal + rightVal
	case "-":
		return leftVal - rightVal
	case "*":
		return leftVal * rightVal
	case "/":
		if rightVal == 0 {
			panic("division by zero")
		}
		return leftVal / rightVal
	}

	return 0
}

// Example: Postfix expression for "(3 + 5) * (10 - 2)" → "3 5 + 10 2 - *"
func mainSta() {
	postfix := "3 5 + 10 2 - *"
	root := BuildExpressionTree(postfix)

	// Print the expression in infix notation
	fmt.Print("Expression: ")
	InOrderTraversalSta(root)
	fmt.Println()

	// EvaluateSta the expression
	result := EvaluateSta(root)
	fmt.Println("Result:", result)
}

// EvaluatePostfix evaluates a postfix expression using a stack
func EvaluatePostfix(expression string) float64 {
	stack := []float64{}
	tokens := strings.Split(expression, " ")

	for _, token := range tokens {
		if isOperator(token) {
			// Pop the top two elements
			if len(stack) < 2 {
				panic("Invalid postfix expression")
			}
			b := stack[len(stack)-1] // Second operand (right)
			stack = stack[:len(stack)-1]
			a := stack[len(stack)-1] // First operand (left)
			stack = stack[:len(stack)-1]

			// Perform the operation and push the result
			var result float64
			switch token {
			case "+":
				result = a + b
			case "-":
				result = a - b
			case "*":
				result = a * b
			case "/":
				if b == 0 {
					panic("division by zero")
				}
				result = a / b
			}

			stack = append(stack, result)
		} else {
			// Convert number and push to stack
			num, err := strconv.ParseFloat(token, 64)
			if err != nil {
				panic("Invalid token in expression")
			}
			stack = append(stack, num)
		}
	}

	// The final result should be the only element in the stack
	if len(stack) != 1 {
		panic("Invalid postfix expression")
	}

	return stack[0]
}

func mainNoTree() {
	// Example: Postfix expression for "(3 + 5) * (10 - 2)" → "3 5 + 10 2 - *"
	postfix := "3 5 + 10 2 - *"
	result := EvaluatePostfix(postfix)
	fmt.Println("Result:", result) // Output: 64
}

// How It Works

// 	1.	Tree Structure (Example for (3 + 5) * (10 - 2))

//         *
//        / \
//       +   -
//      / \  / \
//     3   5 10 2

// 	2.	In-Order Traversal Output (Expression Reconstruction)

// (3+5)*(10-2)

// 	3.	Evaluation
// 	•	3 + 5 = 8
// 	•	10 - 2 = 8
// 	•	8 * 8 = 64

// Output

// Expression: (3+5)*(10-2)
// Result: 64

// This implementation supports basic arithmetic operations and can be extended for more complex expressions. 🚀
