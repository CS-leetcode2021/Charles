package dp21

/**
 *  @ClassName:11_263_isUgly
 *  @Description:TODO
 *  @Author:jackey
 *  @Create:2021/8/26 下午1:57
 */


var factors = []int{2, 3, 5}

func isUgly(n int) bool {
	if n <= 0 {
		return false
	}
	for _, f := range factors {
		for n%f == 0 {
			n /= f
		}
	}
	return n == 1
}