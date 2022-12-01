package main

func main() {

}

type MyQueue struct {
	inStack []int
	outStack []int

}

func Constructor() MyQueue {
	return MyQueue{
		inStack: make([]int, 0),
		outStack: make([]int, 0),
	}
}

func (q *MyQueue) Push(x int)  {
	q.inStack = append(q.inStack, x)
}

func (q *MyQueue) in2out() {
	for len(q.inStack) > 0 {
		q.outStack = append(q.outStack, q.inStack[len(q.inStack)-1])
		q.inStack = q.inStack[:len(q.inStack)-1]
	}
}

func (q *MyQueue) Pop() int {
	if len(q.outStack) == 0 {
		q.in2out()
	}

	x := q.outStack[len(q.outStack)-1]
	q.outStack = q.outStack[:len(q.outStack)-1]
	return x
}

func (q *MyQueue) Peek() int {
	if len(q.outStack) == 0 {
		q.in2out()
	}

	x := q.outStack[len(q.outStack)-1]
	return x
}

func (q *MyQueue) Empty() bool {
	return len(q.inStack) == 0 && len(q.outStack) == 0
}
