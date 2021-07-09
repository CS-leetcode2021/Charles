package binary_tree

/**
 *  @ClassName:tree_util
 *  @Description:TODO
 *  @Author:jackey
 *  @Create:2021/3/18 上午10:19
 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

type Node struct {
	Val int
	Left *Node
	Right *Node
	Next *Node
}
type Queue struct {
	queue []int
}

func (q Queue) Size() int {
	return len(q.queue)
}

func (q *Queue) EnQueue(v int) {
	if q.Size() == 0 {
		q.queue = make([]int,0)
	}

	q.queue = append(q.queue,v)
}

func (q *Queue) Dequeue() int {
	if q.Size() == 0 {
		return -1
	}

	tmp := q.queue[0]
	q.queue = q.queue[1:]
	return tmp
}


type Pos struct {
	posX int
	posY int
	val  int
}

type QueueP struct {
	queue []*Pos
}

func (q *QueueP) Size() int {
	return len(q.queue)
}

func (q *QueueP) EnQueue(v *Pos) {
	if q.Size() == 0 {
		q.queue = make([]*Pos,0)
	}

	q.queue = append(q.queue,v)
}

func (q *QueueP) Dequeue() *Pos {
	if q.Size() == 0 {
		return nil
	}

	tmp := q.queue[0]
	q.queue = q.queue[1:]
	return tmp
}