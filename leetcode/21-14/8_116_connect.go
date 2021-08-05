package _1_14

/**
 *  @ClassName:8_116_connect
 *  @Description:填充每一个节点的下一个右侧节点指针
 *  @Author:jackey
 *  @Create:2021/7/30 下午7:11
 */

// 递归
// 思考每一个root应该怎么做
// 95/99
func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	connectTwoNode(root.Left, root.Right)
	return root
}
func connectTwoNode(n1 *Node, n2 *Node) {
	if n1 == nil || n2 == nil {
		return
	}

	n1.Next = n2
	connectTwoNode(n1.Left, n1.Right)
	connectTwoNode(n2.Left, n2.Right)
	connectTwoNode(n1.Right, n2.Left)
}
