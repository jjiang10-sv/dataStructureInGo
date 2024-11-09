package graph

import (
	"container/heap"
)

func robHouseDp(nums []uint) uint {
	houses := len(nums)
	if houses == 0 {
		return 0
	} else if houses == 1 {
		return nums[0]
	} else if houses == 2 {
		return maxVal(nums[0], nums[1])
	} else {
		dp := make([]uint, houses)
		dp[0] = nums[0]
		dp[1] = maxVal(nums[0], nums[1])
		for i := 2; i < houses; i++ {
			dp[i] = maxVal(nums[i]+dp[i-2], dp[i-1])
		}
		return dp[houses-1]
	}

}

func maxVal(a, b uint) uint {
	if a > b {
		return a
	}
	return b
}

func zagLagEncrypt(s string, depth uint) string {
	rows := make([]string, depth)

	row := 0
	down := false
	for _, char := range s {
		rows[row] += string(char)
		if row == 0 || row == (int(depth)-1) {
			down = (!down)
		}
		if down {
			row += 1
		} else {
			row -= 1
		}
	}
	res := ""
	for _, row := range rows {
		res += row
	}
	return res
}

// func zagLagDecrypt(s string, depth uint) string {
// 	rails := make([][]rune,depth)

// 	row, col  := 0,0
// 	down := false
// 	for j := 0 ; j < len(s); j++ {
// 		rails[row][col] = '*'
// 		if row == 0 || row == (int(depth)-1) {
// 			down = (!down)
// 		}
// 		if down{
// 			row+=1
// 		}else{
// 			row-=1
// 		}
// 		col += 1

// 	}
// 	index := 0
// 	res := ""
// 	for i := 0; i < int(depth) ; i++{
// 		for j := 0 ; j<= len(s); j++ {
// 			if rails[i][j] == '*' && index < len(s) {
// 				rails[i][j] = rune(s[index])
// 				index++
// 			}
// 		}
// 	}

// }

type nodeEle struct {
	val       string
	frequency uint
	left      *nodeEle
	right     *nodeEle
}

func NewNodeEle(val string, frequency uint) *nodeEle {
	return &nodeEle{val: val, frequency: frequency}
}

type nodeEles struct {
	data []*nodeEle
}

func NewNodeEles(data []*nodeEle) *nodeEles {
	return &nodeEles{data: data}
}
func (n *nodeEles) Less(i, j int) bool {
	eles := n.data
	return eles[i].frequency < eles[j].frequency
}
func (n *nodeEles) Swap(i, j int) {
	eles := n.data
	eles[i], eles[j] = eles[j], eles[i]
}

func (n *nodeEles) Len() int {
	return len(n.data)
}

func (n *nodeEles) Push(item interface{}) {
	nodeEle := item.(*nodeEle)
	n.data = append(n.data, nodeEle)

}

func (n *nodeEles) Pop() interface{} {
	len := n.Len()
	if n.Len() > 0 {
		res := n.data[len-1]
		n.data = n.data[:len-1]
		return res
	} else {
		return nil
	}

}

func hoffman_code_tree(vals []string, frequencies []uint) *nodeEle {

	nodeEleArr := make([]*nodeEle, 0)
	for i, val := range vals {
		nodeEleArr = append(nodeEleArr, NewNodeEle(val, frequencies[i]))
	}
	nodeEles := NewNodeEles(nodeEleArr)
	heap.Init(nodeEles)
	for nodeEles.Len() > 1 {

		var nodeEleLeft, nodeEleRight, parentNode *nodeEle
		eleLeft := heap.Pop(nodeEles)
		nodeEleLeft = eleLeft.(*nodeEle)
		eleRight := heap.Pop(nodeEles)
		nodeEleRight = eleRight.(*nodeEle)
		parentNode = NewNodeEle("", nodeEleLeft.frequency+nodeEleRight.frequency)
		parentNode.left = nodeEleLeft
		parentNode.right = nodeEleRight
		heap.Push(nodeEles, parentNode)

	}
	return nodeEles.data[0]

}

func (node *nodeEle) getValCodefromHoffmanTree(codes map[string]string, code string) {

	if node.val != "" {
		codes[node.val] = code
	} else {

		node.left.getValCodefromHoffmanTree(codes, code+"0")
		node.right.getValCodefromHoffmanTree(codes, code+"1")
	}
}

type hoffManCodes struct {
	codes map[string]string
}

func NewHoffmanCodes(codes map[string]string) *hoffManCodes{
	return &hoffManCodes{codes:codes}
}
func (h *hoffManCodes) code(s string) string{
	res := ""
	for _ , char := range(s) {
		res+=h.codes[string(char)]
	}
	return res
}