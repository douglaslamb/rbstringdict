package main

type StringNode struct {
	value  string
	left   *StringNode
	right  *StringNode
	parent *StringNode
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

// setLeft sets the nodes left child.
func (s *StringNode) setLeft(node *StringNode) {
	s.left = node
	node.parent = s
}

// setRight sets the nodes right child.
func (s *StringNode) setRight(node *StringNode) {
	s.right = node
	node.parent = s
}
