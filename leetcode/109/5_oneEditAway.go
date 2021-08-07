package _09

/**
 *  @ClassName:5_oneEditAway
 *  @Description:一次编辑
 *  @Author:jackey
 *  @Create:2021/8/6 下午4:40
 */

func oneEditAway(first string, second string) bool {
	// 字符串有两种类型，一种是长度相等，只要有一个字符不通就可以完成替换
	// 如果是长度差距为1,那就要删除或者插入一个字符事项相同

	diff := len(first) - len(second)

	if diff < 0 {
		diff = -diff
	}

	if diff > 1 {
		return false
	}

	m, n := len(first), len(second)
	p1, p2 := 0, 0
	tag := false
	for p1 < m && p2 < n {
		if first[p1] == second[p2] {
			p1++
			p2++
		} else {
			if tag {
				return false
			}
			tag = true
			if m > n {
				p1++
			}else if m == n {	// 可能长度相同但是只有一个字母是不同的
				p1++
				p2++
			}else{
				p2++
			}
		}
	}

	return true
}
