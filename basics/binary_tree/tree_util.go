package binary_tree

/**
 *  @ClassName:tree_util
 *  @Description:TODO
 *  @Author:jackey
 *  @Create:2021/3/18 上午10:19
 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

type Node struct {
	Val int
	Left *Node
	Right *Node
	Next *Node
}