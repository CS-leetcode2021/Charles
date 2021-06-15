package main

/**
 *  @ClassName:rob.go
 *  @Description:198 打家劫舍
 *  @Author:jackey
 *  @Create:2021/6/15 上午11:41
 */
func main() {

}

func rob(nums []int) int {
	// 当前的数据主要分两种：偷还是不偷
	if len(nums) <= 0 {
		return 0
	}

	if len(nums) == 1 { // 只有一家，一定偷
		return nums[0]
	}

	if len(nums) == 2 { // 只有两家，只偷最大的
		return robMax(nums[0], nums[1])
	}

	res := make([]int, len(nums)+1)

	res[1] = nums[0]
	res[2] = robMax(nums[0], nums[1])

	for i := 3; i <= len(nums); i++ {
		res[i] = robMax(res[i-2]+nums[i-1], res[i-1])
	}

	return res[len(nums)]

}

func robMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 打家劫舍2,首尾相连，只能选其一，可以将整个数组拆分成两个，一个只包含第一个，一个只包含最后一个

func rob2(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 2 {
		return robMax(nums[0], nums[1])
	}

	res := robMax(robHelper(nums[:len(nums)-1]), robHelper(nums[1:]))

	return res

}

func robHelper(nums []int) int {
	pre, cur := 0, 0

	for _, v := range nums {
		pre, cur = cur, robMax(cur, pre+v)
	}
	return cur
}

// 打家劫舍3：二叉树结构
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob3(root *TreeNode) int {
	memo := make(map[*TreeNode]int)

	return rober(root, memo)

}

func rober(root *TreeNode, memo map[*TreeNode]int) int {
	if root == nil {
		return 0
	}
	if _, ok := memo[root]; ok {
		return memo[root]
	}

	resInroot := root.Val
	resNotroot := 0

	if root.Left != nil {
		resInroot += rober(root.Left.Left, memo) + rober(root.Left.Right, memo)
	}

	if root.Right != nil {
		resInroot += rober(root.Right.Left,memo) + rober(root.Right.Right,memo)
	}

	resNotroot += rober(root.Left,memo) + rober(root.Right,memo)
	res := robMax(resInroot,resNotroot)
	memo[root] = res
	return res
}
