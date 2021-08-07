package _weeks

/**
 *  @ClassName:1_15_threeSum
 *  @Description:三数之和
 *  @Author:jackey
 *  @Create:2021/8/7 下午7:44
 */

// 优化代码
// 使用辅助函数判定两数之和是否等于target
// 存在重复的问题
func threeSum(nums []int) [][]int {
	n := len(nums)
	if n < 3 {
		return nil
	}

	res := make([][]int, 0)
	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		target := (-1) * nums[i]

		tmp := twoSum(nums, target)

		tmp = append(tmp, nums[i])
		res = append(res, tmp)

	}
	return res
}

func twoSum(nums []int, target int) []int {
	res := make(map[int]int)

	for i, v := range nums {
		targetNum := target - v

		if _, ok := res[targetNum]; ok {
			return []int{v, targetNum}
		} else {
			res[v] = i
		}
	}
	return nil
}
