package _1

/**
 *  @ClassName:13_191_hammingWeight
 *  @Description:判断二进制中出现的个数
 *  @Author:jackey
 *  @Create:2021/8/3 下午7:59
 */

// n *= n-1
func hammingWeight(num uint32) (ones int) {
	for ; num > 0; num &= num - 1 {
		ones++
	}
	return
}

// 与2的i次方进行&运算
func hammingWeight2(num uint32) (ones int) {
	for i := 0; i < 32; i++ {
		if 1<<i&num > 0 {
			ones++
		}
	}
	return
}
