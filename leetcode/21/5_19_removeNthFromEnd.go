package _1

/**
 *  @ClassName:5_19_removeNthFromEnd
 *  @Description:删除链表中倒数第N个节点
 *  @Author:jackey
 *  @Create:2021/7/26 上午11:55
 */

// 快慢指针
// 100/100
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	res := &ListNode{0,head}
	first, second := head, res
	for i := 0; i < n; i++ {
		first = first.Next
	}
	for first != nil {
		first = first.Next
		second = second.Next
	}
	second.Next = second.Next.Next
	return res.Next
}

// 使用栈结构
func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	nodes := []*ListNode{}
	dummy := &ListNode{0, head}
	for node := dummy; node != nil; node = node.Next {
		nodes = append(nodes, node)
	}
	prev := nodes[len(nodes)-1-n]
	prev.Next = prev.Next.Next
	return dummy.Next
}
