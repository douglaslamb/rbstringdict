package rbstringdict

import "strings"

type StringRBTree struct {
	rootNode *StringNode
	dummy    *StringNode
}

// NewStringRBTree() returns a new StringRBTree
// with an initialized nil dummy
func NewStringRBTree() *StringRBTree {
	tree := &StringRBTree{}
	tree.dummy = &StringNode{}
	tree.dummy.isBlack = true
	tree.setRoot(tree.dummy)
	return tree
}

// newStringNode() returns a new StringNode
// with two dummy children
func (s *StringRBTree) newStringNode() *StringNode {
	node := &StringNode{}
	node.left = s.dummy
	node.right = s.dummy
	node.parent = s.dummy
	return node
}

// Insert adds a string to the dictionary.
func (s *StringRBTree) Insert(key string) {
	if key == "" {
		return
	}
	node := s.insertBST(key)
	if node != s.dummy {
		s.insertFixup(node)
	}
}

// insertBST inserts the provided key in the dictionary and returns
// the added node. If the key has already been inserted
// insertBST returns dummy.
func (s *StringRBTree) insertBST(key string) *StringNode {
	// BST insert
	if s.isEmpty() {
		node := s.newStringNode()
		node.value = key
		s.rootNode = node
		return s.rootNode
	}
	currNode := s.rootNode
	for {
		comparison := strings.Compare(key, currNode.value)
		if comparison < 0 {
			if currNode.left == s.dummy {
				node := s.newStringNode()
				node.value = key
				currNode.left = node
				node.parent = currNode
				return node
			} else {
				currNode = currNode.left
			}
		} else if comparison > 0 {
			if currNode.right == s.dummy {
				node := s.newStringNode()
				node.value = key
				currNode.right = node
				node.parent = currNode
				return node
			} else {
				currNode = currNode.right
			}
		} else {
			// key already existed in tree
			return s.dummy
		}
	}
}

// insertFixup corrects red-black tree violations resulting
// from insertBST.
func (s *StringRBTree) insertFixup(node *StringNode) {
	for node.parent != s.dummy && !node.parent.isBlack {
		// case 1
		uncle := s.uncle(node)
		if uncle != s.dummy && !uncle.isBlack {
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

// Remove removes a key from the dictionary.
func (s *StringRBTree) Remove(key string) {
	// BST remove
	var extraBlack *StringNode
	var parent *StringNode
	currNode := s.rootNode
	for currNode != s.dummy {
		comparison := strings.Compare(key, currNode.value)
		if comparison < 0 {
			currNode = currNode.left
		} else if comparison > 0 {
			currNode = currNode.right
		} else {
			// found key
			if currNode.left != s.dummy && currNode.right != s.dummy {
				// two children
				successor := s.successor(currNode)
				currNode.value = successor.value
				currNode = successor
			}
			// remeber parent if we need to do fixup
			parent = currNode.parent
			// one or zero children delete procedure
			replacer := s.dummy
			if currNode.left != s.dummy {
				replacer = currNode.left
			} else if currNode.right != s.dummy {
				replacer = currNode.right
			}
			if currNode.parent != s.dummy {
				if currNode.parent.left == currNode {
					s.setLeft(currNode.parent, replacer)
				} else {
					s.setRight(currNode.parent, replacer)
				}
			} else {
				s.setRoot(replacer)
			}
			if currNode.isBlack {
				// set to trigger fixup remove
				extraBlack = replacer
			}
			currNode = s.dummy
		}
	}
	if extraBlack != nil {
		for extraBlack != s.rootNode && extraBlack.isBlack {
			if extraBlack == parent.left {
				sibling := parent.right
				if !sibling.isBlack {
					sibling.isBlack = true
					parent.isBlack = false
					s.leftRotate(parent)
					sibling = parent.left
				}
				if sibling.left.isBlack && sibling.right.isBlack {
					sibling.isBlack = false
					extraBlack = parent
				} else {
					if sibling.right.isBlack {
						sibling.left.isBlack = true
						sibling.isBlack = false
						s.rightRotate(sibling)
						sibling = parent.right
					}
					sibling.isBlack = parent.isBlack
					parent.isBlack = true
					sibling.right.isBlack = true
					s.leftRotate(parent)
					extraBlack = s.rootNode
				}
			} else {
				sibling := parent.left
				if !sibling.isBlack {
					sibling.isBlack = true
					parent.isBlack = false
					s.rightRotate(parent)
					sibling = parent.right
				}
				if sibling.left.isBlack && sibling.right.isBlack {
					sibling.isBlack = false
					extraBlack = parent
				} else {
					if sibling.left.isBlack {
						sibling.right.isBlack = true
						sibling.isBlack = false
						s.leftRotate(sibling)
						sibling = parent.left
					}
					sibling.isBlack = parent.isBlack
					parent.isBlack = true
					sibling.left.isBlack = true
					s.rightRotate(parent)
					extraBlack = s.rootNode
				}
			}
			extraBlack.isBlack = true
		}
	}
}

// Contains returns true if the key is in the dictionary
// and false otherwise.
func (s *StringRBTree) Contains(key string) bool {
	if s.isEmpty() {
		return false
	}
	currNode := s.rootNode
	for {
		comparison := strings.Compare(key, currNode.value)
		if comparison < 0 {
			if currNode.left == s.dummy {
				return false
			}
			currNode = currNode.left
		} else if comparison > 0 {
			if currNode.right == s.dummy {
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
	return s.rootNode == s.dummy
}

// leftRotate rotates a node and its children to the left
// and updates its parent
func (s *StringRBTree) leftRotate(node *StringNode) {
	// original right child
	rightChild := node.right
	if rightChild == s.dummy {
		// if rightChild is nil we cannot rotate
		return
	}
	// set parent reference to node's right child
	if node.parent != s.dummy {
		if node.parent.left == node {
			s.setLeft(node.parent, rightChild)
		} else {
			s.setRight(node.parent, rightChild)
		}
	} else {
		s.setRoot(rightChild)
	}
	// set node's right child to orig right child's left child
	s.setRight(node, rightChild.left)
	// set orig right child's left child to node
	s.setLeft(rightChild, node)
}

// rightRotate rotates a node and its children to the right
// and updates its parent
func (s *StringRBTree) rightRotate(node *StringNode) {
	// original left child
	leftChild := node.left
	if leftChild == s.dummy {
		// if leftChild is nil we cannot rotate
		return
	}
	// set parent reference to node's left child
	if node.parent != s.dummy {
		if node.parent.left == node {
			s.setLeft(node.parent, leftChild)
		} else {
			s.setRight(node.parent, leftChild)
		}
	} else {
		s.setRoot(leftChild)
	}
	// set node's left child to orig left child's right child
	s.setLeft(node, leftChild.right)
	// set orig left child's right child to node
	s.setRight(leftChild, node)
}

// setRoot sets rootNode to node.
func (s *StringRBTree) setRoot(node *StringNode) {
	if node == s.dummy {
		s.rootNode = s.dummy
	} else {
		s.detachParent(node)
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
		if node == s.dummy {
			return
		}
		if node.left != s.dummy && strings.Compare(node.left.value, node.value) != -1 {
			isBST = false
			return
		}
		if node.right != s.dummy && strings.Compare(node.right.value, node.value) != 1 {
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
	if s.rootNode == s.dummy {
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
		if node == s.dummy {
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
			if node.left != s.dummy && node.right != s.dummy {
				if !(node.left.isBlack && node.right.isBlack) {
					redChildrenBothBlack = false
				}
			} else if !(node.left == s.dummy && node.right == s.dummy) {
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
	if node == s.dummy || node.right == s.dummy {
		return s.dummy
	}
	currNode := node.right
	for currNode.left != s.dummy {
		currNode = currNode.left
	}
	return currNode
}

// setLeft sets the node's left child.
func (s *StringRBTree) setLeft(node *StringNode, left *StringNode) {
	// disconnect left child
	if node.left != s.dummy {
		node.left.parent = s.dummy
	}
	if left != s.dummy {
		// if node is already child of another node
		// disconnect it
		s.detachParent(left)
		left.parent = node
	}
	node.left = left
}

// setRight sets the node's right child.
func (s *StringRBTree) setRight(node *StringNode, right *StringNode) {
	// disconnect right child
	if node.right != s.dummy {
		node.right.parent = s.dummy
	}
	if right != s.dummy {
		// if node is already child of another node
		// disconnect it
		s.detachParent(right)
		right.parent = node
	}
	node.right = right
}

// detachParent calls setLeft or setRight on
// node's parent if any. setLeft or setRight
// set pointers accordingly.
func (s *StringRBTree) detachParent(node *StringNode) {
	if node.parent != s.dummy {
		if node.parent.left == node {
			s.setLeft(node.parent, s.dummy)
		} else {
			s.setRight(node.parent, s.dummy)
		}
	}
}

// uncle returns the node's uncle. It returns
// nil if node has no parent, no grandparent, or no uncle
func (s *StringRBTree) uncle(node *StringNode) *StringNode {
	if node.parent != s.dummy && node.parent.parent != s.dummy {
		if node.parent == node.parent.parent.left {
			return node.parent.parent.right
		} else {
			return node.parent.parent.left
		}
	}
	return s.dummy
}
