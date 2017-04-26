package main

import "strings"

type StringRBTree struct {
	rootNode *StringNode
}

// insert adds a string to the dictionary.
func (s *StringRBTree) insert(key string) {
	_ = s.insertBST(key)
}

// insertBST inserts the provided key in the dictionary.
// If the key has already been inserted insertBST returns nil.
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

func (s *StringRBTree) remove(key string) {
}

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

func (s *StringRBTree) isEmpty() bool {
	return s.rootNode == nil
}
