package main

import (
	"errors"
	"fmt"
)

type Node struct {
	Data interface{}
	Next *Node
}

func NewNode(data interface{}) *Node {
	return &Node{
		Data: data,
		Next: nil,
	}
}

type linkedList interface {
	InsertHead(v interface{}) error
	InsertTail(v interface{}) error
	InsertK(k int, v interface{}) error
	Delete(k int)	error
	Len() int
	Search(v interface{}) bool
}

// SList 带头节点的链表, 头指针就是链表的名字，仅仅是一个指针而已
// 头节点是为了操作方便而设立的，放在第一个元素节点之前，其数据没有意义(当然也可以存放链表的长度，用做监视等)
// 首元素也是第一个元素的节点，它是头节点后边的第一个节点
// 头节点不是链表必须的，也可以不要头节点
type SList struct {
	Head *Node
	Length int
}

func NewSingleList() *SList {
	return &SList{
		Head:   NewNode(nil),
		Length: 0,
	}
}

func (s *SList) InsertHead(v interface{}) error {
	return s.InsertK(1, v)
}

func (s *SList) InsertTail(v interface{}) error {
	return s.InsertK(s.Length, v)
}

func (s *SList) InsertK(k int, v interface{}) error {
	if k <= 0 || k > s.Length + 1 {
		return errors.New("插入数据位置越界")
	}
	node := NewNode(v)

	pre := s.Head
	for i := 0; i <= k; i++ {
		if i == k - 1 {				// 在记录点的前面插入元素
			node.Next = pre.Next
			pre.Next = node
			s.Length++
		}
		pre = pre.Next
	}
	return nil
}

func (s *SList) Delete(k int) error {
	if k <= 0 || k > s.Length {
		return errors.New("删除数据位置越界")
	}
	pre := s.Head
	for k > 1 {
		pre = pre.Next
		k--
	}
	pre.Next = pre.Next.Next

	return nil
}

func (s *SList) Len() int {
	return s.Length
}

func (s *SList) Search(v interface{}) bool {
	pre := s.Head
	for pre.Next != nil {
		pre = pre.Next
		if pre.Data == v {
			return true
		}
	}
	return false
}

func (s *SList) PrintList() {
	pre := s.Head	// 头节点不存储任何数据
	for pre.Next != nil {
		pre = pre.Next
		fmt.Printf("%v ", pre.Data)
	}
	fmt.Println()
}

func main()  {
	s := NewSingleList()

	_ = s.InsertHead(1)
	_ = s.InsertHead(2)
	_ = s.InsertHead(3)
	_ = s.InsertHead(4)
	_ = s.InsertHead(5)
	_ = s.InsertHead(6)
	s.PrintList()			// 6 5 4 3 2 1
	fmt.Println(s.Len())	// 6

	_ = s.Delete(3)
	s.PrintList()	//6 5 3 2 1

	fmt.Println(s.Search(4))	//false
	fmt.Println(s.Search(5)) //true

}





