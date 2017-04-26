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
