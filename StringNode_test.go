package main

import "testing"

func TestSetLeft(t *testing.T) {
	parent := &StringNode{}
	formerChild := &StringNode{}
	parent.setLeft(formerChild)
	child := &StringNode{}
	formerParent := &StringNode{}
	formerParent.setLeft(child)
	parent.setLeft(child)
	if parent.left != child {
		t.Errorf("parent.setLeft failed; parent.left = %v; want %v", parent.left, child)
	}
	if child.parent != parent {
		t.Errorf("parent.setLeft failed; child.parent = %v; want %v", child.parent, parent)
	}
	if formerChild.parent != nil {
		t.Errorf("parent.setLeft failed; formerChild.parent = %v; want %v", formerChild.parent, nil)
	}
	if formerParent.left != nil {
		t.Errorf("parent.setLeft failed; formerParent.left = %v; want %v", formerParent.left, nil)
	}
}

func TestSetRight(t *testing.T) {
	parent := &StringNode{}
	formerChild := &StringNode{}
	parent.setRight(formerChild)
	child := &StringNode{}
	formerParent := &StringNode{}
	formerParent.setRight(child)
	parent.setRight(child)
	if parent.right != child {
		t.Errorf("parent.setRight failed; parent.right = %v; want %v", parent.right, child)
	}
	if child.parent != parent {
		t.Errorf("parent.setRight failed; child.parent = %v; want %v", child.parent, parent)
	}
	if formerChild.parent != nil {
		t.Errorf("parent.setRight failed; formerChild.parent = %v; want %v", formerChild.parent, nil)
	}
	if formerParent.right != nil {
		t.Errorf("parent.setRight failed; formerParent.right = %v; want %v", formerParent.right, nil)
	}
}

func TestDetachParent(t *testing.T) {
	parent := &StringNode{}
	leftChild := &StringNode{}
	parent.setLeft(leftChild)
	rightChild := &StringNode{}
	parent.setRight(rightChild)
	leftChild.detachParent()
	rightChild.detachParent()
	if parent.left != nil {
		t.Errorf("detachParent failed; parent.left = %v; want %v", parent.left, nil)
	}
	if parent.right != nil {
		t.Errorf("detachParent failed; parent.right = %v; want %v", parent.right, nil)
	}
	if leftChild.parent != nil {
		t.Errorf("detachParent failed; leftChild.parent = %v; want %v", leftChild.parent, nil)
	}
	if rightChild.parent != nil {
		t.Errorf("detachParent failed; rightChild.parent = %v; want %v", rightChild.parent, nil)
	}
}
