package main

/**
 *  @ClassName:18_deleteNode
 *  @Description:剑指offer-18 删除链表的节点
 *  @Author:jackey
 *  @Create:2021/7/13 下午5:07
 */

func deleteNode(head *ListNode, val int) *ListNode {

	head.DeList(val)

	return head
}

func deleteNode2(head *ListNode, val int) *ListNode {
	if head == nil {
		 return nil
	}

	if head.Val == val{
		return head.Next
	}


	p := head
	q := p.Next

	for q != nil {
		if q.Val == val {
			break
		}
		p = q
		q = q.Next
	}

	p.Next = q.Next

	return head
}


// 双百
func deleteNode3(head *ListNode, val int) *ListNode {
	if head == nil  {
		return head
	}

	if head.Val == val {
		return head.Next
	}

	p := head

	for p.Next != nil {
		if p.Next.Val == val {
			p.Next = p.Next.Next
			break
		}
		p = p.Next
	}


	return head
}