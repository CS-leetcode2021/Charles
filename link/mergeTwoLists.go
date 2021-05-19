package link

/**
 *  @ClassName:mergeTwoLists
 *  @Description: leetcode 21 合并两个有序链表
 *  @Author:jackey
 *  @Create:2021/4/2 下午6:48
 */

func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	result := &ListNode{}

	root := result

	for l1 != nil && l2!= nil {
		if l1.Val>l2.Val {
			root.Next=l2
			l2=l2.Next
		}else{
			root.Next=l1
			l1=l1.Next
		}
		root = root.Next
	}

	if l1 == nil {
		root.Next=l2
	}else {
		root.Next=l1
	}

	return result.Next
}
