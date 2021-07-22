package main

/**
 *  @ClassName:27_mirrorTree
 *  @Description:剑指offer-27 二叉树的镜像 同 leetcode-226 反转二叉树
 *  @Author:jackey
 *  @Create:2021/7/14 下午7:28
 */

// 100/100
func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	mirrorTree(root.Left)
	mirrorTree(root.Right)

	tmp := root.Left

	root.Left = root.Right
	root.Right = tmp
	return root
}
