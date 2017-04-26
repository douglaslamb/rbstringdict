package main

import "strings"

type StringRBTree struct {
	rootNode *StringNode
}

// insert adds a string to the dictionary.
func (s *StringRBTree) insert(key string) {
	_ = s.insertBST(key)
}

// insertBST inserts the provided key in the dictionary and returns
// the added node. If the key has already been inserted
// insertBST returns nil.
func (s *StringRBTree) insertBST(key string) *StringNode {
	// BST insert
	if s.isEmpty() {
		node := &StringNode{}
		node.value = key
		s.rootNode = node
		return s.rootNode
	}
	currNode := s.rootNode
	for {
		comparison := strings.Compare(key, currNode.value)
		if comparison < 0 {
			if currNode.left == nil {
				node := &StringNode{}
				node.value = key
				currNode.left = node
				node.parent = currNode
				return node
			} else {
				currNode = currNode.left
			}
		} else if comparison > 0 {
			if currNode.right == nil {
				node := &StringNode{}
				node.value = key
				currNode.right = node
				node.parent = currNode
				return node
			} else {
				currNode = currNode.right
			}
		} else {
			// key already existed in tree
			return nil
		}
	}
}

// remove removes a key from the dictionary.
func (s *StringRBTree) remove(key string) error {
	// BST remove
	currNode := s.rootNode
	for currNode != nil {
		comparison := strings.Compare(key, currNode.value)
		if comparison < 0 {
			currNode = currNode.left
		} else if comparison > 0 {
			currNode = currNode.right
		} else {
			// found key
			// check if left or right child and set parent pointer
			// to nil
			isLeft, hasParent := currNode.isLeftChild()
			if !hasParent {
				s.rootNode = nil
				return nil
			}
			if isLeft {
				currNode.parent.left = nil
			} else {
				currNode.parent.right = nil
			}
			return nil
		}
	}
	// key was not found
	return nil
}

// contains returns true if the key is in the dictionary
// and false otherwise.
func (s *StringRBTree) contains(key string) bool {
	if s.isEmpty() {
		return false
	}
	currNode := s.rootNode
	for {
		comparison := strings.Compare(key, currNode.value)
		if comparison < 0 {
			if currNode.left == nil {
				return false
			}
			currNode = currNode.left
		} else if comparison > 0 {
			if currNode.right == nil {
				return false
			}
			currNode = currNode.right
		} else {
			// key exists in tree
			return true
		}
	}
}

// isEmpty returns a boolean indicating whether the dictionary
// has no keys in it.
func (s *StringRBTree) isEmpty() bool {
	return s.rootNode == nil
}

// leftRotate rotates a node and its children to the left
// and updates its parent
func (s *StringRBTree) leftRotate(node *StringNode) {
	// original right child
	rightChild := node.right
	if rightChild == nil {
		// if rightChild is nil we cannot rotate
		return
	}
	// set parent reference to node's right child
	isLeft, hasParent := node.isLeftChild()
	if hasParent {
		if isLeft {
			node.parent.setLeft(rightChild)
		} else {
			node.parent.setRight(rightChild)
		}
	} else {
		s.setRoot(rightChild)
	}
	// set node's right child to orig right child's left child
	node.setRight(rightChild.left)
	// set orig right child's left child to node
	rightChild.setLeft(node)
}

// rightRotate rotates a node and its children to the right
// and updates its parent
func (s *StringRBTree) rightRotate(node *StringNode) {
	// original left child
	leftChild := node.left
	if leftChild == nil {
		// if leftChild is nil we cannot rotate
		return
	}
	// set parent reference to node's left child
	isLeft, hasParent := node.isLeftChild()
	if hasParent {
		if isLeft {
			node.parent.setLeft(leftChild)
		} else {
			node.parent.setRight(leftChild)
		}
	} else {
		s.setRoot(leftChild)
	}
	// set node's left child to orig left child's right child
	node.setLeft(leftChild.right)
	// set orig left child's right child to node
	leftChild.setRight(node)
}

// setRoot sets rootNode to node.
func (s *StringRBTree) setRoot(node *StringNode) {
	node.detachParent()
	node.isBlack = true
	s.rootNode = node
}

// isBST tests that the tree is a valid binary
// search tree.
func (s *StringRBTree) isBST() bool {
	isBST := true
	var checkBST func(node *StringNode)
	checkBST = func(node *StringNode) {
		if node == nil {
			return
		}
		if node.left != nil && strings.Compare(node.left.value, node.value) != -1 {
			isBST = false
			return
		}
		if node.right != nil && strings.Compare(node.right.value, node.value) != 1 {
			isBST = false
			return
		}
		checkBST(node.right)
		checkBST(node.left)
	}
	checkBST(s.rootNode)
	return isBST
}

// isRedBlackTree tests that the tree is a valid
// redBlackTree
func (s *StringRBTree) isRedBlackTree() bool {
	if !s.rootNode.isBlack {
		return false
	}
	// check that each red node has two black children or nils
	redChildrenBothBlack := true
	sameBlackHeight := true

	var checkRB func(node *StringNode) int
	checkRB = func(node *StringNode) int {
		if node == nil {
			return 1
		}
		// check that left and right black heights are equal
		leftHeight := checkRB(node.left)
		rightHeight := checkRB(node.right)
		if leftHeight != rightHeight {
			sameBlackHeight = false
		}
		if !node.isBlack {
			// check for red without two black children
			if node.left != nil && node.right != nil {
				if !(node.left.isBlack && node.right.isBlack) {
					redChildrenBothBlack = false
				}
			} else if !(node.left == nil && node.right == nil) {
				redChildrenBothBlack = false
			}
			// does not matter if we use leftHeight or rightHeight
			return leftHeight
		} else {
			return leftHeight + 1
		}
	}
	checkRB(s.rootNode)
	return redChildrenBothBlack && sameBlackHeight
}
