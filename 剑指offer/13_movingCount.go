package main

import "container/list"
// 剑指offer-13 机器人移动问题
// 1、BFS、DFS都可以求解，本次采用DFS
// 2、需要辅助函数来将坐标差分为个位数组，然后返回相加的数字之和
// 3、需要方位数组：上下左右
// 4、需要判定函数
// 5、需要设置标记数组

var Mx = [4]int{1, -1, 0, 0}
var My = [4]int{0, 0, 1, -1}

func movingCount(m, n int, k int) int {
	tag := make([][]int,m)
	for i := 0; i < m; i++ {
		tag[i] = make([]int,n)
	}
	 dfsMov(tag,m,n,0,0,k)
	num := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if tag[i][j] == 1 {
				num++
			}
		}
	}
	return  num
}

func dfsMov(tag [][]int,m, n, x, y ,k int){

	tag[x][y] = 1
	for i := 0; i < 4; i++ {
		tmp_x := x + Mx[i]
		tmp_y := y + My[i]
		if isInMov(m, n, tmp_x, tmp_y) && numSum(tmp_x,tmp_y) <=k  && tag[tmp_x][tmp_y] != 1 {
			dfsMov(tag,m,n,tmp_x,tmp_y,k)
		}
	}
	return
}

func numSum(x, y int) int {
	sum_x := 0
	sum_y := 0
	for x/10 != 0 {
		sum_x += x%10
		x = x/10
	}
	sum_x += x

	for y/10 != 0 {
		sum_y += y%10
		y = y/10
	}
	sum_y += y
	return sum_y+sum_x
}

func isInMov(m, n, x, y int) bool {
	if x < 0 || y < 0 || x >= m || y >= n {
		return false
	}
	return true
}

// ------------------------BFS
func movingCount1(m int, n int, k int) int {
	queue := list.List{}
	visited := make([][]bool, m)
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, n)
	}
	queue.PushBack([]int{0, 0, 0, 0})
	res := 0
	for queue.Len() > 0 {
		back := queue.Back()
		queue.Remove(back)
		bfs := back.Value.([]int)
		i := bfs[0]
		j := bfs[1]
		si := bfs[2]
		sj := bfs[3]
		if i >= m || j >= n || si+sj > k || visited[i][j] {
			continue
		}
		res++
		visited[i][j] = true

		sj1 := sj + 1
		si1 := si + 1
		if (j+1)%10 == 0 {
			sj1 = sj - 8
		}
		if (i+1)%10 == 0 {
			si1 = si - 8
		}

		queue.PushBack([]int{i + 1, j, si1, sj})
		queue.PushBack([]int{i, j + 1, si, sj1})
	}
	return res
}