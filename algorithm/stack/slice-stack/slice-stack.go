package main

import (
	"errors"
	"fmt"
)

type Element interface {}

type stacker interface {
	IsEmpty() bool
	Len() int
	Push(e Element) error	// 入栈
	Pop() (Element, error)	// 出栈
	Clear() error   		// 清空栈
	PrintStack()			// 打印栈
}

type Stack struct {
	list []Element
}

func NewStack() *Stack {
	return &Stack{list: make([]Element, 0)}
}

func (s *Stack) IsEmpty() bool {
	if len(s.list) == 0 {
		return true
	}
	return false
}

func (s *Stack) Len() int {
	return len(s.list)
}

func (s *Stack) Push(e Element) error {
	s.list = append(s.list, e)
	return nil
}

func (s *Stack) Pop() (Element, error) {
	if s.Len() < 1 {
		return nil, errors.New("栈为空，没有元素")
	}
	e := s.list[s.Len()-1]
	s.list = s.list[:s.Len()-1]
	return e, nil
}

func (s *Stack) Clear() error {
	if s.Len() == 0 {
		return nil
	}
	s.list = make([]Element, 0)
	return nil
}

func (s *Stack) PrintStack() {
	fmt.Println(s.list)
}

// 顺序栈
func main() {
	s := NewStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.PrintStack()
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	s.Push(1)
	s.Push(1)
	s.Push(1)
	s.PrintStack()
	s.Clear()	//清空栈
	s.PrintStack()

}
