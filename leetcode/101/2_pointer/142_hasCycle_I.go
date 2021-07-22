package __pointer

/**
 *  @ClassName:142_hasCycle_I
 *  @Description:easy 快慢指针
 *  @Author:jackey
 *  @Create:2021/7/22 上午10:57
 */

// 快慢指针
// 98/63
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	fast, slow := head.Next, head
	for fast != slow{

		if fast == nil || fast.Next == nil {
			return false
		}
		fast = fast.Next.Next
		slow = slow.Next
	}

	return true
}

// map也可以做