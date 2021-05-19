package binary_tree

/**
 *  @ClassName:searchBST
 *  @Description: searchBST 700
 *  @Author:jackey
 *  @Create:2021/3/23 下午12:07
 */

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val==val {
		return root
	}else if root.Val < val {
		return searchBST(root.Right,val)
	}else {
		return searchBST(root.Left,val)
	}
}
