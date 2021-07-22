package binary_tree

/**
 *  @ClassName:flatten
 *  @Description:将一个二叉树转化为链表
 *  @Author:jackey
 *  @Create:2021/3/18 上午11:55
 */

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	flatten(root.Left)
	flatten(root.Right)

	/*左子树*/
	left := root.Left
	/*右子树*/
	right := root.Right

	/*将左子树挂到右子树上*/
	root.Right = left

	/*将左子树置为空*/
	root.Left = nil

	p := root

	for p.Right != nil {
		p = p.Right
	}

	/*将最初的 右子树 挂靠到 原左子树 也就是 新的右子树 上*/
	p.Right= right

}
