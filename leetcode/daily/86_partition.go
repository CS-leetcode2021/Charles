package daily

/**
 *  @ClassName:86_partition
 *  @Description:leetcode 86
 *  @Author:jackey
 *  @Create:2021/8/20 下午7:14
 */

// 100/67
func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 特殊处理
	if head.Val >= x {
		newNode := &ListNode{-1,head}
		return Help(newNode,x).Next
	}

	return Help(head,x)
}


func Help(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// find first >= x
	p1 := head.Next
	p2 := head
	for p1 != nil && p1.Val < x {
		p1 = p1.Next
		p2 = p2.Next
	}

	if p1 == nil {
		return head
	}
	tmp := p1
	for p1 != nil {
		if p1.Val < x {
			tmp.Next = p1.Next
			p1.Next = p2.Next
			p2.Next = p1
			p2 = p2.Next
			p1 = tmp.Next

		} else {
			tmp = p1
			p1 = p1.Next
		}
	}

	return head
}
