package _1_21

/**
 *  @ClassName:04_11_maxArea
 *  @Description:盛最多水的容器
 *  @Author:jackey
 *  @Create:2021/8/10 下午10:14
 */

// 双指针遍历，每次移动高度小的
func maxArea(height []int) int {
	i, j := 0, len(height) - 1
	m := 0
	for i < j {
		// 计算当前最大面积
		cur := (j - i) * min(height[i], height[j])
		if cur > m {
			m = cur
		}

		// 移动较小的一侧指针
		if (height[i] < height[j]) {
			i++
		} else {
			j--
		}
	}
	return m
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}