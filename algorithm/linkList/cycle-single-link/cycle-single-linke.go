package main

import (
	"errors"
	"fmt"
)

type Node struct {
	data interface{}
	next *Node
}

func NewNode(data interface{}) *Node {
	return &Node{
		data: data,
		next: nil,
	}
}

// 定义链表相关接口
type linkedList interface {
	InsertHead(v interface{}) error
	InsertTail(v interface{}) error
	InsertK(k int, v interface{}) error
	Delete(k int)	error
	Len() int
	Search(v interface{}) (int, error)
}

// CycleSingleList 环形链表所有的节点形成一个环
// 单链表的尾节点指向首节点形成单向循环链表
type CycleSingleList struct {
	head *Node
	height int
}

func NewCycleSingleList() *CycleSingleList {
	head := NewNode(nil)
	head.next = head
	return &CycleSingleList{
		head:   head,
		height: 0,
	}
}

func (c *CycleSingleList) InsertHead(v interface{}) error {
	return c.InsertK(1, v)
}

func (c *CycleSingleList) InsertTail(v interface{}) error {
	return c.InsertK(c.Len()+1, v)
}

func (c *CycleSingleList) InsertK(k int, v interface{}) error {
	if k < 1 || k > c.Len() + 1 {
		return errors.New("插入数据位置越界")
	}
	p := c.head
	for k > 1 {
		p = p.next
		k--
	}

	node := NewNode(v)
	node.next = p.next
	p.next = node
	c.height++

	return nil
}

func (c *CycleSingleList) Delete(k int) error {
	if k <= 0 || k > c.Len() {
		return errors.New("删除数据位置越界")
	}
	pre := c.head
	for k > 1 {
		pre = pre.next
		k--
	}
	pre.next = pre.next.next
	return nil
}

func (c *CycleSingleList) Len() int {
	return c.height
}

func (c *CycleSingleList) Search(v interface{}) (int, error) {
	pre := c.head
	k := 0
	for pre.next != nil {
		if k > c.Len() {
			return -1, errors.New("not found")
		}
		pre = pre.next
		k++
		if pre.data == v {
			return k, nil
		}
	}
	return -1, errors.New("not found")
}

func (c *CycleSingleList) PrintList() {
	pre := c.head	// 头节点不存储任何数据
	l := c.Len()
	for pre.next != nil && l > 0 {
		l--
		pre = pre.next
		fmt.Printf("%v ", pre.data)
	}
	fmt.Println()
}



// 循环单链表
func main() {

	s := NewCycleSingleList()

	_ = s.InsertHead(1)
	_ = s.InsertHead(2)
	_ = s.InsertHead(3)
	_ = s.InsertHead(4)
	_ = s.InsertHead(5)
	_ = s.InsertHead(6)
	_ = s.InsertTail(7)
	s.PrintList()			// 6 5 4 3 2 1 7
	fmt.Println(s.Len())	// 6

	_ = s.Delete(3)
	s.PrintList()	//6 5 3 2 1 7

	fmt.Println(s.Search(4))	//-1 not found
	fmt.Println(s.Search(5)) //2 <nil>

}
