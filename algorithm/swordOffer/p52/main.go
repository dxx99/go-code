package main

func main() {

}


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val int
	Next *ListNode
}

// hash表
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	hash := make(map[*ListNode]bool)
	for p := headA; p != nil;  p = p.Next {
		hash[p] = true
	}

	for cur := headB; cur != nil; cur = cur.Next {
		if _, ok := hash[cur]; ok {
			return cur
		}
	}
	return nil
}

//todo 双指针解法，不是很清楚
func getIntersectionNodeV2(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pa, pb := headA, headB
	for pa != pb {
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}
		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}
	return pa
}
