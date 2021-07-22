package link

/**
 *  @ClassName:isPalidrome
 *  @Description:TODO
 *  @Author:jackey
 *  @Create:2021/3/17 上午11:22
 */

/*1、时间复杂度太大，空间复杂度哦也太大*/
func isPalindrome(head *ListNode) bool {
	a := []int{}
	p := head
	for p!=nil {
		a = append(a ,p.Val)
		p= p.Next
	}
	b:= reversePalindrome(head)
	for i:= 0;b != nil;i++ {
		if a[i]==b.Val {
			b=b.Next
		}else {
			return false
		}
	}
	return true
}

func reversePalindrome(head *ListNode) *ListNode {
	if head.Next==nil {
		return head
	}

	last := reversePalindrome(head.Next)

	head.Next.Next= head
	head.Next=nil
	return last
}

/*2、比第一种时间复杂度和空间复杂度都少一下，但还是很大*/
func isPalindrome2(head *ListNode) bool {
	a := []int{}
	p := head
	for p!=nil {
		a = append(a ,p.Val)
		p= p.Next
	}
	for i, j := 0, len(a)-1; i < j; {
		if a[i]==a[j] {
			i++
			j--
		}else {
			return false
		}
	}

	return true
}

/*3、递归查询*/
func isPalindrome3(head *ListNode) bool {
	frontPointer := head
	var recursivelyCheck func(*ListNode) bool
	recursivelyCheck = func(curNode *ListNode) bool {
		if curNode != nil {
			if !recursivelyCheck(curNode.Next) {
				return false
			}
			if curNode.Val != frontPointer.Val {
				return false
			}
			frontPointer = frontPointer.Next
		}
		return true
	}
	return recursivelyCheck(head)
}


/*4、空间复杂度为O(1)*/
/*思路，原地将后半段的链表逆序*/
/*1、找到前半部分链表的尾节点
 *2、反转后半部分的链表
 *3、判断是否是回文，使用双指针遍历两个链表
 *4、恢复链表
 *5、返回结果
 */
func reverseListP(head *ListNode) *ListNode {
	/*使用迭代反转*/
	var prev, cur *ListNode = nil, head
	for cur != nil {
		nextTmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = nextTmp
	}
	return prev
}

func endOfFirstHalf(head *ListNode) *ListNode {
	fast := head
	slow := head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func isPalindrome4(head *ListNode) bool {
	if head == nil {
		return true
	}

	// 找到前半部分链表的尾节点并反转后半部分链表
	firstHalfEnd := endOfFirstHalf(head)
	/*反转后半部分的链表*/
	secondHalfStart := reverseListP(firstHalfEnd.Next)

	// 判断是否回文
	p1 := head
	p2 := secondHalfStart
	result := true
	for result && p2 != nil {
		if p1.Val != p2.Val {
			result = false
		}
		p1 = p1.Next
		p2 = p2.Next
	}

	// 还原链表并返回结果
	firstHalfEnd.Next = reverseList(secondHalfStart)
	return result
}

/*5、可以借助二叉树后序遍历的思路，不需要显式反转链表也可以倒序遍历链表*/
var left *ListNode
/*算法复杂度和空间复杂度*/
func isPalindrome5(head *ListNode) bool{
	left =head
	return traverse(head)
}

/*链表的后序遍历*/
func traverse(right *ListNode) bool {
	if right==nil {
		return true
	}

	res := traverse(right.Next)

	res = res && right.Val==left.Val

	left= left.Next
	return res
}

