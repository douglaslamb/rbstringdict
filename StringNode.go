package main

import "strings"

type StringNode struct {
	value  string
	left   *StringNode
	right  *StringNode
	parent *StringNode
}

func (s *StringNode) compare(node *StringNode) int {
	return strings.Compare(s.value, node.value)
}
