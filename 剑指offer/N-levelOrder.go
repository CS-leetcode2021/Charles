package main

/**
 *  @ClassName:N-levelOrder
 *  @Description:leetcode-429. N 叉树的层序遍历
 *  @Author:jackey
 *  @Create:2021/7/8 下午3:56
 */

func levelOrder4(root *Node) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)
	qn := new(QueueNode)
	qn.EnQueue(root)

	for qn.Size() != 0 {
		cursize := qn.Size()
		tmp := make([]int,0)
		for i := 0; i < cursize; i++ {
			n := qn.DeQueue()
			tmp = append(tmp,n.Val)
			if len(n.Children) != 0 {
				for i := 0; i < len(n.Children); i++ {
					qn.EnQueue(n.Children[i])
				}
			}
		}
		res = append(res,tmp)
	}
	return res
}

