package link

/**
 *  @ClassName:reverseBetween
 *  @Description:TODO
 *  @Author:jackey
 *  @Create:2021/3/17 上午9:34
 */


var successor *ListNode

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head==nil{
		return nil
	}
	if left==1 {
		return reverseN(head,right)
	}

	head.Next= reverseBetween(head.Next,left-1,right-1)
	return head

}

func reverseN(head *ListNode, m int) *ListNode {
	if m == 1 {
		/*记录n+1的节点*/
		successor = head.Next
		return head
	}


	/*返回第N个节点*/
	last := reverseN(head.Next, m-1)
	head.Next.Next = head
	/*指向N+1个节点*/
	head.Next = successor

	return last

}
