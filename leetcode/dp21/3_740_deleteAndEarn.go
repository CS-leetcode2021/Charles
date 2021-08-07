package dp21

/**
 *  @ClassName:3_740_deleteAndEarn
 *  @Description:删除并获得点数
 *  @Author:jackey
 *  @Create:2021/8/7 下午5:53
 */
// 打家劫舍问题的变形
// 我们在原来的 nums 的基础上构造一个临时的数组 all，这个数组，以元素的值来做下标，下标对应的元素是原来的元素的个数。
//
//举个例子：
//
//nums = [2, 2, 3, 3, 3, 4]
//
//构造后：
//
//all=[0, 0, 2, 3, 1];
//
//就是代表着 22 的个数有两个，33 的个数有 33 个，44 的个数有 11 个。
//
//其实这样就可以变成打家劫舍的问题了呗。
//
//我们来看看，打家劫舍的最优子结构的公式：
//
//dp[i] = Math.max(dp[i - 1], dp[i - 2] + nums[i]);
//
//再来看看现在对这个问题的最优子结构公式：
//
//dp[i] = Math.max(dp[i - 1], dp[i - 2] + i * all[i]);
//

func deleteAndEarn(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	tmap := make(map[int]int)
	tmpMax := 0
	for i := 0; i < len(nums); i++ {
		if tmpMax < nums[i] {
			tmpMax = nums[i]
		}
		tmap[nums[i]]++
	}

	res := make([]int, tmpMax+1)

	for i, V := range tmap {
		res[i] = i*V
	}

	return rob(res)
}
