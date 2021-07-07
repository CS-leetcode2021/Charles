package binary_tree

/**
 *  @ClassName:levelOrderBottom
 *  @Description:leetcode-107 二叉树的层次遍历 自底向上
 *  @Author:jackey
 *  @Create:2021/7/7 下午8:58
 */

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	res := [][]int{}
	queue := []*TreeNode{root}

	for len(queue) != 0 {
		currSize := len(queue)
		tmp := []int{}

		for i := 0; i < currSize; i++ {
			tmp = append(tmp, queue[i].Val)
			if queue[i].Left != nil {
				queue = append(queue,queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue,queue[i].Right)
			}
		}
		res = append(res,tmp)
		queue = queue[currSize:]
		tmp = tmp[0:0]
	}

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j+1 {
		res[i],res[j] = res[j],res[i]
	}

	return res
}
