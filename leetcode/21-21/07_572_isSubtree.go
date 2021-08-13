package _1_21

/**
 *  @ClassName:07_572_isSubtree
 *  @Description:子树
 *  @Author:jackey
 *  @Create:2021/8/13 下午8:25
 */

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}
	return check(root, subRoot) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func check(s, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}

	if s == nil || t == nil {
		return false
	}

	if s.Val == t.Val {
		return check(s.Left, t.Left) && check(s.Right, t.Right)
	}
	return false
}
