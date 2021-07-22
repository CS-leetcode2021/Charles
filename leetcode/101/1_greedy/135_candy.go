package __greedy

/**
 *  @ClassName:135_candy
 *  @Description:hard
 *  @Author:jackey
 *  @Create:2021/7/22 上午9:58
 */

// 贪心，每次只考虑一侧的问题
// 99/81
func candy(ratings []int) int {
	n := len(ratings)

	tmp := make([]int, n)

	for i := 0; i < n; i++ {
		tmp[i] = 1
	}

	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			tmp[i] = tmp[i-1] + 1
		}
	}

	for i := n - 1; i > 0; i-- {
		if ratings[i] < ratings[i-1] {
			// 添加max函数。因为可能左侧的本身就比右侧的大
			tmp[i-1] = MaxCandy(tmp[i]+1,tmp[i-1])
		}
	}
	sum := 0

	for i := 0; i < n; i++ {
		sum += tmp[i]
	}

	return sum
}

func MaxCandy(i, j int) int {
	if i > j {
		return i
	}
	return j
}
