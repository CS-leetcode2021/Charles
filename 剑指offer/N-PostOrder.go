package main

/**
 *  @ClassName:N-PostOrder
 *  @Description:leetcode-590 N 叉树的后序遍历
 *  @Author:jackey
 *  @Create:2021/7/8 下午3:06
 */

// stack
func postorder(root *Node) []int {
	if root == nil {
		return nil
	}
	res := make([]int,0)
	s := new(Stack)
	s.EnStack(root)

	for s.Size() != 0 {
		n := s.DeStack()
		res = append(res,n.Val)
		for i := 0;i<len(n.Children);i++ {
			s.EnStack(n.Children[i])
		}
	}

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i],res[j] = res[j],res[i]
	}

	return res
}

// DFS
func postorder2(root *Node) []int {
	if root == nil {
		return nil
	}
	res := make([]int,0)

	for _, node := range root.Children {
		res = append(res,postorder2(node)...)
	}
	res = append(res,root.Val)
	return res
}