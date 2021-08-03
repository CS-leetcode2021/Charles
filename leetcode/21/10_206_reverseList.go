package _1

/**
 *  @ClassName:10_206_reverseList
 *  @Description:反转链表
 *  @Author:jackey
 *  @Create:2021/8/3 下午4:01
 */


// 100/O(n)
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next  == nil {
		return head
	}

	post :=reverseList(head.Next)

	head.Next.Next = head
	head.Next = nil

	return  post
}

// 双指针遍历  100/O(1)空间复杂度
func reverseList2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	r := head
	p := head.Next
	q := r


	for p != nil {
		r.Next = p.Next
		p.Next = q
		q = p
		p = r.Next
	}


	return q
}