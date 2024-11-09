package binaryTree

import (
	"context"
	"fmt"
	"sync"
)

type Node struct {
	num   int
	Left  *Node
	Right *Node
}

func isValidBstNode(node *Node, min int, max int) bool {
	if node == nil {
		return true
	}
	if node.num <= min || node.num >= max {
		return false
	}
	return isValidBstNode(node.Left, min, node.num) && isValidBstNode(node.Right, node.num, max)
}

func isBinarySearchTree(root *Node) bool {
	if root == nil {
		return true
	}
	return isValidBstNode(root, -999999, 999999)
}

type CtxPath string

const (
	ctxPath CtxPath = "ctxPath"
)

func maxDepth(root *Node) ([]string, []string) {
	var maxDepthHelper func(root *Node, ctx context.Context, isLeft bool, wg *sync.WaitGroup, res *[][]string)
	maxDepthHelper = func(root *Node, ctx context.Context, isLeft bool, wg *sync.WaitGroup, res *[][]string) {
		wg.Add(1)
		defer wg.Done()

		val := ctx.Value(ctxPath).([]string)

		if isLeft {
			val = append(val, "L")

		} else {
			val = append(val, "R")
		}
		tmp := []string{}
		tmp = append(tmp, val...)
		if root == nil {
			fmt.Println(tmp)
			// if len(tmp) == 4 && tmp[2] == "R" && tmp[3] == "L" && tmp[0] == "R" && tmp[1] == "R" {
			// 	fmt.Println("")
			// }
			*res = append(*res, tmp)
			return
		}
		ctxWithValue := context.WithValue(ctx, ctxPath, val)
		maxDepthHelper(root.Left, ctxWithValue, true, wg, res)
		maxDepthHelper(root.Right, ctxWithValue, false, wg, res)
		//return

	}
	wg := sync.WaitGroup{}
	res := [][]string{}
	ctxWithValue := context.WithValue(context.Background(), ctxPath, []string{})
	maxDepthHelper(root.Left, ctxWithValue, true, &wg, &res)
	maxDepthHelper(root.Right, ctxWithValue, false, &wg, &res)
	wg.Wait()
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

func middleOrderWalk(root *Node) {
	if root == nil {
		return
	}
	middleOrderWalk(root.Left)
	fmt.Print("	", root.num)
	middleOrderWalk(root.Right)
}

func preOrderWalk(root *Node) {
	if root == nil {
		return
	}
	fmt.Print("	", root.num)
	preOrderWalk(root.Left)
	preOrderWalk(root.Right)
}

func postOrderWalk(root *Node) {
	if root == nil {
		return
	}
	postOrderWalk(root.Left)
	postOrderWalk(root.Right)
	fmt.Print("	", root.num)

}

func levelWalk(root *Node) {

	var tmp func(currentLevel []*Node)

	tmp = func(currentLevel []*Node) {
		if len(currentLevel) == 0 {
			return
		}
		nextLevel := []*Node{}
		for i := 0; i < len(currentLevel); i++ {
			tmp := currentLevel[i]
			if tmp.Left != nil {
				nextLevel = append(nextLevel, tmp.Left)
			}
			if tmp.Right != nil {
				nextLevel = append(nextLevel, tmp.Right)
			}
			fmt.Print(" ", tmp.num)
		}
		tmp(nextLevel)
	}

	tmp([]*Node{root})
}
