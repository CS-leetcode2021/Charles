package __greedy

/**
 *  @ClassName:605_canPlaceFlowers
 *  @Description:canPlaceFlowers 贪心算法
 *  @Author:jackey
 *  @Create:2021/7/22 下午2:34
 */

// 暴力解，问题太多，通过率低下
func canPlaceFlowers(flowerbed []int, n int) bool {
	if len(flowerbed) == 0 {
		return false
	}
	addCount := 0
	if len(flowerbed) == 1 && flowerbed[0] == 0 {
		if n <= 1 {
			return true
		}
		return false
	}
	for i := 0; i < len(flowerbed)-1; i++ {
		if i == 0 && flowerbed[0] == 0 && flowerbed[1] == 0 {
			flowerbed[i] = 1
			addCount++
		} else if i >= 1 && i < len(flowerbed)-1 && flowerbed[i] == 0 && flowerbed[i-1] == 0 && flowerbed[i+1] == 0 {
			flowerbed[i] = 1
			addCount++
		} else if i == len(flowerbed)-2 && flowerbed[i] == 0 && flowerbed[i+1] == 0 {
			flowerbed[i+1] = 1
			addCount++
		}
	}

	if addCount >= n {
		return true
	}

	return false
}

// 空间复杂度过高
func canPlaceFlowers2(flowerbed []int, n int) bool {
	tmp := []int{0}

	tmp = append(tmp, flowerbed...)
	tmp = append(tmp, 0)

	for i := 1; i < len(tmp); i++ {

		if tmp[i] == 0 && tmp[i-1] == 0 && tmp[i+1] == 0 {
			tmp[i] = 1
			n--
		}
		if n <= 0 { // 防在后面判断，不然会越界，导致下一轮进不来
			return true
		}
	}
	return false
}

// 优化 跳格子
// 找到当前值：
// 1、如果是0,则查看后面是否是0.如果是则插入，往回走两格
// 2、如果是0，则后面是1,则插入，往后走三格
// 3、如果是1,则向后跳两个

// 90/99
func canPlaceFlowers3(flowerbed []int, n int) bool {

	for i := 0; i < len(flowerbed); {
		if flowerbed[i] == 1 {
			i += 2
		} else if i == len(flowerbed)-1 || flowerbed[i+1] == 0 {
			n--
			i += 2
		} else {
			i += 3
		}
		if n <= 0 {
			return true
		}
	}

	return false
}
