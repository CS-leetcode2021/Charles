package binary_tree

import "math"

/**
 *  @ClassName:countNodes
 *  @Description:222 统计完全二叉树的节点
 *  @Author:jackey
 *  @Create:2021/3/29 下午1:44
 */

func countNodes(root *TreeNode) int {
	/*
	 *  @Description:   完全二叉树、普通二叉树、满二叉树各有不同，
	 *  @Param:         输入的跟节点
	 *  @Return:        返回节点数
	 */

	/*完全二叉树并不一定是满二叉树，因此需要结合普通二叉树的方式实现节点的统计*/

	l := root
	r := root

	hl := 0
	hr := 0

	for l != nil {
		l = l.Left
		hl++
	}

	for r != nil {
		r = r.Right
		hr++
	}

	if hl == hr {
		return int(math.Pow(2, float64(hl)) - 1)
	}

	return 1+countNodes(root.Left)+countNodes(root.Right)

}

/*普二叉树的计算方式*/
func OrdinaryBinTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + OrdinaryBinTree(root.Left) + OrdinaryBinTree(root.Right)
}

/*满二叉树的计算方式*/
func FullBinTree(root *TreeNode) int {
	h := 0
	for root != nil {
		root = root.Left
		h++
	}
	return int(math.Pow(2, float64(h)) - 1)
}
