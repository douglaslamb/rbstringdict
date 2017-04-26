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
	// add to set left and set right that they need to call setleft or setright whichever is appropriate
	// on node's previous parent if it has one
	if s.left != nil {
		s.left.parent = nil
	}
	s.left = node
	if node != nil {
		if node.parent != nil {
			if node.parent.left == node {
				node.parent.setLeft(nil)
			} else {
				node.parent.setRight(nil)
			}
		}
		node.parent = s
	}
}

// setRight sets the node's right child.
func (s *StringNode) setRight(node *StringNode) {
	if s.right != nil {
		s.right.parent = nil
	}
	s.right = node
	if node != nil {
		node.parent = s
	}
}
