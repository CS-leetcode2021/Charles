package backtrack

/**
 *  @ClassName:subsets
 *  @Description:78 leetcode 集合的所有子集
 *  @Author:jackey
 *  @Create:2021/4/8 下午8:29
 */
func subsetBackTrack(pathNum,nums []int, k, start int, res *[][]int, ) {

	if len(pathNum) == k {
		temp := make([]int, k)
		copy(temp, pathNum)
		*res = append(*res, temp)
		return
	}

	for i := start; i < len(nums); i++ {
		pathNum = append(pathNum,nums[i])

		subsetBackTrack(pathNum,nums,k,i+1,res)

		pathNum = pathNum[:len(pathNum)-1]
	}

}

func subsets(nums []int) [][]int {
	n := len(nums)
	pathNum := []int{}
	start := 0
	res := [][]int{}
	for k := 0; k<= n;k++ {
		subsetBackTrack(pathNum, nums,k, start, &res)
	}

	return res
}