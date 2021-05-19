package backtrack

/**
 *  @ClassName:combine
 *  @Description:77\78 leetcode 组合
 *  @Author:jackey
 *  @Create:2021/4/8 下午6:45
 */

func combine(n int, k int) [][]int {
	/*回溯：选择列表、选择、路径、回撤*/
	pathNum := []int{}
	start := 1
	res := [][]int{}
	combineBackTrack(pathNum, k, start, n, &res)

	return res
}

/*关于引用，如果不修改内部的数据，可以直接传*/
/*如果要修复引用数据的内部数据，一定要传地址，不然无法更改*/
func combineBackTrack(pathNum []int, k, start, n int, res *[][]int, ) {

	if len(pathNum) == k {
		temp := make([]int, k)
		copy(temp, pathNum)
		*res = append(*res, temp)
		return
	}

	for i := start; i <= n; i++ {
		pathNum = append(pathNum,i)

		combineBackTrack(pathNum,k,i+1,n,res)

		pathNum = pathNum[:len(pathNum)-1]
	}

}
