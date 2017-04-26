package main

import "testing"

func TestStringNodeCompare(t *testing.T) {
	node := &StringNode{}
	node.value = `f`
	// test lesser
	lesser := &StringNode{}
	lesser.value = `a`
	if node.compare(lesser) != 1 {
		t.Errorf("node.compare(StringNode{`a`}) = %v; want %v", node.compare(lesser), 1)
	}
	// test greater
	greater := &StringNode{}
	greater.value = `z`
	if node.compare(greater) != -1 {
		t.Errorf("node.compare(StringNode{`z`}) = %v; want %v", node.compare(greater), -1)
	}
	// test equal
	equal := &StringNode{}
	equal.value = `f`
	if node.compare(equal) != 0 {
		t.Errorf("node.compare(StringNode{`f`}) = %v; want %v", node.compare(equal), 0)
	}
}
