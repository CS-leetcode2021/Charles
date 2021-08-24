package dp21

/**
 *  @ClassName:common
 *  @Description:TODO
 *  @Author:jackey
 *  @Create:2021/8/24 上午11:41
 */

func Max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func Min(i, j int) int {
	if i > j {
		return j
	}
	return i
}
