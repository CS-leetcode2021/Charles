package _1_21

/**
 *  @ClassName:06_547_findCircleNum
 *  @Description:TODO
 *  @Author:jackey
 *  @Create:2021/8/13 下午7:55
 */

func findCircleNum(isConnected [][]int) int {
	m := len(isConnected)

	vis := make([]int, m)

	count := 0

	for i := 0; i < m; i++ {	// 只需要遍历行就可以
		if vis[i] != 1 {
			dfs547(isConnected, vis, i)
			count++
		}
	}

	return count
}

func dfs547(grid [][]int, vis []int, i int) {

	vis[i] = 1
	for k := 0; k < len(grid); k++ {
		if grid[i][k]== 1 &&  vis[k] == 0 {
			dfs547(grid, vis, k)
		}
	}

	return

}