package binary_tree

/**
 *  @ClassName:minDepth
 *  @Description:111 二叉树的最小深度
 *  @Author:jackey
 *  @Create:2021/4/8 下午8:45
 */
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	depth := 1

	for len(queue) != 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			cur := queue[0]
			queue = queue[1:]
			if cur.Left == nil && cur.Right == nil {
				return depth
			}
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}

		}
		depth++
	}
	return depth
}
