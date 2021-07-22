package binary_tree

/**
 *  @ClassName:deleteNode
 *  @Description:删除二叉树节点 450
 *  @Author:jackey
 *  @Create:2021/3/23 上午11:40
 */

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root==nil {
		return nil
	}
	if root.Val==key{
		if root.Left ==nil {
			return root.Right
		}
		if root.Right==nil {
			return root.Left
		}
		root.Val = getmin(root.Right).Val

		root.Right=deleteNode(root.Right,getmin(root.Right).Val)
	}else if root.Val<key{
		root.Right=deleteNode(root.Right,key)
	}else {
		root.Left= deleteNode(root.Left,key)
	}
	return root
}

func getmin(node *TreeNode) *TreeNode {
	for node.Left!=nil {
		node = node.Left
	}
	return node
}

func getmax(node *TreeNode) *TreeNode {
	for node.Right!=nil {
		node = node.Right
	}
	return node
}