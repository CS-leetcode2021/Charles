package _1_21

/**
 *  @ClassName:util
 *  @Description:TODO
 *  @Author:jackey
 *  @Create:2021/8/7 下午4:34
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// 二叉树
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
