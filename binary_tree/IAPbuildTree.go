package binary_tree

/**
 *  @ClassName:IAPbuildTree
 *  @Description: 106. 从中序与后序遍历序列构造二叉树
 *  @Author:jackey
 *  @Create:2021/3/22 上午11:47
 */

func buildTreePost(inorder []int, postorder []int) *TreeNode {

	return buildPost(inorder,0,len(inorder)-1,postorder,0,len(postorder)-1)
}

func buildPost( inorder []int, in_start int,in_end int,postorder []int, post_strat int , post_end int) *TreeNode {

	if post_strat>post_end {
		return nil
	}

	rootVal := postorder[post_end]

	index := 0

	for i := in_start; i <= in_end; i++ {
		if inorder[i]==rootVal {
			index = i
			break
		}
	}

	size := index-in_start

	root := &TreeNode{
		Val: rootVal,
	}


	root.Left=buildPost(inorder,in_start,index,postorder,post_strat,post_strat+size-1)

	root.Right=buildPost(inorder,index+1,in_end,postorder,post_strat+size,post_end-1)

	return root
}