package main

type StringNode struct {
	value   string
	left    *StringNode
	right   *StringNode
	parent  *StringNode
	isBlack bool
}

// isLeftChild returns true if the node is
// its parent's left child. It returns false
// if the node is its parent's right child. The
// second return value is false if parent is nil.
func (s *StringNode) isLeftChild() (bool, bool) {
	if s.parent == nil {
		return false, false
	}
	if s.parent.left == s {
		return true, true
	} else {
		return false, true
	}
}

// setLeft sets the node's left child.
func (s *StringNode) setLeft(node *StringNode) {
	// disconnect left child
	if s.left != nil {
		s.left.parent = nil
	}
	if node != nil {
		// if node is already child of another node
		// disconnect it
		node.detachParent()
		node.parent = s
	}
	s.left = node
}

// setRight sets the node's right child.
func (s *StringNode) setRight(node *StringNode) {
	// disconnect right child
	if s.right != nil {
		s.right.parent = nil
	}
	if node != nil {
		// if node is already child of another node
		// disconnect it
		node.detachParent()
		node.parent = s
	}
	s.right = node
}

// detachParent calls setLeft or setRight on
// node's parent if any. setLeft or setRight
// set pointers accordingly.
func (s *StringNode) detachParent() {
	isLeft, hasParent := s.isLeftChild()
	if hasParent {
		if isLeft {
			s.parent.setLeft(nil)
		} else {
			s.parent.setRight(nil)
		}
	}
}

// uncle returns the node's uncle. It returns
// nil if node has no parent, no grandparent, or no uncle
func (s *StringNode) uncle() *StringNode {
	if s.parent != nil {
		isLeft, hasGrandparent := s.parent.isLeftChild()
		if hasGrandparent {
			grandparent := s.parent.parent
			if isLeft {
				return grandparent.right
			} else {
				return grandparent.left
			}
		}
	}
	return nil
}
