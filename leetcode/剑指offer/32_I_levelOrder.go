package main

/**
 *  @ClassName:levelOrder
 *  @Description:剑指 Offer 32 - I. 从上到下打印二叉树
 *  @Author:jackey
 *  @Create:2021/7/8 下午12:47
 */

func levelOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := []int{}
	q := new(Queue)
	q.EnQueue(root)
	for q.Size() != 0 {
		tmp := q.DeQueue()
		res = append(res,tmp.Val)
		if tmp.Left != nil {
			q.EnQueue(tmp.Left)
		}
		if tmp.Right != nil {
			q.EnQueue(tmp.Right)
		}

	}
	return res
}
