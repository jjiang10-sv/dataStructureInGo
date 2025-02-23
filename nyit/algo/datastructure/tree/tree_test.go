package tree

import (
	"testing"
)

func Test_avlTree(t *testing.T) {
	avlTree := &AVLTree{root: InitAVLNode(5)}

	vals := []int{4, 6, 2, 1, 8, 9, 11, 2, 32, 21, 15}

	//vals := []int{11}
	for _, val := range vals {
		avlTree.insert(val)
	}
	avlTree.delete(9)
	avlTree.delete(11)
	avlTree.delete(6)
	avlTree.delete(2)
	avlTree.delete(4)
	avlTree.insert(35)
	avlTree.delete(1)
	
	//avlTree.preorder()
	avlTree.inorder()
}
