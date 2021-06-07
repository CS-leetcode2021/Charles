package link

/**
 *  @ClassName:reverseList
 *  @Description:TODO
 *  @Author:jackey
 *  @Create:2021/3/17 上午9:08
 */
/*递归实现*/
func reverseList(head *ListNode) *ListNode {
	if head==nil{
		return nil
	}
	if head.Next==nil {
		return head
	}
	last := reverseList(head.Next)
	head.Next.Next= head
	head.Next=nil
	return last
}

/*for循环*/
func reverseList2(head *ListNode) *ListNode {
	var newListNode *ListNode
	for head != nil {
		p := head.Next
		head.Next=newListNode
		newListNode=head
		head= p
	}

	return newListNode
}