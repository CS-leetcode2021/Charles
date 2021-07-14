package main

// 剑指offer-26 树的子结构
// 1、首先判断root的值是否一样，如果不一样就去左子树或者右子树进行查找
// 2、如果root的值一样，就去判断子结构是否一样
// 3、判断子结构也是递归遍历
// 4、如果val是double类型，需要注意相等的判断方式
// 5、https://leetcode-cn.com/problems/shu-de-zi-jie-gou-lcof/solution/yi-pian-wen-zhang-dai-ni-chi-tou-dui-che-uhgs/
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	res := false
	if A != nil && B != nil {
		if Equal(A.Val,B.Val){
			res = isSubTree(A,B)
		}
		if !res {
			res = isSubStructure(A.Left,B)
		}

		if !res {
			res = isSubStructure(A.Right,B)
		}
	}

	return res
}

func Equal(a, b int) bool {
	if a == b {
		return true
	}
	return false
}

func isSubTree(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return true
	}
	if A == nil {
		return false
	}

	if !Equal(A.Val, B.Val) {
		return false
	}

	return isSubTree(A.Left,B.Left)&&isSubTree(A.Right,B.Right)
}
