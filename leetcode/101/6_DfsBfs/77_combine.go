package __DfsBfs

/**
 *  @ClassName:77_combine
 *  @Description:组合 回溯法
 *  @Author:jackey
 *  @Create:2021/7/25 下午6:14
 */

func combine(n int, k int) [][]int {
	// 创建合格的路径数组
	path := make([]int, 0)

	// 创建1-n的数字集合
	Nums := make([]int, n)
	for i := 0; i < n; i++ {
		Nums[i] = i + 1
	}

	// 创建返回结果集
	res := make([][]int, 0)
	backtracking77(Nums, path, 0, k, &res)
	return res
}

func backtracking77(Nums, path []int, level, k int, res *[][]int) {
	if len(path) == k {
		tmp := make([]int, k)
		copy(tmp, path)
		*res = append(*res, tmp)
	}

	for i := level; i < len(Nums); i++ {
		path = append(path, Nums[i])
		backtracking77(Nums, path, level+1, k, res)
		path = path[:len(path)-1]
		level++		// 避免重复选
	}
	return
}

// 优化，无需创建数字数组

func combine2(n int, k int) [][]int {
	// 判断边界问题
	if n <= 0 || k <= 0 || k > n {
		return [][]int{}
	}
	c, res := []int{}, [][]int{}
	generateCombinations(n, k, 1, c, &res)
	return res
}

// dfs+回溯
func generateCombinations(n, k, start int, path []int, res *[][]int) {
	// copy复合标准的数字组合
	if len(path) == k {
		b := make([]int, len(path))
		copy(b, path)
		*res = append(*res, b)
		return
	}
	// i will at most be n - (k - c.size()) + 1
	// start代表数字
	// c是path
	// k == n - (k-len(path)) +1
	for i := start; i <= n-(k-len(path))+1; i++ {
		path = append(path, i)
		generateCombinations(n, k, i+1, path, res)
		path = path[:len(path)-1]
	}
	return
}
