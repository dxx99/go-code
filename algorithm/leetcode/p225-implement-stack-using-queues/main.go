package main

func main() {

}

type MyStack struct {
	inQueue []int
	outQueue []int
}


func Constructor() MyStack {
	return MyStack{
		inQueue:  make([]int, 0),
		outQueue: make([]int, 0),
	}
}


func (m *MyStack) Push(x int)  {
	m.inQueue = append(m.inQueue, x)
	for len(m.outQueue) > 0 {
		m.inQueue = append(m.inQueue, m.outQueue[0])
		m.outQueue = m.outQueue[1:]
	}
	m.outQueue, m.inQueue = m.inQueue, m.outQueue
}


func (m *MyStack) Pop() int {
	x := m.outQueue[0]
	m.outQueue = m.outQueue[1:]
	return x
}


func (m *MyStack) Top() int {
	return m.outQueue[0]
}


func (m *MyStack) Empty() bool {
	return len(m.inQueue) == 0 && len(m.outQueue) == 0
}
