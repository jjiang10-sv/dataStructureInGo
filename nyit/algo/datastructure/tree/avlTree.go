package tree

type AVLNode struct {
	val    int
	left   *AVLNode
	right  *AVLNode
	height int
}

type AVLTree struct {
	root *AVLNode
}

func InitAVLNode(val int) *AVLNode {
	return &AVLNode{
		val:    val,
		height: 1,
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (t *AVLTree) getHeight(node *AVLNode) int {
	if node == nil {
		return 0
	}
	return node.height
}

func (t *AVLTree) updateHeight(node *AVLNode) int {
	defer func() {
		if panicVal := recover(); panicVal != nil {
			println(panicVal)
		}
	}()
	if node == nil {
		return 0
	}
	return 1 + max(t.getHeight(node.left), t.getHeight(node.right))
}

func (t *AVLTree) isLeafNode(node *AVLNode) bool {
	if node == nil {
		return true
	}
	if node.left == nil && node.right == nil {
		return true
	}
	return false
}

func (t *AVLTree) getBalance(node *AVLNode) int {
	return t.getHeight(node.left) - t.getHeight(node.right)

}

func (t *AVLTree) copyNode(node *AVLNode) *AVLNode {
	return &AVLNode{val: node.val, height: node.height, left: node.left, right: node.right}
}

func (t *AVLTree) leftRotate(node *AVLNode) *AVLNode {
	println("left rotate on ", node.val)
	// going to return t.right as root so get its left to be the old root's right
	res := node.right
	rightLeft := res.left
	// the old root becomes the new root's left
	res.left = node
	node.right = rightLeft
	node.height = t.updateHeight(node)
	res.height = t.updateHeight(res)
	return res
}

func (t *AVLTree) rightRotate(node *AVLNode) *AVLNode {
	println("right rotate on ", node.val)
	if node.left == nil {
		println("")
	}
	// going to return t.left as root so get its right to be the old root's left
	res := node.left
	leftRight := res.right
	if node.left == nil {
		println("")
	}
	// the old root becomes the new root's right
	res.right = node
	node.left = leftRight
	node.height = t.updateHeight(node)
	res.height = t.updateHeight(res)
	return res
}

func (t *AVLTree) insert(val int) {
	t.root = t.insertAtNode(t.root, val)
}

func (t *AVLTree) insertAtNode(node *AVLNode, val int) *AVLNode {

	if node == nil {
		return InitAVLNode(val)
	}
	if node.val < val {
		node.right = t.insertAtNode(node.right, val)
	} else if node.val > val {
		node.left = t.insertAtNode(node.left, val)
	} else {
		println("can not insert node with same val", val)
	}
	node.height = t.updateHeight(node)
	return t.balanceTreeAtInsertNode(node, val)
}

func (t *AVLTree) balanceTreeAtInsertNode(node *AVLNode, val int) *AVLNode {

	balance := t.getBalance(node)
	if balance < -1 {
		if val > node.right.val {
			node = t.leftRotate(node)
		} else {
			node.right = t.rightRotate(node.right)
			node = t.leftRotate(node)
		}
	} else if balance > 1 {
		if val < node.left.val {
			node = t.rightRotate(node)
		} else {
			node.left = t.leftRotate(node.left)
			node = t.rightRotate(node)
		}
	}
	return node
}

func (t *AVLTree) balanceTreeAtDeleteNode(node *AVLNode, val int) *AVLNode {
	if val == 1 {
		println("")
	}

	balance := t.getBalance(node)
	if balance < -1 {

		if node.left != nil && val > node.left.val {
			node.right = t.rightRotate(node.right)
			node = t.leftRotate(node)
		} else {
			node = t.leftRotate(node)
		}
	} else if balance > 1 {
		// delete roration is different. what if node.right is nil
		// need more testing
		if node.right != nil && val > node.right.val {
			node.left = t.leftRotate(node.left)
			node = t.rightRotate(node)
		} else {
			node = t.rightRotate(node)
		}
	}
	return node
}

func (t *AVLTree) delete(val int) {

	t.root = t.deleteAtNode(t.root, val)
}

func (t *AVLTree) deleteAtNode(node *AVLNode, val int) *AVLNode {
	if node == nil {
		println("node not found at value ", val)
		return nil
	}
	if node.val < val {
		node.right = t.deleteAtNode(node.right, val)
	} else if node.val > val {
		node.left = t.deleteAtNode(node.left, val)
	} else {
		if t.isLeafNode(node) {
			return nil
		} else if node.left != nil && node.right != nil {
			node.val, node.left = t.getRightMostNodeValAndLeftNode(node.left)
			return node
		} else if node.left == nil {
			return node.right
		} else {
			return node.left
		}
	}
	node.height = t.updateHeight(node)
	return t.balanceTreeAtDeleteNode(node, val)
}

func (t *AVLTree) getRightMostNodeValAndLeftNode(node *AVLNode) (int, *AVLNode) {
	if node.right == nil {
		val := node.val
		node = node.left
		return val, node
	}
	for node.right.right != nil {
		node = node.right
	}
	res := t.copyNode(node.right)
	node.right = res.left
	return res.val, node
}

func (t *AVLTree) preorder() {
	t.preorderAtNode(t.root)
}

func (t *AVLTree) preorderAtNode(node *AVLNode) {
	if node == nil {
		return
	}
	println(node.val)
	t.preorderAtNode(node.left)
	t.preorderAtNode(node.right)
}

func (t *AVLTree) inorder() {
	t.inorderAtNode(t.root)
}

func (t *AVLTree) inorderAtNode(node *AVLNode) {
	if node == nil {
		return
	}
	t.inorderAtNode(node.left)
	println(node.val)
	t.inorderAtNode(node.right)
}
