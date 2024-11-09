package main

import (
	"fmt"
	"testing"
	"time"
)

// var arr = []int{3, 5, 2, 11, 4, 1, 8, 2}
func TestCocktailSort(t *testing.T) {
	data1 := []int{3, 5, 2, 11, 4, 1, 8, 2}
	now := time.Now()
	CocktailSort0428(data1)
	cost1 := time.Since(now).Nanoseconds()
	data2 := []int{3, 5, 2, 11, 4, 1, 8, 2}
	fmt.Println(data2)
	now = time.Now()
	BubbleSourt0428(data2)
	cost2 := time.Since(now).Nanoseconds()
	fmt.Println(data2)
	// if cost2 > cost1 {
	// 	fmt.Println(cost2-cost1)
	// }
	fmt.Println(cost2 - cost1)
	fmt.Println(cost2)
	fmt.Println(cost1)
}

func TestCombSort(t *testing.T) {
	data1 := []int{3, 5, 2, 11, 4, 1, 8, 2}
	//data1 = append(data1, data1...)
	//now := time.Now()
	//CombSort(data1)
	//CountingSort0428(data1,12)
	data1 = quickSort0430(data1)
	//cost1 := time.Since(now).Nanoseconds()

	fmt.Println(data1)
}

func TestTreeSort(t *testing.T) {
	data1 := []int{3, 5, 2, 11, 4, 1, 8, 2}
	bst := new(bst)

	bst.data = data1[0]
	for i := 1; i < len(data1); i++ {
		insertNode(data1[i], bst)
	}
	data1 = loadBstData(bst)

	fmt.Println(data1)
	isValidBst(bst)

	node := searchNodeFromBst(5, bst)
	fmt.Println(node.data)

	getPaths(bst)
	delNodeFromBst(5, bst)
	data1 = loadBstData(bst)

	fmt.Println("after delete", data1)

	min := findMin(bst)
	fmt.Println(min)
	max := findMax(bst)
	fmt.Println(max)

	height := findHeight(bst, -1)
	fmt.Println(height)
}

func TestTreeFindHeight(t *testing.T) {

	bst := new(bst)

	bst.data = 1
	bst.left = newBst(0)
	bst.right = newBst(2)
	bst.right.left = newBst(1)
	bst.right.right = newBst(3)
	bst.right.right.right = newBst(11)
	bst.right.right.right.left = newBst(3)
	bst.right.right.right.left.right = newBst(8)
	bst.right.right.right.left.right.left = newBst(4)
	//height := findHeight(bst,-1)
	height := findHeight1(bst)
	fmt.Println(height)
}
