package main

/**
 *  @ClassName:28_isSymmetric
 *  @Description:剑指offer-28 对称的二叉树 同 leetcode-101
 *  @Author:jackey
 *  @Create:2021/7/14 下午7:36
 */

// 中序遍历，返回数组，判断数组的首尾是否是对称的

var res = []int{}

func isSymmetric(root *TreeNode) bool {
	res := make([]int,0)
	interNode(root,&res)

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		if res[i] != res[j] {
			return false
		}
	}
	return true
}


// 中序需要保存空的节点，不然不能通过
func interNode(root *TreeNode,res *[]int){
	if root == nil {
		return
	}

	interNode(root.Left,res)
	*res = append(*res,root.Val)
	interNode(root.Right,res)
}


func isSymmetric2(root *TreeNode) bool {
	if root == nil {
		return true
	}


	return isEqual(root.Left,root.Right)
}

func isEqual(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if left ==  nil	 || right == nil {
		return false
	}


	return (left.Val == right.Val)&&isEqual(left.Left,right.Right)&&isEqual(left.Right,right.Left)

}