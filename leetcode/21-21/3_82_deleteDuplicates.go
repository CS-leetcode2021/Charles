package _1_21

/**
 *  @ClassName:3_82_deleteDuplicates
 *  @Description:删除排序链表中的重复元素
 *  @Author:jackey
 *  @Create:2021/8/7 下午4:32
 */

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := new(ListNode)
	dummy.Next = head

	prev := dummy
	for head != nil && head.Next != nil {
		if head.Val == head.Next.Val {
			for head.Next != nil && head.Val == head.Next.Val {
				head.Next = head.Next.Next
			}
			prev.Next = head.Next
		} else {
			prev = prev.Next
		}
		head = head.Next
	}
	return dummy.Next
}

