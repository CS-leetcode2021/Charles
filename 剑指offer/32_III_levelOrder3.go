package main

/**
 *  @ClassName:levelOrder
 *  @Description:剑指 Offer 32 - III. 从上到下打印二叉树 III 之字形顺序打印二叉树
 *  @Author:jackey
 *  @Create:2021/7/8 下午1:22
 */

func levelOrder3(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	res := [][]int{}
	q := new(Queue)
	q.EnQueue(root)
	level := 0
	for q.Size() != 0 {
		cursize := q.Size()
		tmp := []int{}
		level++
		if level%2 == 1 {
			for i := 0; i < cursize; i++ {
				tag := q.DeQueue()
				tmp = append(tmp, tag.Val)
				if tag.Left != nil {
					q.EnQueue(tag.Left)
				}
				if tag.Right != nil {
					q.EnQueue(tag.Right)
				}
			}
		} else {
			for i := 0; i < cursize; i++ {
				tag := q.DeQueue()
				tmp = append(tmp, tag.Val)
				if tag.Left != nil {
					q.EnQueue(tag.Left)
				}
				if tag.Right != nil {
					q.EnQueue(tag.Right)
				}
			}
			for i, j := 0, len(tmp)-1; i < j; i, j = i+1, j-1 {
				tmp[i],tmp[j] = tmp[j],tmp[i]
			}
		}

		res = append(res,tmp)
	}
	return res
}
