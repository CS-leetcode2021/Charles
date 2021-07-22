package Figure

/**
 *  @ClassName:common
 *  @Description:common
 *  @Author:jackey
 *  @Create:2021/7/7 下午12:48
 */

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