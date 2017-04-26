package main

import "testing"

func TestInsert(t *testing.T) {
	tree := &StringRBTree{}
	tree.insert("foo")
	if tree.rootNode.value != "foo" {
		t.Errorf("tree.contains(\"foo\") = %v; want %v", tree.contains("foo"), true)
	}
}

func TestRemove(t *testing.T) {
	tree := &StringRBTree{}
	node := &StringNode{}
	node.value = "foo"
	tree.rootNode = node
	tree.remove("foo")
	if tree.contains("foo") {
		t.Errorf("tree.contains(\"foo\") = %v; want %v", tree.contains("foo"), false)
	}
}

func TestContains(t *testing.T) {
	tree := &StringRBTree{}
	node := &StringNode{}
	node.value = "foo"
	tree.rootNode = node
	if !tree.contains("foo") {
		t.Errorf("tree.contains(\"foo\") = %v; want %v", tree.contains("foo"), true)
	}
}

func TestLeftRotate(t *testing.T) {
	// rotation when node has parent
	tree := &StringRBTree{}
	parent := &StringNode{}
	node := &StringNode{}
	right := &StringNode{}
	rightLeftChild := &StringNode{}
	parent.setLeft(node)
	node.setRight(right)
	right.setLeft(rightLeftChild)
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
	tree = &StringRBTree{}
	node = &StringNode{}
	right = &StringNode{}
	rightLeftChild = &StringNode{}
	tree.rootNode = node
	node.setRight(right)
	right.setLeft(rightLeftChild)
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
	tree := &StringRBTree{}
	parent := &StringNode{}
	node := &StringNode{}
	left := &StringNode{}
	leftRightChild := &StringNode{}
	parent.setRight(node)
	node.setLeft(left)
	left.setRight(leftRightChild)
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
	tree = &StringRBTree{}
	node = &StringNode{}
	left = &StringNode{}
	leftRightChild = &StringNode{}
	tree.rootNode = node
	node.setLeft(left)
	left.setRight(leftRightChild)
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
	tree := &StringRBTree{}
	root := &StringNode{}
	root.value = `f`
	left := &StringNode{}
	left.value = `z`
	tree.rootNode = root
	root.left = left
	if tree.isBST() {
		t.Errorf("tree.isBST() = %v; want %v", tree.isBST(), false)
	}

	// is BST
	tree = &StringRBTree{}
	root = &StringNode{}
	root.value = `f`
	left = &StringNode{}
	left.value = `a`
	tree.rootNode = root
	root.left = left
	if !tree.isBST() {
		t.Errorf("tree.isBST() = %v; want %v", tree.isBST(), true)
	}
}
