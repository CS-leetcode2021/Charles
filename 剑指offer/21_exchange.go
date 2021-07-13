package main

/**
 *  @ClassName:21_exchange
 *  @Description:剑指offer-21 调整数组顺序使奇数位于偶数的前面
 *  @Author:jackey
 *  @Create:2021/7/13 下午5:31
 */


// 单指针滑动解法 96/70
func exchange(nums []int) []int {
	n := len(nums)
	first := 0

	for i := 0; i < n; i++ {
		if nums[i]&1 == 1 { // 奇数才需要调整
			tmp := nums[first]
			nums[first] = nums[i]
			nums[i] = tmp
			first++
		}
	}
	return nums
}

//优化过以后 99/87
func exchange2(nums []int) []int {
	n := len(nums)
	first := 0

	for i := 0; i < n; i++ {
		if nums[i]&1 == 1 { // 奇数才需要调整
			nums[first],nums[i] = nums[i],nums[first]
			first++
		}
	}
	return nums
}

func exchange3(nums []int) []int {
	if len(nums) <= 0 {
		return nil
	}

	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		if nums[i]&1 == 0 {
			for nums[j]&1 != 1{
				j--
				if j == 0 {
					break
				}
			}
			nums[i],nums[j] = nums[j],nums[i]
		}
	}
	return nums
}