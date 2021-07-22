package __greedy

import "sort"

/**
 *  @ClassName:406_reconstructQueue
 *  @Description:reconstructQueue
 *  @Author:jackey
 *  @Create:2021/7/22 下午4:57
 */

// 71/36,空间复杂度太高了啊
func reconstructQueue(people [][]int) [][]int {
	// 先把同等身高的人排在一起,不等的身高降序排列
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] > people[j][0] {
			return true
		} else if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		} else {
			return false
		}
	})

	// [[7, 0], [7, 1], [6, 1], [5, 0], [5, 2], [4, 4]]
	// 构建返回数组
	res := make([][]int, 0)

	for i := 0; i < len(people); i++ {
		if people[i][1] >= len(res) {
			res = append(res, people[i])
		} else {
			tmp := make([][]int,len(res)-people[i][1])
			copy(tmp,res[people[i][1]:])
			res = res[:people[i][1]]
			res = append(res,people[i])
			res = append(res,tmp...)
		}
	}
	return res
}
