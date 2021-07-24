package __sort

/**
 *  @ClassName:75_sortColors
 *  @Description:不同的颜色排序归类
 *  @Author:jackey
 *  @Create:2021/7/24 下午6:12
 */

// 要求：原地排序

// 3for
func sortColors(nums []int)  {

	tag := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums[tag],nums[i] = nums[i],nums[tag]
			tag++
		}
	}

	for i := tag; i < len(nums); i++ {
		if nums[i] == 1 {
			nums[tag],nums[i] = nums[i],nums[tag]
			tag++
		}
	}
}

func sortColors2(nums []int) {
	// 双指针，一次遍历 : 存在边界问题
	tag0 , tag1 := 0,0
	for i := 0; i <len(nums); i++{
		if nums[i] == 0 {
			nums[i],nums[tag0] = nums[tag0],nums[i]
			tag0++
			if tag0 < tag1{
				nums[i],nums[tag1] = nums[tag1],nums[i]
			}
			tag1++
		}else if nums[i] == 1{
			nums[i],nums[tag1] = nums[tag1],nums[i]
			tag1++
		}

	}

}


// 双指针一左一右
func sortColors3(nums []int) {
	p0, p2 := 0, len(nums)-1
	for i := 0; i <= p2; i++ {
		// 换回来的可能还是2
		for ; i <= p2 && nums[i] == 2; p2-- {
			nums[i], nums[p2] = nums[p2], nums[i]
		}
		if nums[i] == 0 {
			nums[i], nums[p0] = nums[p0], nums[i]
			p0++
		}
	}
}

// 三路快排排
func sortColors4(nums []int)  {
	zero, two := -1, len(nums)
	for i := 0; i < two; {
		if nums[i] == 1 {
			i++
		} else if nums[i] == 2 {
			two -= 1
			nums[i], nums[two] = nums[two], nums[i]
		} else {
			zero += 1
			nums[i], nums[zero] = nums[zero], nums[i]
			i++
		}
	}
}

