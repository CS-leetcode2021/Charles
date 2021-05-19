package binary_tree

/**
 *  @ClassName:constructMaximumBinaryTree
 *  @Description: leetcode-654 最大二叉树
 *  @Author:jackey
 *  @Create:2021/3/19 上午10:16
 */

func constructMaximumBinaryTree(nums []int) *TreeNode {
	return build01(nums, 0, len(nums))
}

func build01(nums []int, left, right int) *TreeNode {
	if nums == nil {
		return nil
	}
	if left > right {
		return nil
	}

	index := left
	Max := nums[left]

	for i := left; i < right; i++ {
		if nums[i] > Max {
			index = i
			Max = nums[i]
		}
	}

	root := new(TreeNode)
	root.Val = Max
	root.Left = build01(nums, left, index-1)
	root.Right = build01(nums, index+1, right)

	return root

}
