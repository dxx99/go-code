package main

func main() {

}

type Trie struct {
	children [26]*Trie
	isEnd bool
}


func Constructor() Trie {
	return Trie{}
}


// Insert 迭代插入数据
func (t *Trie) Insert(word string)  {
	node := t	// 存储当前节点，循环迭代
	for _, ch := range word {
		ch -= 'a'
		if node.children[ch] == nil {
			node.children[ch] = &Trie{}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

// 找到索引节点
func (t *Trie) searchPrefix(prefix string) *Trie {
	node := t
	for _, ch := range prefix {
		ch -= 'a'
		if node.children[ch] == nil {
			return nil
		}
		node = node.children[ch]
	}
	return node
}


func (t *Trie) Search(word string) bool {
	node := t.searchPrefix(word)
	if node != nil {
		return node.isEnd
	}
	return false
}


func (t *Trie) StartsWith(prefix string) bool {
	return t.searchPrefix(prefix) == nil
}
