package binaryTree

import (
	"fmt"
	"testing"
)

func compare(x interface{}, y interface{}) bool {
	return x.(int) < y.(int)
}

func Test_binaryTree(t *testing.T) {
	tree := New(compare)

	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(3)

	findTree := tree.Search(2)
	if findTree.node != 2 {
		t.Error("[Error] Search error")
	}

	findNilTree := tree.Search(100)

	if findNilTree != nil {
		t.Error("[Error] 2. Search error")
	}
}

func Test_minmax(t *testing.T) {
	tree := New(compare)

	testValues := []int{4, 5, 3, 2, 9}
	for _, i := range testValues {
		tree.Insert(i)
	}

	max := tree.Max()
	if max != 9 {
		t.Errorf("[Error] max: expected 9, got %d", max)
	}

	min := tree.Min()
	if min != 2 {
		t.Errorf("[Error] max: expected 2, got %d", min)
	}
}

func Test_IsBst(t *testing.T) {
	root := Node{num: 6}
	root.Left = &Node{num: 4}
	root.Right = &Node{num: 7}
	root.Left.Left = &Node{num: 1}
	root.Left.Right = &Node{num: 3}
	root.Right.Left = &Node{num: 2}
	root.Right.Right = &Node{num: 8}
	root.Right.Right.Right = &Node{num: 10}
	res := isBinarySearchTree(&root)
	fmt.Println(res)
}

func Test_IsBst1(t *testing.T) {
	root := Node{num: 6}
	root.Left = &Node{num: 4}
	root.Right = &Node{num: 7}
	root.Left.Left = &Node{num: 1}
	root.Left.Left.Right = &Node{num: 3}
	root.Left.Left.Right.Left = &Node{num: 2}
	root.Right.Right = &Node{num: 8}
	root.Right.Right.Right = &Node{num: 10}
	res := isBinarySearchTree(&root)
	fmt.Println(res)
}

func TestMaxDepth(t *testing.T) {
	root := Node{num: 6}
	root.Left = &Node{num: 4}
	root.Right = &Node{num: 7}
	root.Left.Left = &Node{num: 1}
	root.Left.Left.Right = &Node{num: 3}
	root.Left.Left.Right.Left = &Node{num: 2}
	root.Right.Right = &Node{num: 8}
	root.Right.Right.Right = &Node{num: 10}
	fmt.Println(maxDepth(&root))

}

func TestMiddleOrderWalk(t *testing.T) {
	root := Node{num: 6}
	root.Left = &Node{num: 4}
	root.Right = &Node{num: 7}
	root.Left.Left = &Node{num: 1}
	root.Left.Left.Right = &Node{num: 3}
	root.Left.Left.Right.Left = &Node{num: 2}
	root.Right.Right = &Node{num: 8}
	root.Right.Right.Right = &Node{num: 10}
	middleOrderWalk(&root)
	fmt.Println("")

}

func TestPreOrderWalk(t *testing.T) {
	root := Node{num: 6}
	root.Left = &Node{num: 4}
	root.Right = &Node{num: 7}
	root.Left.Left = &Node{num: 1}
	root.Left.Left.Right = &Node{num: 3}
	root.Left.Left.Right.Left = &Node{num: 2}
	root.Right.Right = &Node{num: 8}
	root.Right.Right.Right = &Node{num: 10}
	preOrderWalk(&root)
	fmt.Println("")
}

func TestPostOrderWalk(t *testing.T) {
	root := Node{num: 6}
	root.Left = &Node{num: 4}
	root.Right = &Node{num: 7}
	root.Left.Left = &Node{num: 1}
	root.Left.Left.Right = &Node{num: 3}
	root.Left.Left.Right.Left = &Node{num: 2}
	root.Right.Right = &Node{num: 8}
	root.Right.Right.Right = &Node{num: 10}
	postOrderWalk(&root)
	fmt.Println("")
}

func TestPostOrderWalk1(t *testing.T) {
	root := Node{num: 6}
	root.Left = &Node{num: 4}
	root.Right = &Node{num: 7}
	root.Left.Left = &Node{num: 1}
	root.Left.Right = &Node{num: 3}
	root.Right.Left = &Node{num: 2}
	root.Right.Right = &Node{num: 8}
	root.Right.Right.Right = &Node{num: 10}
	postOrderWalk(&root)
	fmt.Println("")
}


func TestLevelWalk(t *testing.T) {
	root := Node{num: 6}
	root.Left = &Node{num: 4}
	root.Right = &Node{num: 7}
	root.Left.Left = &Node{num: 1}
	root.Left.Right = &Node{num: 3}
	root.Right.Left = &Node{num: 2}
	root.Right.Right = &Node{num: 8}
	root.Right.Right.Right = &Node{num: 10}
	levelWalk(&root)
	fmt.Println("")
}

func TestLevelWalk1(t *testing.T) {
	root := Node{num: 6}
	root.Left = &Node{num: 4}
	root.Right = &Node{num: 7}
	root.Left.Left = &Node{num: 1}
	root.Left.Left.Right = &Node{num: 3}
	root.Left.Left.Right.Left = &Node{num: 2}
	root.Right.Right = &Node{num: 8}
	root.Right.Right.Right = &Node{num: 10}
	levelWalk(&root)
	fmt.Println("")
}
