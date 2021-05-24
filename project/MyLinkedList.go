package project

/**
 *  @ClassName:MyLinkedList
 *  @Description:707 leetcode 设计链表,有头节点
 *  @Author:jackey
 *  @Create:2021/5/24 下午9:43
 */
type ListNode struct {
	val int
	next *ListNode
}
type MyLinkedList struct {
	size int
	root *ListNode
}

func Constructor() MyLinkedList {
	return MyLinkedList{
		0,
		&ListNode{},
	}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index>this.size {
		return -1
	}

	prev := this.root
	for i := 0; i < index; i++ {
		prev = prev.next
	}

	return prev.next.val
}
func (this *MyLinkedList) AddAtHead(val int) {
	this.AddAtIndex(0, val)
}

func (this *MyLinkedList) AddAtTail(val int) {
	this.AddAtIndex(this.size, val)
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index > this.size {
		return
	}

	prev := this.root
	for i := 0; i < index; i++ {
		prev = prev.next
	}
	node := &ListNode{val,nil}
	node.next = prev.next
	prev.next = node

	this.size++
}

func (this MyLinkedList) DeleteAtIndex(index int) {
	if index <0 || index > this.size {
		return
	}

	prev := this.root
	for i := 0; i < index; i++ {
		prev = prev.next
	}

	prev.next = prev.next.next
	this.size--


}