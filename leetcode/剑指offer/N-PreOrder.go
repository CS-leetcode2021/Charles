package main

/**
 *  @ClassName:N-PreOrder
 *  @Description:leetcode-589 N  叉树的前序遍历
 *  @Author:jackey
 *  @Create:2021/7/8 下午2:15
 */

//  递归 空间复杂度大
func preorder(root *Node) []int {
	if root == nil {
		return nil
	}

	res := []int{}

	res = append(res, root.Val)
	for _, node := range root.Children {
		res = append(res, preorder(node)...)
	}

	return res
}

// stack

func preorder2(root *Node) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)

	stack := []*Node{root}

	for len(stack) != 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, cur.Val)
		for i := len(cur.Children) - 1; i >= 0; i-- {
			stack = append(stack, cur.Children[i])
		}
	}
	return res
}

func preorder3(root *Node) []int {
	if root == nil {
		return nil
	}

	res := make([]int, 0)
	s := new(Stack)
	s.EnStack(root)
	for s.Size() != 0 {
		n := s.DeStack()
		res = append(res,n.Val)
		for i := len(n.Children) - 1; i >= 0; i-- {
			s.EnStack(n.Children[i])
		}
	}
	return res
}
