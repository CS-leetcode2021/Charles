package binary_tree

/**
 *  @ClassName:insertIntoBST
 *  @Description:insertIntoBST 701
 *  @Author:jackey
 *  @Create:2021/3/23 上午11:57
 */

func insertIntoBST(root *TreeNode, key int) *TreeNode {
	if root ==nil {
		return &TreeNode{
			Val: key,
		}
	}
	if root.Val > key {
		root.Left=insertIntoBST(root.Left,key)
	}else {
		root.Right=insertIntoBST(root.Right,key)
	}

	return root
}