package binary_tree

/**
 *  @ClassName:connect_node
 *  @Description:链接同层的右侧节点
 *  @Author:jackey
 *  @Create:2021/3/18 上午11:34
 */

func connect(root *Node) *Node {
	if root == nil || root.Left == nil {
		return root
	}

	root.Left.Next = root.Right
	connect(root.Left)
	connect(root.Right)

	return root

}

/*需要解决的问题是如何跨父节点进行连接*/

func Node2(root *Node) *Node {
	if root == nil {
		return root
	}
	connectTwoNode(root.Left,root.Right)
	return root
}

func connectTwoNode(n1 *Node, n2 *Node) {
	if n1 == nil || n2 == nil {
		return
	}
	n1.Next=n2

	connectTwoNode(n1.Left,n1.Right)
	connectTwoNode(n2.Left,n2.Right)

	connectTwoNode(n1.Right,n2.Left)
}
