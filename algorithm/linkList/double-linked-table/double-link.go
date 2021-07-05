package main

import (
	"errors"
	"fmt"
)

// Node 节点元素的数据
type Node struct {
	data interface{}
	prev *Node
	next *Node
}

func NewNode(data interface{}) *Node {
	return &Node{
		data: data,
		prev: nil,
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

// DList head与tail指针不存储数据，只做一个前端的位置标记
type DList struct {
	head *Node
	tail *Node
	length int
}

func NewDoubleList() *DList {
	head := NewNode(nil)
	tail := NewNode(nil)
	head.next = tail
	tail.prev = head
	return &DList{
		head:   head,
		tail:   tail,
		length: 0,
	}
}

func (d *DList) InsertHead(v interface{}) error {
	return d.InsertK(1, v)
}

func (d *DList) InsertTail(v interface{}) error {
	return d.InsertK(d.length + 1, v)
}

func (d *DList) InsertK(k int, v interface{}) error {
	if k < 1 || k > d.length + 1 {
		return errors.New("插入的元素不在规定的索引位置")
	}
	node := NewNode(v)

	p := d.head
	for k > 1 {
		p = p.next
		k--
	}

	// 画图可以深刻理解
	node.next = p.next		// 1. 插入元素的next指针，是他前面元素的指针
	node.next.prev = node	// 2. 后面一个元素的prev指针，应该指向插入的元素
	p.next = node			// 3. 前面一个元素的next指针，应该指向插入的元素
	node.prev = p			// 4. 插入元素的prev指针，应指向前面的元素
	d.length++
	return nil
}

func (d *DList) Delete(k int) error {
	if k < 1 || k > d.length {
		return errors.New("删除的元素不在规定的索引位置")
	}

	p := d.head
	for k > 1 {
		p = p.next
		k--
	}
	p.next = p.next.next
	p.next.prev = p
	d.length--
	return nil
}

func (d *DList) Len() int {
	return d.length
}

func (d *DList) Search(v interface{}) (int,  error) {
	p := d.head
	k := 0
	if p.next != nil {
		p = p.next
		k++
		if p.data == v {
			return k, nil
		}
	}
	return -1, errors.New("not found")
}

func (d *DList) PrintList() {
	p := d.head.next
	for p.next != nil {
		fmt.Printf("%v ", p.data)
		p = p.next
	}
	fmt.Println()
}

// 双链表实现
func main() {
	l := NewDoubleList()

	l.InsertHead(1)
	l.InsertHead(1)
	l.InsertHead(2)
	l.InsertHead(2)
	l.InsertTail(3)
	l.InsertTail(5)
	l.PrintList()	// 2 2 1 1 3 5
	fmt.Println(l.Len())	//6

}
