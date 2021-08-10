package _1_21

/**
 *  @ClassName:04_986_intervalIntersection
 *  @Description:区间列表的交集
 *  @Author:jackey
 *  @Create:2021/8/10 下午9:35
 */


func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	// 正常有一个为空,就不需要求了
	if len(firstList) == 0 || len(secondList) == 0 {
		return nil
	}

	// 定义两个指针移动
	var i, j int
	// 定义一个交集
	var re [][]int
	// 当指针都移动到了各自的末尾,就结束循环
	for len(firstList) > i && len(secondList) > j {
		// 合并两个数组的交集
		res := combine(firstList[i], secondList[j])
		// 如果有交集,放到交集列表里
		if res != nil {
			re = append(re, res)
		}

		// 判断谁该往后挪, 可以看上面题解
		// 谁比较小,就谁挪
		if firstList[i][1] < secondList[j][1] {
			i++
		} else {
			j++
		}

	}
	return re
}

func combine(a, b []int) (r []int) {

	// 定义最小值的最大值,和最大值的最小值
	var minMax, maxMin int
	if a[0] > b[0] {
		minMax = a[0]
	} else {
		minMax = b[0]
	}

	if a[1] > b[1] {
		maxMin = b[1]
	} else {
		maxMin = a[1]
	}
	// 毫无相交
	if minMax > maxMin {
		return nil
	}

	return []int{minMax, maxMin}
}