package binary_tree

/**
 *  @ClassName:lowestCommonAncestor
 *  @Description:236 二叉树寻找最近的公共祖先
 *  @Author:jackey
 *  @Create:2021/3/29 上午11:29
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	/*这里是寻找最近的公共祖先，所以使用后序遍历会好一下，返回的一定是结果*/
	if root == nil {
		return nil
	}
	if root == q || root == p {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}

	/*如果q,p都不在root中，直接返回nil*/
	if left == nil && right == nil {
		return nil
	}

	/*最近的祖先可能是节点本身*/
	/*增加情况的判断*/
	if left == nil {
		return right
	}
	return left
}
