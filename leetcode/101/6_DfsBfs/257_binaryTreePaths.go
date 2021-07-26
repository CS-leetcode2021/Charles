package __DfsBfs

import "strconv"

/**
 *  @ClassName:257_binaryTreePaths
 *  @Description:二叉树的所有路径
 *  @Author:jackey
 *  @Create:2021/7/26 下午8:58
 */

// 前序遍历 dfs
// string 可以直接拼接
// 空间复杂度过高
func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	res := []string{}

	if root.Left == nil && root.Right == nil {
		return []string{strconv.Itoa(root.Val)}
	}

	tmpLeft := binaryTreePaths(root.Left)
	for i := 0; i < len(tmpLeft); i++ {
		res = append(res, strconv.Itoa(root.Val)+"->"+tmpLeft[i])
	}
	tmpRight := binaryTreePaths(root.Right)
	for i := 0; i < len(tmpRight); i++ {
		res = append(res, strconv.Itoa(root.Val)+"->"+tmpRight[i])
	}
	return res

}
