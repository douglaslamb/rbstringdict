package rbstringdict

import "testing"

func TestInsert(t *testing.T) {
	tree := NewStringRBTree()
	tree.insert("foo")
	if tree.rootNode.value != "foo" {
		t.Errorf("tree.contains(\"foo\") = %v; want %v", tree.contains("foo"), true)
	}
}

func TestInsertFixup(t *testing.T) {
	tree := NewStringRBTree()
	tree.insert("m")
	tree.insert("g")
	tree.insert("b")
	tree.insert("d")
	tree.insert("k")
	tree.insert("jkdk")
	tree.insert("h")
	tree.insert("t")
	tree.insert("q")
	tree.insert("w")
	tree.insert("y")
	tree.insert("t")
	tree.insert("z")
	tree.insert("w")
	if !tree.isRedBlackTree() {
		t.Errorf("tree.isRedBlackTree() = %v; want %v", tree.isRedBlackTree(), true)
	}
}

func TestRemove(t *testing.T) {
	tree := NewStringRBTree()
	node := tree.newStringNode()
	node.value = "foo"
	tree.setRoot(node)
	tree.remove("foo")
	if tree.contains("foo") {
		t.Errorf("tree.contains(\"foo\") = %v; want %v", tree.contains("foo"), false)
	}

	// breaks red black tree
	tree = NewStringRBTree()
	tree.insert("w")
	tree.insert("c")
	tree.insert("a")
	tree.insert("b")
	tree.remove("w")
	if !tree.isRedBlackTree() {
		t.Errorf("tree.isRedBlackTree() = %v; want %v", tree.isRedBlackTree(), true)
	}
}

func TestContains(t *testing.T) {
	tree := NewStringRBTree()
	node := tree.newStringNode()
	node.value = "foo"
	tree.rootNode = node
	if !tree.contains("foo") {
		t.Errorf("tree.contains(\"foo\") = %v; want %v", tree.contains("foo"), true)
	}
}

func TestLeftRotate(t *testing.T) {
	// rotation when node has parent
	tree := NewStringRBTree()
	parent := tree.newStringNode()
	node := tree.newStringNode()
	right := tree.newStringNode()
	rightLeftChild := tree.newStringNode()
	tree.setLeft(parent, node)
	tree.setRight(node, right)
	tree.setLeft(right, rightLeftChild)
	tree.rootNode = parent
	tree.leftRotate(node)
	if parent.left != right {
		t.Errorf("leftRotate failed; parent.left = %v; want %v", parent.left, right)
	}
	if node.right != rightLeftChild {
		t.Errorf("leftRotate failed; node.right = %v; want %v", node.right, rightLeftChild)
	}
	if right.left != node {
		t.Errorf("leftRotate failed; right.left = %v; want %v", right.left, node)
	}

	// rotation when node is rootNode
	tree = NewStringRBTree()
	node = tree.newStringNode()
	right = tree.newStringNode()
	rightLeftChild = tree.newStringNode()
	tree.rootNode = node
	tree.setRight(node, right)
	tree.setLeft(right, rightLeftChild)
	tree.leftRotate(node)
	if tree.rootNode != right {
		t.Errorf("leftRotate failed; tree.rootNode = %v; want %v", tree.rootNode, right)
	}
	if node.right != rightLeftChild {
		t.Errorf("leftRotate failed; node.right = %v; want %v", node.right, rightLeftChild)
	}
	if right.left != node {
		t.Errorf("leftRotate failed; right.left = %v; want %v", right.left, node)
	}
}

func TestRightRotate(t *testing.T) {
	// rotation when node has parent
	tree := NewStringRBTree()
	parent := tree.newStringNode()
	node := tree.newStringNode()
	left := tree.newStringNode()
	leftRightChild := tree.newStringNode()
	tree.setRight(parent, node)
	tree.setLeft(node, left)
	tree.setRight(left, leftRightChild)
	tree.rootNode = parent
	tree.rightRotate(node)
	if parent.right != left {
		t.Errorf("rightRotate failed; parent.right = %v; want %v", parent.right, left)
	}
	if node.left != leftRightChild {
		t.Errorf("rightRotate failed; node.left = %v; want %v", node.left, leftRightChild)
	}
	if left.right != node {
		t.Errorf("rightRotate failed; left.right = %v; want %v", left.right, node)
	}

	// rotation when node is rootNode
	tree = NewStringRBTree()
	node = tree.newStringNode()
	left = tree.newStringNode()
	leftRightChild = tree.newStringNode()
	tree.rootNode = node
	tree.setLeft(node, left)
	tree.setRight(left, leftRightChild)
	tree.rightRotate(node)
	if tree.rootNode != left {
		t.Errorf("rightRotate failed; tree.rootNode = %v; want %v", tree.rootNode, left)
	}
	if node.left != leftRightChild {
		t.Errorf("rightRotate failed; node.left = %v; want %v", node.left, leftRightChild)
	}
	if left.right != node {
		t.Errorf("rightRotate failed; left.right = %v; want %v", left.right, node)
	}
}

func TestIsBST(t *testing.T) {
	// not BST
	tree := NewStringRBTree()
	root := tree.newStringNode()
	root.value = `f`
	left := tree.newStringNode()
	left.value = `z`
	tree.rootNode = root
	root.left = left
	if tree.isBST() {
		t.Errorf("tree.isBST() = %v; want %v", tree.isBST(), false)
	}

	// is BST
	tree = NewStringRBTree()
	root = tree.newStringNode()
	root.value = `f`
	left = tree.newStringNode()
	left.value = `a`
	tree.rootNode = root
	root.left = left
	if !tree.isBST() {
		t.Errorf("tree.isBST() = %v; want %v", tree.isBST(), true)
	}
}

func TestIsRedBlackTree(t *testing.T) {
	// not red black tree
	// root is red
	tree := NewStringRBTree()
	root := tree.newStringNode()
	tree.rootNode = root
	if tree.isRedBlackTree() {
		t.Errorf("tree.isRedBlackTree() = %v; want %v", tree.isRedBlackTree(), false)
	}

	// not red black tree
	// red node with red children
	tree = NewStringRBTree()
	root = tree.newStringNode()
	parent := tree.newStringNode()
	left := tree.newStringNode()
	right := tree.newStringNode()
	tree.rootNode = root
	root.isBlack = true
	root.left = parent
	parent.left = left
	parent.right = right
	if tree.isRedBlackTree() {
		t.Errorf("tree.isRedBlackTree() = %v; want %v", tree.isRedBlackTree(), false)
	}

	// not red black tree
	// unequal black heights
	tree = NewStringRBTree()
	root = tree.newStringNode()
	left = tree.newStringNode()
	tree.rootNode = root
	root.isBlack = true
	root.left = left
	left.isBlack = true
	if tree.isRedBlackTree() {
		t.Errorf("tree.isRedBlackTree() = %v; want %v", tree.isRedBlackTree(), false)
	}

	// is red black tree
	tree = NewStringRBTree()
	root = tree.newStringNode()
	parent = tree.newStringNode()
	uncle := tree.newStringNode()
	left = tree.newStringNode()
	right = tree.newStringNode()
	tree.rootNode = root
	root.isBlack = true
	root.left = parent
	root.right = uncle
	uncle.isBlack = true
	parent.left = left
	parent.right = right
	left.isBlack = true
	right.isBlack = true
	if !tree.isRedBlackTree() {
		t.Errorf("tree.isRedBlackTree() = %v; want %v", tree.isRedBlackTree(), true)
	}
}

func TestSetLeft(t *testing.T) {
	tree := NewStringRBTree()
	parent := tree.newStringNode()
	formerChild := tree.newStringNode()
	tree.setLeft(parent, formerChild)
	child := tree.newStringNode()
	formerParent := tree.newStringNode()
	tree.setLeft(formerParent, child)
	tree.setLeft(parent, child)
	if parent.left != child {
		t.Errorf("setLeft failed; parent.left = %v; want %v", parent.left, child)
	}
	if child.parent != parent {
		t.Errorf("setLeft failed; child.parent = %v; want %v", child.parent, parent)
	}
	if formerChild.parent != tree.dummy {
		t.Errorf("setLeft failed; formerChild.parent = %v; want %v", formerChild.parent, tree.dummy)
	}
	if formerParent.left != tree.dummy {
		t.Errorf("setLeft failed; formerParent.left = %v; want %v", formerParent.left, tree.dummy)
	}
}

func TestSetRight(t *testing.T) {
	tree := NewStringRBTree()
	parent := tree.newStringNode()
	formerChild := tree.newStringNode()
	tree.setRight(parent, formerChild)
	child := tree.newStringNode()
	formerParent := tree.newStringNode()
	tree.setRight(formerParent, child)
	tree.setRight(parent, child)
	if parent.right != child {
		t.Errorf("setRight failed; parent.right = %v; want %v", parent.right, child)
	}
	if child.parent != parent {
		t.Errorf("setRight failed; child.parent = %v; want %v", child.parent, parent)
	}
	if formerChild.parent != tree.dummy {
		t.Errorf("setRight failed; formerChild.parent = %v; want %v", formerChild.parent, tree.dummy)
	}
	if formerParent.right != tree.dummy {
		t.Errorf("setRight failed; formerParent.right = %v; want %v", formerParent.right, tree.dummy)
	}
}

func TestDetachParent(t *testing.T) {
	tree := NewStringRBTree()
	parent := tree.newStringNode()
	leftChild := tree.newStringNode()
	tree.setLeft(parent, leftChild)
	rightChild := tree.newStringNode()
	tree.setRight(parent, rightChild)
	tree.detachParent(leftChild)
	tree.detachParent(rightChild)
	if parent.left != tree.dummy {
		t.Errorf("detachParent failed; parent.left = %v; want %v", parent.left, tree.dummy)
	}
	if parent.right != tree.dummy {
		t.Errorf("detachParent failed; parent.right = %v; want %v", parent.right, tree.dummy)
	}
	if leftChild.parent != tree.dummy {
		t.Errorf("detachParent failed; leftChild.parent = %v; want %v", leftChild.parent, tree.dummy)
	}
	if rightChild.parent != tree.dummy {
		t.Errorf("detachParent failed; rightChild.parent = %v; want %v", rightChild.parent, tree.dummy)
	}
}
