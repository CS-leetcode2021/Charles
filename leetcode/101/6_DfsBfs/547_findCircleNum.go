package __DfsBfs

/**
 *  @ClassName:547_findCircleNum
 *  @Description:省份数量
 *  @Author:jackey
 *  @Create:2021/7/24 下午8:16
 */

// 实质是图
func findCircleNum(isConnected [][]int) int {
	m := len(isConnected)

	vis := make([]int, m)

	count := 0

	for i := 0; i < m; i++ {
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