package main

import (
	"context"
	"fmt"
)

// definition of a bst node
type node struct {
	val   int
	left  *node
	right *node
}

// definition of a node
type btree struct {
	root *node
}

// allocating a new node
func newNode(val int) *node {
	return &node{val, nil, nil}
}

// insert nodes into a binary search tree
func insert(root *node, val int) *node {
	if root == nil {
		return newNode(val)
	}
	if val < root.val {
		root.left = insert(root.left, val)
	} else {
		root.right = insert(root.right, val)
	}
	return root
}

// inorder traversal algorithm
// Copies the elements of the bst to the array in sorted order
func inorderCopy(n *node, array []int, index *int) {
	if n != nil {
		inorderCopy(n.left, array, index)
		array[*index] = n.val
		*index++
		inorderCopy(n.right, array, index)
	}
}

func treesort(array []int, tree *btree) {
	// build the binary search tree
	for _, element := range array {
		tree.root = insert(tree.root, element)
	}
	index := 0
	// perform inorder traversal to get the elements in sorted order
	inorderCopy(tree.root, array, &index)
}

// // tester
// func gotest() {
// 	tree := &btree{nil}
// 	numbers := []int{5, 4, 3, 2, 1, -1, 0}
// 	fmt.Println("numbers : ", numbers)
// 	treesort(numbers, tree)
// 	fmt.Println("numbers : ", numbers)
// }

type bst struct {
	data  int
	left  *bst
	right *bst
}

func newBst(i int) *bst {
	return &bst{data: i}
}
func insertNode(i int, b *bst) *bst {
	if b == nil {
		return &bst{data: i}
	}
	if i < b.data {
		b.left = insertNode(i, b.left)
	} else {
		b.right = insertNode(i, b.right)
	}
	return b
}

func loadBstData(b *bst) []int {
	res := []int{}
	asdWalk(b, &res)
	//descWalk(b, &res)
	//breadthWalk(b, &res)
	return res
}

func asdWalk(b *bst, arr *[]int) {

	if b == nil {
		return
	}

	asdWalk(b.left, arr)
	*arr = append(*arr, b.data)
	asdWalk(b.right, arr)

}

func descWalk(b *bst, arr *[]int) {

	if b == nil {
		return
	}
	descWalk(b.right, arr)
	*arr = append(*arr, b.data)
	descWalk(b.left, arr)

}

func breadthWalk(b *bst, arr *[]int) {
	currentLevel := []*bst{b}
	for len(currentLevel) != 0 {
		nextLevel := []*bst{}
		for _, node := range currentLevel {
			*arr = append(*arr, node.data)
			if node.left != nil {
				nextLevel = append(nextLevel, node.left)
			}
			if node.right != nil {
				nextLevel = append(nextLevel, node.right)
			}
		}
		currentLevel = nextLevel
	}
}

func isValidBst(b *bst) bool {
	if b == nil {
		return true
	}
	if b.left != nil && b.left.data >= b.data {
		return false
	}
	if b.right != nil && b.right.data < b.data {
		return false
	}
	return isValidBst(b.left) && isValidBst(b.right)

}

// func getMinNode(b *bst) *bst{
// 	if b.left == nil {
// 		return b
// 	}
// 	return getMinNode(b.left)
// }

func delNodeWithTwoChildren(b *bst) *bst {
	successor := findMin(b.right)
	delNodeFromBst(successor, b)
	b.data = successor
	return b

}

func delNodeFromBst(i int, b *bst) *bst {
	if b == nil {
		panic("not exit, can not del")
	}

	if i == b.data {
		if b.left == nil && b.right == nil {
			return nil
		} else if b.left != nil && b.right != nil {
			successorData := findSuccessorNode(b.right, nil)
			b.data = successorData
			// successor = nil
			return b
			//return delNodeWithTwoChildren(b)
		} else if b.left == nil {
			return b.right
		} else {
			return b.left
		}

	} else if i < b.data {
		b.left = delNodeFromBst(i, b.left)
	} else {
		b.right = delNodeFromBst(i, b.right)
	}
	return b
}
func findSuccessorNode(b, bParent *bst) int {
	if b.left == nil {
		res := b.data
		if bParent != nil {
			bParent.left = nil
		}
		b = nil
		return res
	} else {
		return findSuccessorNode(b.left, b)
	}
}
func searchNodeFromBst(i int, b *bst) *bst {
	if b == nil {
		panic("not exit")
	}
	if i == b.data {
		return b
	} else if i < b.data {
		return searchNodeFromBst(i, b.left)
	} else {
		return searchNodeFromBst(i, b.right)
	}

}
func findMin(b *bst) int {

	if b.left == nil {
		return b.data
	}
	return findMin(b.left)
}
func findMax(b *bst) int {

	if b.right == nil {
		return b.data
	}
	return findMax(b.right)
}

func findHeight(b *bst, height int) int {
	if b == nil {
		return height
	}

	left := findHeight(b.left, height+1)
	right := findHeight(b.right, height+1)
	fmt.Println("left is ", left, " right is ", right)
	if left < right {
		return right
	}
	return left
}

func findHeight1(b *bst) int {
	if b == nil {
		return -1
	}

	left := findHeight1(b.left)
	right := findHeight1(b.right)
	fmt.Println("left is ", left, " right is ", right)
	if left > right {
		return left + 1
	}
	return right + 1

}

type CtxPath string

const (
	ctxPath CtxPath = "ctxPath"
)

func getPaths(root *bst) ([]string, []string) {
	var maxDepthHelper func(root *bst, ctx context.Context, isLeft bool, res *[][]string)
	maxDepthHelper = func(root *bst, ctx context.Context, isLeft bool, res *[][]string) {

		if root == nil {
			return
		}
		val := ctx.Value(ctxPath).([]string)

		if isLeft {
			val = append(val, "L")
		} else {
			val = append(val, "R")
		}
		// leaf node
		if root.left == nil && root.right == nil {
			// needcopy
			tmp := []string{}
			tmp = append(tmp, val...)
			//fmt.Println(tmp)
			// if len(tmp) == 4 && tmp[2] == "R" && tmp[3] == "L" && tmp[0] == "R" && tmp[1] == "R" {
			// 	fmt.Println("")
			// }
			*res = append(*res, tmp)
			return
		}
		ctxWithValue := context.WithValue(ctx, ctxPath, val)
		maxDepthHelper(root.left, ctxWithValue, true, res)
		maxDepthHelper(root.right, ctxWithValue, false, res)
		//return

	}

	res := [][]string{}
	ctxWithValue := context.WithValue(context.Background(), ctxPath, []string{})
	maxDepthHelper(root.left, ctxWithValue, true, &res)
	maxDepthHelper(root.right, ctxWithValue, false, &res)

	fmt.Println("the res is ", res)
	var maxDepth, minDepth []string = res[0], res[0]
	for i := 1; i < len(res); i++ {

		if len(res[i]) > len(maxDepth) {
			maxDepth = res[i]
		}
		if len(res[i]) < len(minDepth) {
			minDepth = res[i]
		}
	}
	return minDepth, maxDepth

}
