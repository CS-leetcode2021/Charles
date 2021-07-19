package main

/**
 *  @ClassName:binTree
 *  @Description:二叉树的前中后遍历模板
 *  @Author:jackey
 *  @Create:2021/7/19 下午2:31
 */

func preOrder(root *TreeNode,queue *[]*TreeNode) {
	if root == nil {
		return
	}
	*queue = append(*queue,root)
	preOrder(root.Left,queue)
	preOrder(root.Right,queue)
}

func midOrder(root *TreeNode,queue *[]*TreeNode) {
	if root == nil {
		return
	}
	midOrder(root.Left,queue)
	*queue = append(*queue,root)
	midOrder(root.Right,queue)
}
func posOrder(root *TreeNode,queue *[]*TreeNode) {
	if root == nil {
		return
	}
	posOrder(root.Left,queue)
	posOrder(root.Right,queue)
	*queue = append(*queue,root)
}
