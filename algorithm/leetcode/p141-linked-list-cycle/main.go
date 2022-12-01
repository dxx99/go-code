package main

func main() {

}

type ListNode struct {
	Val int
	Next *ListNode
}


// hash法
func hasCycleV1(head *ListNode) bool {
	m := map[*ListNode]bool{}
	for head != nil {
		if _, ok := m[head]; ok {
			return true
		} else {
			m[head] = true
		}
		head = head.Next
	}

	return false
}

// 快慢指针法
func hasCycle(head *ListNode) bool {
	fast := head
	slow := head

	for fast != nil && fast.Next != nil && slow != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}
