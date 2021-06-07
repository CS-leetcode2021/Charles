package array

/**
 *  @ClassName:twoSum
 *  @Description:TODO
 *  @Author:jackey
 *  @Create:2021/3/16 下午5:42
 */

func twoSum(nums []int,target int) []int {
	ans := map[int]int{}

	for i, v := range nums {
		targetNum := target-v
		pos , ok := ans[targetNum]
		if ok {
			return []int{i,pos}
		}else {
			ans[v] = i
		}
	}
	return []int{}
}