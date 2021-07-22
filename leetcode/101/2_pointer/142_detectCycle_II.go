package __pointer

/**
 *  @ClassName:142_detectCycle_II
 *  @Description:medium
 *  @Author:jackey
 *  @Create:2021/7/22 上午11:14
 */

// 97/100 快慢指针
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	fast, slow := head.Next, head // fast 先走一步
	for fast != slow {
		if fast == nil || fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	fast = head
	slow = slow.Next // slow 先走一步
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}

	return fast

}
