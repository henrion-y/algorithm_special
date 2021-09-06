package main

//给你一个链表的头节点 head ，旋转链表，将链表每个节点向右移动 k 个位置。
//
//
//
// 示例 1：
//
//
//输入：head = [1,2,3,4,5], k = 2
//输出：[4,5,1,2,3]
//
//
// 示例 2：
//
//
//输入：head = [0,1,2], k = 4
//输出：[2,0,1]
//
//
//
//
// 提示：
//
//
// 链表中节点的数目在范围 [0, 500] 内
// -100 <= Node.val <= 100
// 0 <= k <= 2 * 10⁹
//
// Related Topics 链表 双指针 👍 621 👎 0

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

func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil || head.Next == nil {
		return head
	}
	n:=1
	p := head
	for p.Next != nil {
		p = p.Next
		n++
	}
	// todo 这里为什么直接移动k个位置会出错呢！！！
	add := n - k%n
	if add == n {
		return head
	}
	p.Next = head
	for i := 0; i < add; i++ {
		p = p.Next
	}
	ret := p.Next
	p.Next = nil
	return ret

	//if k == 0 || head == nil || head.Next == nil {
	//	return head
	//}
	//n := 1
	//iter := head
	//for iter.Next != nil {
	//	iter = iter.Next
	//	n++
	//}
	//add := n - k%n
	//if add == n {
	//	return head
	//}
	//iter.Next = head
	//for add > 0 {
	//	iter = iter.Next
	//	add--
	//}
	//ret := iter.Next
	//iter.Next = nil
	//return ret

}

//leetcode submit region end(Prohibit modification and deletion)
