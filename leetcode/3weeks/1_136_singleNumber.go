package _weeks

/**
 *  @ClassName:1_136_singleNumber
 *  @Description:只出现一次的数字
 *  @Author:jackey
 *  @Create:2021/8/7 下午7:35
 */

func singleNumber(nums []int) int {
	tMap := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		tMap[nums[i]]++
	}

	for i, v := range tMap {
		if v == 1 {
			return i
		}
	}
	return -1
}

func singleNumber2(nums []int) int {
	single := 0
	for _, num := range nums {
		single ^= num
	}
	return single
}
