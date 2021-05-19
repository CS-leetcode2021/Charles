package binary_tree

/**
 *  @ClassName: invertTree
 *  @Description: 翻转二叉树
 *  @Author: jack
 *  @Create: 2021/3/18 上午10:18
 */

func invertTree1(root *TreeNode) *TreeNode {

	/*这里使用前序遍历或者 后序遍历进行翻转*/
	if root==nil{
		return nil
	}

	tmp := root.Left;
	root.Left=root.Right
	root.Right=tmp

	invertTree1(root.Left)
	invertTree1(root.Right)

	return root

	return nil
}

/*如果使用中序遍历的话，这里要先进行左子树的翻转
 *然后进行左右子树的交换
 *再进行右子树的翻转，其实是对原始左子树进行了两次翻转
 *实际中的右子树没有进行遍历和翻转
*/
func invertTree2(root *TreeNode) *TreeNode {
	if root==nil{
		return nil
	}

	invertTree2(root.Left)
	tmp := root.Left;
	root.Left=root.Right
	root.Right=tmp
	invertTree2(root.Right)

	return root
}

/*改进，进行两次左子树的遍历翻转解决以上问题*/
/*同样的如果是中序遍历翻转，可以进行两次右子树进行翻转*/
/*以后多使用java调试*/
func invertTree3(root *TreeNode) *TreeNode {
	if root==nil{
		return nil
	}

	invertTree3(root.Left)
	tmp := root.Left;
	root.Left=root.Right
	root.Right=tmp
	invertTree3(root.Right)

	return root
}