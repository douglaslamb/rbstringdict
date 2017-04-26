package main

import "testing"

func TestInsert(t *testing.T) {
	tree := StringRBTree{}
	tree.insert("foo")
	if !tree.contains("foo") {
		t.Errorf("tree.contains(\"foo\") = %v; want %v", tree.contains("foo"), true)
	}
}

func TestRemove(t *testing.T) {
	tree := StringRBTree{}
	tree.insert("foo")
	tree.remove("foo")
	if tree.contains("foo") {
		t.Errorf("tree.contains(\"foo\") = %v; want %v", tree.contains("foo"), false)
	}
}
