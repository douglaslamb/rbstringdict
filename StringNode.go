package rbstringdict

type StringNode struct {
	value       string
	left        *StringNode
	right       *StringNode
	parent      *StringNode
	isBlack     bool
	doubleBlack bool
}
