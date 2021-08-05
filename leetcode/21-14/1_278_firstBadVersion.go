package _1_14

/**
 *  @ClassName:1_firstBadVersion
 *  @Description:第一个错误的版本 278
 *  @Author:jackey
 *  @Create:2021/7/26 上午10:33
 */

// 100/100
// 二分查找

func firstBadVersion(n int) int {
	return binSearch278(n, 1, n)
}

func binSearch278(n, l, r int) int {
	for l <= r {
		mid := (r + l) >> 1
		if isBadVersion(mid) { // 当前版本是错误的
			if mid == 1 || !isBadVersion(mid-1) { // 当前的前一个版本不是错误的
				return mid
			}
			// 前一个版本也是错误的
			r = mid - 1
		} else { // 当前版本是正确的
			if mid == n {
				return -1
			}
			l = mid + 1
		}
	}
	return -1
}

func isBadVersion(version int) bool {
	return false
}
