package _1

/**
 *  @ClassName:10_21_mergeTwoLists
 *  @Description:合并两个有序链表
 *  @Author:jackey
 *  @Create:2021/8/3 下午3:55
 */

// 100/100
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	// nil
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	res := new(ListNode)

	root := res
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			root.Next = l2
			l2 = l2.Next
		} else {
			root.Next = l1
			l1 = l1.Next
		}
		root = root.Next
	}

	if l1 == nil {
		root.Next = l2
	} else {
		root.Next = l1
	}

	return res.Next

}
