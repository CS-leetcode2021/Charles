package main

import "go/doc"

/**
 *  @ClassName:34_pathSum
 *  @Description:剑指offer-34 二叉树中和为某一值的路径 同 leetcode-113
 *  @Author:jackey
 *  @Create:2021/7/17 下午9:50
 */

// dfs + 回溯

func pathSum2(root *TreeNode, target int) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int,0)
	path := make([]int,0)
	path = append(path,root.Val)
	dfsPath(root,path,&res, target-root.Val)
	return res
}

func dfsPath(root *TreeNode, path []int, res *[][]int, target int) {
	if root.Left == nil && root.Right == nil {
		if target == 0 {
			tmp := make([]int,len(path))	// 为了防止切片的底层修改，使用copy函数进行新的复制
			copy(tmp,path)
			*res = append(*res,tmp)
		}
		return
	}

	if root.Left != nil {
		path = append(path,root.Left.Val)
		dfsPath(root.Left,path,res,target-root.Left.Val)
		// 回溯
		path = path[:len(path)-1]
	}
	if root.Right != nil {
		path = append(path,root.Right.Val)
		dfsPath(root.Right,path,res,target-root.Right.Val)
		// 回溯
		path = path[:len(path)-1]
	}

}

func pathSum3(root *TreeNode, sum int) [][]int {
	if root == nil {
		return nil
	}
	var ret [][]int
	dfs(root,sum,[]int{},&ret)
	return ret
}

func dfs(root *TreeNode,sum int,arr []int,ret *[][]int){
	if root == nil{
		return
	}
	arr = append(arr,root.Val)

	if root.Val == sum && root.Left == nil && root.Right == nil {
		//slice是一个指向底层的数组的指针结构体
		//因为是先序遍历，如果 root.Right != nil ,arr 切片底层的数组会被修改
		//所以这里需要 copy arr 到 tmp，再添加进 ret，防止 arr 底层数据修改带来的错误
		tmp := make([]int,len(arr))
		copy(tmp,arr)
		*ret = append(*ret,tmp)
	}

	dfs(root.Left,sum - root.Val,arr,ret)
	dfs(root.Right,sum - root.Val,arr,ret)

	arr = arr[:len(arr)-1]	// 回溯操作
}