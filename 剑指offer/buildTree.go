package main

/**
 *  @ClassName:buildTree
 *  @Description:键值offer-07 重建二叉树、Leetcode 105
 *  @Author:jackey
 *  @Create:2021/7/1 下午8:33
 */
/*
 *  @Description:   使用递归将两个序列结合划分，分别将左子树和右子树找出来
 *  @Param:
 *  @Return:
 */
func buildTree(preorder []int, inorder []int) *TreeNode {

	return build(preorder, 0, len(preorder)-1, inorder, 0, len(inorder)-1)
}

func build(pre []int, pStart int, pEnd int, in []int, inStart int, inEnd int) *TreeNode {
	if pStart > pEnd {
		return nil
	}

	rootVal := pre[pStart]

	index := 0

	for i := inStart; i <= inEnd; i++ {
		if in[i] == rootVal {
			index = i
			break
		}
	}

	size := index-inStart

	root := &TreeNode{
		Val: rootVal,
	}
	root.Left = build(pre,pStart+1,pStart+size,in,inStart,index-1)
	root.Right = build(pre,pStart+size+1,pEnd,in,index+1,inEnd)

	return root
}
