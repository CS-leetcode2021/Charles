package binary_tree

/**
 *  @ClassName:shortestBridge
 *  @Description:leetcode-934 最短的桥 DFS+BFS
 *  @Author:jackey
 *  @Create:2021/7/9 下午3:42
 */

// dfs找到其中一个岛后，问题转化为求两个岛屿之间的最短距离，bfs即可。
// 一次 DFS 标记完第一个岛屿，碰到 0 （边界）即加入队列并返回（不需要再扩散）。
// 之后对另一个岛进行BFS，向外进行扩展，当遇到2时，扩展的层数即是答案
var queue [][]int
var xShift []int = []int{-1, 1,  0, 0}
var yShift []int = []int{ 0, 0, -1, 1}
func shortestBridge(grid [][]int) int {
	// used := make([][]bool, len(grid))
	// for index := range used{
	//     used[index] = make([]bool, len(grid[0]))
	// }
	queue = make([][]int,0)
	findFirst := true
	for i:=0; i < len(grid) && findFirst; i++{
		for j:=0; j < len(grid[0]); j++{
			if 1 == grid[i][j]{
				dfsFindBridge(grid, i, j)
				findFirst=false
				break
			}
		}
	}
	ans := bfsFindBridge(grid)
	return ans
}

func dfsFindBridge(grid [][]int, i int, j int){
	if i < 0 || j < 0 || i >= len(grid)  || j >= len(grid[0]){
		return
	}
	if grid[i][j] == 0 {
		// 到达边界，坐标入队并停止搜索
		queue = append(queue,[]int{i, j})
	} else if grid[i][j] == 1 {
		grid[i][j] = 2
		dfsFindBridge(grid, i-1, j)
		dfsFindBridge(grid, i+1, j)
		dfsFindBridge(grid, i, j-1)
		dfsFindBridge(grid, i, j+1)
	}
}
func bfsFindBridge(grid [][]int)int{
	ans := 0
	for len(queue) > 0{
		ans++
		j:=len(queue)-1
		for ; j>=0; j--{
			frontI := queue[0][0]
			frontJ := queue[0][1]
			queue = queue[1:]
			for index:=0; index < 4; index++{
				newfrontI := frontI + xShift[index]
				newfrontJ := frontJ + yShift[index]
				if newfrontI >=0 && newfrontI < len(grid) && newfrontJ >=0 && newfrontJ < len(grid[0]) {
					if grid[newfrontI][newfrontJ] == 1{
						return ans
					}
					if grid[newfrontI][newfrontJ] == 2{
						continue
					}
					queue = append(queue,[]int{newfrontI, newfrontJ})
					grid[newfrontI][newfrontJ] = 2
				}
			}
		}
	}
	return ans
}