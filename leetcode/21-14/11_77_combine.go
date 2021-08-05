package _1_14

/**
 *  @ClassName:11_77_combine
 *  @Description:组合+回溯+dfs
 *  @Author:jackey
 *  @Create:2021/8/3 下午4:13
 */

func combine(n int, k int) [][]int {
	// 创建path
	path := make([]int, 0)

	// 创建nums
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i + 1
	}

	res := make([][]int, 0)
	backtracking77(nums, path, 0, k, &res)
	return res
}

func backtracking77(nums, path []int, level, k int, res *[][]int) {
	if len(path) == k {
		tmp := make([]int, k)
		copy(tmp, path)
		*res = append(*res, tmp)
		return
	}

	for i := level; i < len(nums); i++ {
		path = append(path, nums[i])
		backtracking77(nums, path, level+1, k, res)
		path = path[:len(path)-1]
		level++
	}
	return
}
