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
