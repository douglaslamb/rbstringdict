package main

import "strings"

type StringRBTree struct {
	rootNode *StringNode
}

// insert adds a string to the dictionary.
func (s *StringRBTree) insert(key string) {
	if key == "" {
		return
	}
	node := s.insertBST(key)
	if node != nil {
		s.insertFixup(node)
	}
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

// insertFixup corrects red-black tree violations resulting
// from insertBST.
func (s *StringRBTree) insertFixup(node *StringNode) {
	for node.parent != nil && !node.parent.isBlack {
		// case 1
		uncle := node.uncle()
		if uncle != nil && !uncle.isBlack {
			node.parent.parent.isBlack = false
			node.parent.isBlack = true
			uncle.isBlack = true
			node = node.parent.parent
		} else {
			// case 2
			if node == node.parent.right && node.parent == node.parent.parent.left {
				node = node.parent
				s.leftRotate(node)
			} else if node == node.parent.left && node.parent == node.parent.parent.right {
				node = node.parent
				s.rightRotate(node)
			}
			// case 3
			node.parent.isBlack = true
			node.parent.parent.isBlack = false
			if node == node.parent.left {
				s.rightRotate(node.parent.parent)
			} else {
				s.leftRotate(node.parent.parent)
			}
		}
	}
	s.rootNode.isBlack = true
}

// remove removes a key from the dictionary.
func (s *StringRBTree) remove(key string) {
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
			if currNode.left != nil && currNode.right != nil {
				// two children
				successor := s.successor(currNode)
				currNode.value = successor.value
				s.easyRemove(successor)
			} else {
				s.easyRemove(currNode)
			}
			currNode = nil
		}
	}
}

// easyRemove deletes the given node as long as
// the node has less than two children. Otherwise
// easyRemove does nothing.
func (s *StringRBTree) easyRemove(node *StringNode) {
	if node.left != nil && node.right != nil {
		return
	}
	var replacer *StringNode
	if node.left != nil {
		replacer = node.left
	} else if node.right != nil {
		replacer = node.right
	}
	if node.parent != nil {
		if node.parent.left == node {
			node.parent.setLeft(replacer)
		} else {
			node.parent.setRight(replacer)
		}
	} else {
		s.setRoot(replacer)
	}
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
	if node.parent != nil {
		if node.parent.left == node {
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
	if node.parent != nil {
		if node.parent.left == node {
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
	if node == nil {
		s.rootNode = nil
	} else {
		node.detachParent()
		node.isBlack = true
		s.rootNode = node
	}
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
	if s.rootNode == nil {
		return true
	}
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

// successor returns a node's in-order successor.
// It returns nil if non exists.
func (s *StringRBTree) successor(node *StringNode) *StringNode {
	if node == nil || node.right == nil {
		return nil
	}
	currNode := node.right
	for currNode.left != nil {
		currNode = currNode.left
	}
	return currNode
}
