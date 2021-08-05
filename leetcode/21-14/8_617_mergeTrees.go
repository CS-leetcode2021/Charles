package _1_14

/**
 *  @ClassName:8_617_mergeTrees
 *  @Description:合并二叉树
 *  @Author:jackey
 *  @Create:2021/7/30 下午6:54
 */

// 递归
// 如果值都不空，则直接相加
// 如果root2分支不为nil，root1为nil，直接添加

// 98/39
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 != nil && root2 != nil {
		root1.Val += root2.Val
	}
	if root2 == nil {
		return root1
	}
	if root1 == nil {
		return root2
	}

	root1.Left = mergeTrees(root1.Left,root2.Left)
	root1.Right = mergeTrees(root1.Right,root2.Right)

	return root1
}