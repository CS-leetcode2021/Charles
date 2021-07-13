package main

/**
 *  @ClassName:common
 *  @Description:common
 *  @Author:jackey
 *  @Create:2021/7/1 下午8:11
 */

type ListNode struct {
	Val int
	Next *ListNode
}

func (ln *ListNode) EnList(node *ListNode) {
	tmp := ln
	for tmp.Next != nil {
		tmp = tmp.Next
	}

	tmp.Next = node
	node.Next = nil
	return
}

func (ln *ListNode) DeList(val int) *ListNode {
	p := ln
	q := p.Next
	for q.Next != nil {
		if q.Val == val {
			break
		}
		p = q
		q = q.Next
	}

	p.Next = q.Next

	return q
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

type Queue struct {
	queue []*TreeNode
}

func (q *Queue) Size() int {
	return len(q.queue)
}

func (q *Queue) EnQueue(v *TreeNode)  {
	if q.Size() == 0 {
		q.queue = make([]*TreeNode,0)
	}
	q.queue = append(q.queue,v)
	return
}

func (q *Queue) DeQueue() *TreeNode {
	if q.Size() <= 0 {
		return nil
	}

	v := q.queue[0]
	q.queue = q.queue[1:]
	return v
}

type Node struct {
	Val int
	Children []*Node
}

type Stack struct {
	stack []*Node
}

func (s *Stack) Size() int {
	return len(s.stack)
}

func (s *Stack) EnStack(n *Node) {
	if s.Size() == 0 {
		s.stack = make([]*Node,0)
	}

	s.stack = append(s.stack,n)
	return
}

func (s *Stack) DeStack() *Node {
	if s.Size() <= 0 {
		return nil
	}

	node := s.stack[s.Size()-1]
	s.stack = s.stack[:s.Size()-1]

	return node
}

func (s *Stack) Last() *Node {
	return s.stack[s.Size()-1]
}


type QueueNode struct {
	queue []*Node
}

func (qn *QueueNode) Size() int {
	return len(qn.queue)
}

func (qn *QueueNode) EnQueue(n *Node) {
	if qn.Size() == 0 {
		qn.queue = make([]*Node,0)
	}
	qn.queue = append(qn.queue,n)
	return
}

func (qn *QueueNode) DeQueue() *Node {
	if qn.Size() == 0 {
		return nil
	}
	n := qn.queue[0]
	qn.queue= qn.queue[1:]
	return n
}