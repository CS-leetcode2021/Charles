package binary_tree

/**
 *  @ClassName:buildTree
 *  @Description: 105. 从前序与中序遍历序列构造二叉树
 *  @Author:jackey
 *  @Create:2021/3/22 上午11:24
 */

func buildTree(preorder []int, inorder []int) *TreeNode {

	return build(preorder,0,len(preorder)-1,inorder,0,len(inorder)-1)
}

func build(preorder []int, pre_start int, pre_end int, inorder []int, in_start int,in_end int) *TreeNode {

	if pre_start>pre_end {
		return nil
	}

	rootVal := preorder[pre_start]

	index := 0

	for i := in_start; i <= in_end; i++ {
		if inorder[i]==rootVal{
			index = i
			break
		}
	}

	size := index-in_start
	root := &TreeNode{
		Val: rootVal,
	}

	root.Left=build(preorder,pre_start+1,pre_start+size,inorder,in_start,in_end-1)
	root.Right=build(preorder,pre_start+size+1,pre_end,inorder,index+1,in_end)



	return root
}