package main

/**
 *  @ClassName:isCousins
 *  @Description:leetcode-993 二叉树的堂兄弟节点
 *  @Author:jackey
 *  @Create:2021/7/8 下午4:30
 */

func isCousins(root *TreeNode, x int, y int) bool {
	if root == nil {
		return false
	}
	first := findFa(root,x)
	second := findFa(root,y)
	if first == nil ||second == nil || first == second{
		return false
	}
	res := bfs(root)
	l1 := isLevel(first,res)
	l2 := isLevel(second,res)
	if l1 == l2 {
		return true
	}
	return false
}

func findFa(root *TreeNode, v int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Left != nil && root.Left.Val == v{
		return root
	}
	if root.Right != nil && root.Right.Val == v{
		return root
	}

	LeftNode := findFa(root.Left,v)
	RightNode := findFa(root.Right,v)

	if LeftNode != nil {
		return LeftNode
	}
	if RightNode != nil {
		return RightNode
	}
	return nil
}

func isLevel(node *TreeNode,res [][]int) int {
	v := node.Val
	m := len(res)
	tag := 0
	for i := 0; i < m; i++ {
		for j := 0; j < len(res[i]); j++ {
			if res[i][j] == v {
				tag = i
			}
		}
	}
	return tag
}

func bfs(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int,0)
	q := new(Queue)
	q.EnQueue(root)

	for q.Size() != 0 {
		curSize := q.Size()
		tmp := []int{}

		for i := 0; i < curSize; i++ {
			tmp = append(tmp,q.queue[i].Val)
			if q.queue[i].Left != nil {
				q.EnQueue(q.queue[i].Left)
			}

			if q.queue[i].Right != nil {
				q.EnQueue(q.queue[i].Right)
			}
		}
		res = append(res,tmp)
		q.queue = q.queue[curSize:]
		tmp = tmp[0:0]
	}
	return res
}

//
func isCousins2(root *TreeNode, x int, y int) bool {
	var yDepth, xDepth int
	var xParent, yParent *TreeNode

	var dfs func(*TreeNode, int, *TreeNode)
	dfs = func(node *TreeNode, depth int, parent *TreeNode) {
		if node == nil {
			return
		}
		if x == node.Val {
			xDepth = depth
			xParent = parent
		}
		if y == node.Val {
			yDepth = depth
			yParent = parent
		}
		dfs(node.Left, depth + 1, node)
		dfs(node.Right, depth + 1, node)
	}
	dfs(root, 0, nil)
	return xDepth == yDepth && xParent != yParent
}