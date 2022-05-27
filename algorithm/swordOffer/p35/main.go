package main

// 请实现 copyRandomList 函数，复制一个复杂链表。在复杂链表中，每个节点除了有一个 next 指针指向下一个节点，还有一个 random 指针指向链表中的任意节点或者 null。
//
// 
//
//示例 1：
//
//
//
//输入：head = [[7,null],[13,0],[11,4],[10,2],[1,0]]
//输出：[[7,null],[13,0],[11,4],[10,2],[1,0]]
//示例 2：
//
//
//
//输入：head = [[1,1],[2,1]]
//输出：[[1,1],[2,1]]
//示例 3：
//
//
//
//输入：head = [[3,null],[3,0],[3,null]]
//输出：[[3,null],[3,0],[3,null]]
//示例 4：
//
//输入：head = []
//输出：[]
//解释：给定的链表为空（空指针），因此返回 null。
// 
//
//提示：
//
//-10000 <= Node.val <= 10000
//Node.random 为空（null）或指向链表中的节点。
//节点数目不超过 1000 。
// 
//
//注意：本题与主站 138 题相同：https://leetcode-cn.com/problems/copy-list-with-random-pointer/
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/fu-za-lian-biao-de-fu-zhi-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {

}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

type Node struct {
	Val int
	Next *Node
	Random *Node
}


func copyRandomList(head *Node) *Node {
	if head == nil {
		 return nil
	}

	// 判断节点是否已经被计算过了，这个时候就需要直接返回，不需要再生产node
	hashNode := make(map[*Node]*Node)

	var deepCopy func(node *Node) *Node
	deepCopy = func(node *Node) *Node {
		if node == nil {
			return nil
		}

		// 在缓存表中已经存在这个节点了
		if n, ok := hashNode[node]; ok {
			return n
		}

		newNode := &Node{Val: node.Val}
		hashNode[node] = newNode
		newNode.Next = deepCopy(node.Next)
		newNode.Random = deepCopy(node.Random)

		return newNode
	}

	return deepCopy(head)
}
