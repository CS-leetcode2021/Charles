package main

/**
 *  @ClassName:candy.go
 *  @Description:135 分发糖果
 *  @Author:jackey
 *  @Create:2021/6/7 下午7:17
 */

/*
 *  @Description:   贪心算法
 *  @Param:         数组
 *  @Return:        int
 */
/*
 *  @Description:   基本步骤：1,从某个初始解出发；
 *							2,采用迭代的过程，当可以向目标进一步的时，就根据局部最优策略，得到一部分解，缩小问题规模
 *  			    		3，将所有的解综合起来
 *  @Return:
 */

func candy(rating []int) int {
	if len(rating) == 0 {
		return 0
	}
	if len(rating) == 1 {
		return 1
	}


	return 0
}
