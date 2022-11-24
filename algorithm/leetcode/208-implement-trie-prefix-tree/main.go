package main

import "fmt"

func main() {
	obj := Constructor()
	obj.Insert("apple")
	fmt.Println(obj.Search("apple"))

}

type nodeType uint8

const (
	root nodeType = iota + 1
	prefix
)

type Trie struct {
	root *node
}

type node struct {
	path     string
	nType    nodeType
	children []*node
}

func Constructor() Trie {
	return Trie{
		root: &node{
			path:     "",
			nType:    root,
			children: []*node{},
		},
	}
}

func (t *Trie) Insert(word string) {

	cur := t.root

	nType := prefix
	for i := 0; i < len(word); i++ {
		if i == len(word)-1 {
			nType = root
		}

		n, ok := isFind(cur, string(word[i]))
		if !ok {
			newNode := &node{
				path:     string(word[i]),
				nType:    nType,
				children: []*node{},
			}
			cur.children = append(cur.children, newNode)
			cur = newNode
		} else {
			cur = n
		}
	}
}

func isFind(cur *node, p string) (*node, bool) {
	for _, child := range cur.children {
		if child.path == p {
			return child, true
		}
	}
	return nil, false
}

func (t *Trie) Search(word string) bool {
	cur := t.root
	var ans *node
	for _, w := range word {
		ans, ok := isFind(cur, string(w))
		if ok {
			cur = ans
		} else {
			return false
		}
	}
	return (*ans).nType == root

}

func (t *Trie) StartsWith(prefix string) bool {
	cur := t.root
	for _, w := range prefix {
		ans, ok := isFind(cur, string(w))
		if ok {
			cur = ans
		} else {
			return false
		}
	}
	return true
}
