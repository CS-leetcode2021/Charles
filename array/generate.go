package array

/**
 *  @ClassName:generate
 *  @Description:leetcode 118 杨辉三角
 *  @Author:jackey
 *  @Create:2021/5/20 下午12:49
 */

/*
 *  @Description:  	1、第i行有i+1个数据；2、第i行的第j个数据等于第i-1行的第 j-1 个数据加上 j 个数据
 *  @Param:			行数
 *  @Return:		二维数组
 */

func generate(numRows int)[][]int  {

	res := make([][]int,numRows)

	for i := range res{
		res[i] = make([]int,i+1)

		res[i][0] = 1
		res[i][i] = 1

		for j := 1; j < len(res[i])-1; j++ {
			res[i][j]= res[i-1][j-1] + res[i-1][j]
		}
	}
	return res
}
