package _1

/**
 *  @ClassName:5_576_middleNode
 *  @Description:寻找链表的中间结点
 *  @Author:jackey
 *  @Create:2021/7/26 上午11:54
 */

// 三种方法
// 1、额外的数组
// 2、两次遍历
// 3、快慢指针

func middleNode(head *ListNode) *ListNode {
	first, second := head, head

	for first != nil && first.Next != nil {
		second = second.Next
		first = first.Next.Next
	}

	return second
}

func middleNode2(head *ListNode) *ListNode {
	p := head
	count := 0

	for p != nil {
		count++
		p = p.Next
	}

	p = head
	i := 0
	for i < count>>1 {
		p = p.Next
		i++
	}
	return p
}
