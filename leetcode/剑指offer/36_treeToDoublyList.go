package main

/**
 *  @ClassName:36_treeToDoublyList
 *  @Description:剑指offer-36 二叉搜索树转化为排序的双向链表 同leetcode-426
 *  @Author:jackey
 *  @Create:2021/7/19 下午1:56
 */

// 1、递归寻找比根节点更大或更小的节点
func treeToDoublyList(root *TreeNode) *TreeNode{
	if root == nil {
		return nil
	}

	queue := []*TreeNode{}
	traverse(&queue,root)

	head := new(TreeNode)

	p := head

	for _, v := range queue {
		if v == nil {
			continue
		}
		p.Right = v
		if p.Right != nil {
			p.Right.Left = p
			p = p.Right
		}
	}

	p.Right = head.Right
	head.Right.Left = p
	return head.Right
}


// 中序遍历
func traverse(queue *[]*TreeNode,root *TreeNode)  {
	if root == nil {
		return
	}

	traverse(queue,root.Left)

	*queue = append(*queue,root)

	traverse(queue,root.Right)
}