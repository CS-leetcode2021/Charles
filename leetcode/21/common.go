package _1

/**
 *  @ClassName:common
 *  @Description:TODO
 *  @Author:jackey
 *  @Create:2021/7/26 下午3:01
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

type Coordinate struct {
	X   int
	Y   int
	Val int
}

type QueueCoordinate struct {
	queue []*Coordinate
}

func (qc QueueCoordinate) Size() int {
	return len(qc.queue)
}

func (qc *QueueCoordinate) EnQueue(coord *Coordinate) {
	if qc.Size() == 0 {
		qc.queue = make([]*Coordinate, 0)
	}

	qc.queue = append(qc.queue, coord)
	return
}

func (qc *QueueCoordinate) DeQueue() *Coordinate {
	if qc.Size() == 0 {
		return nil
	}

	res := qc.queue[0]
	qc.queue = qc.queue[1:]

	return res
}
