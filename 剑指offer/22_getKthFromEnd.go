package main

/**
 *  @ClassName:22_getKthFromEnd
 *  @Description:剑指offer-22 链表中倒数第k个节点
 *  @Author:jackey
 *  @Create:2021/7/13 下午6:46
 */

// 双指针 100/100
func getKthFromEnd(head *ListNode, k int) *ListNode {
	if head == nil || k == 0{
		return nil
	}

	p := head
	q := head

	for i := 0; i < k-1; i++ {
		if p.Next != nil {
			p = p.Next
		}else {
			return nil
		}
	}

	for p.Next != nil {
		p = p.Next
		q = q.Next
	}

	return q

}
