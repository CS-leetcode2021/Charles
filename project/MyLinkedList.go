package project

/**
 *  @ClassName:MyLinkedList
 *  @Description:707 leetcode 设计链表,有头节点
 *  @Author:jackey
 *  @Create:2021/5/24 下午9:43
 */
type ListNode struct {
	val  int
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
	if index < 0 || index >= this.size {
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
	node := &ListNode{val, nil}
	node.next = prev.next
	prev.next = node

	this.size++
}

func (this MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.size {
		return
	}

	prev := this.root
	for i := 0; i < index; i++ {
		prev = prev.next
	}

	prev.next = prev.next.next
	this.size--

}

/*
 *  @Description:   双链表实现,(有问题)
 *  @Param:			索引是可以从0开始的
 *  @Return:
 */

type DListNode struct {
	Val int
	Next *DListNode
	Prev *DListNode
}

type MyDLinkedList struct {
	size int
	root *DListNode
	tail *DListNode
}

func Init() MyDLinkedList {
	root := &DListNode{}
	root.Next= root
	root.Prev = root
	tail := root
	return MyDLinkedList{
		0,
		root,
		tail,
	}
}

func (this *MyDLinkedList) Get(index int) int {
	if index <0 || index >= this.size {
		return -1
	}

	tmp := this.root
	for i := 0; i < index; i++ {
		tmp = tmp.Next
	}

	return tmp.Next.Val
}

func (this *MyDLinkedList) AddAtHead(val int) {
	this.AddAtIndex(0,val)
}

func (this *MyDLinkedList) AddAtTail(val int)  {
	this.AddAtIndex(this.size,val)
}

func (this *MyDLinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index > this.size  {
		return
	}
	// 尾节点直接用尾指针操作
	if index == this.size {
		node := &DListNode{val,nil,nil}
		this.tail.Next = node
		node.Prev = this.tail
		this.tail = node
		this.size++
		return
	}
	prev := this.root
	for i := 0; i <index; i++ {
		prev = prev.Next
	}

	node := &DListNode{val,nil,nil}
	node.Next = prev.Next
	node.Prev = prev
	prev.Next.Prev = node
	prev.Next = node

	this.size++
}

func (this MyDLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index > this.size {
		return
	}

	if index == this.size {
		this.tail = this.tail.Prev
		this.tail.Next = nil
	}

	prev := this.root
	for i := 0; i < index; i++ {
		prev = prev.Next
	}

	if prev.Next.Next != nil {
		prev.Next.Next.Prev = prev
	}
	prev.Next = prev.Next.Next
	this.size--
}
