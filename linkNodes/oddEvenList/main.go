package oddEvenList

//给定一个单链表，把所有的奇数节点和偶数节点分别排在一起。请注意，这里的奇数节点和偶数节点指的是节点编号的奇偶性，而不是节点的值的奇偶性。
//
// 请尝试使用原地算法完成。你的算法的空间复杂度应为 O(1)，时间复杂度应为 O(nodes)，nodes 为节点总数。
//
// 示例 1:
//
// 输入: 1->2->3->4->5->NULL
//输出: 1->3->5->2->4->NULL
//
//
// 示例 2:
//
// 输入: 2->1->3->5->6->4->7->NULL
//输出: 2->3->6->7->1->5->4->NULL
//
// 说明:
//
//
// 应当保持奇数节点和偶数节点的相对顺序。
// 链表的第一个节点视为奇数节点，第二个节点视为偶数节点，以此类推。
//
// Related Topics 链表 👍 466 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	// even 指向偶数， odd 指向奇数
	evenHead := head.Next
	odd := head
	even := evenHead
	for even != nil && even.Next != nil {
		// 分裂链表
		// 偶数在奇数后面，奇数的next指向偶数当前的next，相当于跳了两步
		odd.Next = even.Next
		// 跟新奇数的位置， 此时奇数在偶数的后面
		odd = odd.Next
		// 偶数的next指向奇数的next，相当于偶数走两步
		even.Next = odd.Next
		// 更新偶数的位置
		even = even.Next
	}
	// 偶数拼接到奇数后面
	odd.Next = evenHead
	return head
}

//leetcode submit region end(Prohibit modification and deletion)
