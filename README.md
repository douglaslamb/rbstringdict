# rbstringdict
A red-black tree dictionary written in Go. It is compatible only with strings. 

A red-black tree is a binary search tree that balances itself. That is, after insertion and deletion it rearranges nodes to balance the tree.

## Install
`go get github.com/douglaslamb/rbstringdict`

## Usage
In your source code at the top of the file:
```
import "rbstringdict"
```
### API

#### `NewStringRBTree() *StringRBTree`
Returns a pointer to an empty StringRBTree.

#### `StringRBTree.Insert(key string)`
Insert a key into the dictionary.

#### `StringRBTree.Remove(key string)`
Remove a key from the dictionary.

#### `StringRBTree.Contains(key string) bool`
Returns a boolean indicating whether the key is in the dictionary.

## Example
```
// create a dictionary
dict := rbstringdict.NewStringRBTree()

// insert some keys
dict.Insert("foo")
dict.Insert("bar")

// remove a key
dict.Remove("bar")

// query the dictionary
fmt.Println(dict.Contains("bar")) // false
fmt.Println(dict.Contains("foo")) // true!
```
