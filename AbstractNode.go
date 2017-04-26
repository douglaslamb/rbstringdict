package main

type NodeInterface interface {
	Left() NodeInterface
	Right() NodeInterface
	Parent() NodeInterface
	Compare(NodeInterface) int
}

type AbstractNode struct {
	left   *AbstractNode
	right  *AbstractNode
	parent *AbstractNode
}

func (a *AbstractNode) Left() *AbstractNode {
	return a.left
}

func (a *AbstractNode) Right() *AbstractNode {
	return a.right
}

func (a *AbstractNode) Parent() *AbstractNode {
	return a.parent
}
