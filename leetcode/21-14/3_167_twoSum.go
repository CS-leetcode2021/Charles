package _1_14

/**
 *  @ClassName:3_167_twoSum
 *  @Description:两数之和进阶版
 *  @Author:jackey
 *  @Create:2021/7/26 上午11:32
 */

func twoSum(numbers []int, target int) []int {
	n := len(numbers)

	for i, j := 0, n-1; i < j; {
		if numbers[i] + numbers[j] == target {
			return []int{i+1,j+1}
		}else if numbers[i]+numbers[j] < target {
			i++
		}else {
			j--
		}
	}
	return nil
}