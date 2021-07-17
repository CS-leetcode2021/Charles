package main

/**
 *  @ClassName:levelOrder2
 *  @Description:剑指 Offer 32 - II. 从上到下打印二叉树 II
 *  @Author:jackey
 *  @Create:2021/7/8 下午1:08
 */

func levleOrder(root *TreeNode)	[][]int  {
	if root == nil {
		return nil
	}
	res := make([][]int,0)
	q := new(Queue)
	q.EnQueue(root)

	for q.Size() != 0 {
		curSize := q.Size()
		tmp := []int{}

		for i := 0; i < curSize; i++ {
			tmp = append(tmp,q.queue[i].Val)
			if q.queue[i].Left != nil {
				q.EnQueue(q.queue[i].Left)
			}

			if q.queue[i].Right != nil {
				q.EnQueue(q.queue[i].Right)
			}
		}
		res = append(res,tmp)
		q.queue = q.queue[curSize:]
		tmp = tmp[0:0]
	}
	return res

}