package main

//给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
//
// 请你将两个数相加，并以相同形式返回一个表示和的链表。
//
// 你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
//
//
//
// 示例 1：
//
//
//输入：l1 = [2,4,3], l2 = [5,6,4]
//输出：[7,0,8]
//解释：342 + 465 = 807.
//
//
// 示例 2：
//
//
//输入：l1 = [0], l2 = [0]
//输出：[0]
//
//
// 示例 3：
//
//
//输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
//输出：[8,9,9,9,0,0,0,1]
//
//
//
//
// 提示：
//
//
// 每个链表中的节点数在范围 [1, 100] 内
// 0 <= Node.val <= 9
// 题目数据保证列表表示的数字不含前导零
//
// Related Topics 递归 链表 数学 👍 6712 👎 0

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

func addTwoNumbers(l1, l2 *ListNode) (head *ListNode) {
	// 这里作为移动head的游标
	var tail *ListNode
	// 这个表示个位
	carry := 0
	// l1和l2都没走到末尾
	for l1 != nil || l2 != nil {
		// 存放 l1 和 l2 的值， 默认为0(走到末尾了)
		n1, n2 := 0, 0
		// 回去l1的值，l1后移
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		// 获取l2的值，l2后移
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		// 计算当前节点和上一个节点的和
		sum := n1 + n2 + carry
		// sum存放10位， carry存放个位(carry留给下一个节点去累加！！)
		sum, carry = sum%10, sum/10
		// 头节点
		if head == nil {
			head = &ListNode{Val: sum}
			tail = head
		} else {
			// 添加新节点
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
	}
	// 如果个位大于零， 表示还有节点没进去， 加进去
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return
}


//leetcode submit region end(Prohibit modification and deletion)
