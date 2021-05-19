package binary_tree

import "strconv"

/**
 *  @ClassName:findDuplicateTree
 *  @Description:寻找重复子树 652
 *  @Author:jackey
 *  @Create:2021/3/22 下午8:32
 */

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	myMap := make(map[string]int)
	ret   := make([]*TreeNode,0)
	traverse(root,myMap,&ret)

	return ret
}

func traverse(root *TreeNode,yMap map[string]int, ret *[]*TreeNode) string{
	if root == nil {
		return " "
	}

	left := traverse(root.Left, yMap,ret)
	right := traverse(root.Right,yMap,ret)

	subTree := left+","+right+","+ strconv.Itoa(root.Val)

	count,ok := yMap[subTree]

	if !ok {
		yMap[subTree]=1
	}else if count ==1 {
		*ret = append(*ret,root)
		yMap[subTree] += 1
	}else{
		yMap[subTree] += 1
	}

	return subTree

}

