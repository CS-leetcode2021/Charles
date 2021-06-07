package binary_tree

/**
 *  @ClassName:isValidBST
 *  @Description:isValidBST 验证是否是二叉搜索树 98
 *  @Author:jackey
 *  @Create:2021/3/23 下午12:11
 */

/*检测不到左右子树的孩子节点*/
func isValidBST01(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if root.Left != nil && getmax(root.Left).Val >= root.Val {
		return false
	}
	if root.Right != nil && getmin(root.Right).Val <= root.Val {
		return false
	}

	return isValidBST(root.Left) && isValidBST(root.Right)
}

/*改进,增加辅助函数*/
func isValidBST(root *TreeNode) bool {
	return isValidBst(root,nil,nil)
}

func isValidBst(root *TreeNode, min *TreeNode, max *TreeNode) bool {
	if root == nil {
		return true
	}
	if min != nil && root.Val <= min.Val {
		return false
	}
	if max != nil && root.Val >= max.Val {
		return false
	}
	return isValidBst(root.Left,min,root)&&isValidBst(root.Right,root,max)
}
