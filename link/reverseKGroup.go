package link

import _ "encoding/hex"

/**
 *  @ClassName:reverseKGroup
 *  @Description:TODO
 *  @Author:jackey
 *  @Create:2021/3/17 上午9:52
 */

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head==nil {
		return  nil
	}

	a := head
	b := head

	for i := 0; i < k; i++ {
		if b == nil {
			return head
		}
		b = b.Next
	}

	newListNode := reverseAB(a,b)
	a.Next= reverseKGroup(b,k)

	return newListNode


}

//反转固定头节点的链表
func reverse(head *ListNode) *ListNode {
	pre := new(ListNode)
	curr := head
	next := head

	for curr!=nil {
		next=curr.Next
		curr.Next=pre
		pre = curr
		curr=next
	}
	return pre

}

func reverseAB(a *ListNode,b *ListNode)  *ListNode{
	pre := new(ListNode)
	curr := a
	next := a

	/*循环的终止条件修改一下*/
	/*反转区间【a，b）的元素，注意是左闭右开*/
	for curr!=b{
		next = curr.Next
		curr.Next= pre
		pre = curr
		curr = next
	}
	/*反转后的头结点*/
	return pre

}