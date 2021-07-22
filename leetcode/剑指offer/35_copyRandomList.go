package main

/**
 *  @ClassName:35_copyRandomList
 *  @Description:剑指offer-35 复杂链表的复制 同 leetcode-138
 *  @Author:jackey
 *  @Create:2021/7/19 下午1:06
 */

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// 1、实现原节点的顺序拷贝
// 2、遍历处理所有的random节点信息
// 3、剥离原始节点和新的节点
func copyRandomList(head *Node) *Node {
	//

	if head == nil || head.Next == nil {
		return nil
	}
	p := head
	for p != nil {
		tmp := new(Node)
		tmp.Val = p.Val
		tmp.Next = p.Next
		p.Next = tmp
		p = tmp.Next
	}

	p = head
	for p != nil {
		if p.Random != nil {
			p.Next.Random = p.Random.Next
		}
		p = p.Next.Next
	}

	res := head.Next
	p = res
	oldNode := head
	for p.Next != nil {	// 注意判断空的条件
		oldNode.Next = p.Next
		p.Next = p.Next.Next
		oldNode = oldNode.Next
		p = p.Next
	}
	oldNode.Next = nil
	return res
}
