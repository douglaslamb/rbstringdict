package main

import "testing"

func TestIsLeftChild(t *testing.T) {
	// is left child
	parent := &StringNode{}
	child := &StringNode{}
	parent.left = child
	child.parent = parent
	isLeft, ok := child.isLeftChild()
	if !ok || !isLeft {
		t.Errorf("child.isLeftChild() = (%v, %v); want true, true)", isLeft, ok)
	}
	// is right child
	parent = &StringNode{}
	child = &StringNode{}
	parent.right = child
	child.parent = parent
	isLeft, ok = child.isLeftChild()
	if !ok || isLeft {
		t.Errorf("child.isLeftChild() = (%v, %v); want false, true)", isLeft, ok)
	}
	// has no parent
	child = &StringNode{}
	isLeft, ok = child.isLeftChild()
	if ok {
		t.Errorf("child.isLeftChild() = (%v, %v); want false, false)", isLeft, ok)
	}
}

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
