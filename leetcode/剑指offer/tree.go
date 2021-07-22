package main

import "fmt"

/**
 *  @ClassName:tree
 *  @Description:tree前序遍历
 *  @Author:jackey
 *  @Create:2021/7/19 下午2:53
 */


func main() {
	root :=&TreeNode{1,nil,nil}
	node2 :=&TreeNode{2,nil,nil}
	node3 :=&TreeNode{3,nil,nil}
	node4 :=&TreeNode{4,nil,nil}
	node5 :=&TreeNode{5,nil,nil}
	node6 := &TreeNode{6,nil,nil}

	root.Left = node2
	root.Right = node3
	node3.Left = node4
	node3.Right = node5
	node2.Left = node6
	res := []int{}
	MIDOrder(root,&res)

	fmt.Println(res)
}

func MIDOrder(root *TreeNode,res *[]int)  {
	if root == nil {
		return
	}

	*res = append(*res,root.Val)
	MIDOrder(root.Left,res)
	MIDOrder(root.Right,res)
	return
}