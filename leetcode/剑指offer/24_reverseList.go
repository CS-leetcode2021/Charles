package main

/**
 *  @ClassName:24_offer
 *  @Description:剑指offer-24 反转链表
 *  @Author:jackey
 *  @Create:2021/7/13 下午7:00
 */

// 3个节点指针
func reverseList(head *ListNode) *ListNode {
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

// 递归
func reverseList2(head *ListNode) *ListNode {
	if head==nil{
		return nil
	}

	if head.Next==nil{
		return head
	}

	last := reverseList(head.Next)

	head.Next.Next=head
	head.Next=nil
	return last
}
