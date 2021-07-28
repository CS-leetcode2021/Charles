package _09

/**
 *  @ClassName:1_isUnique
 *  @Description:TODO
 *  @Author:jackey
 *  @Create:2021/7/28 下午2:49
 */

func isUnique(astr string) bool {
	tmap := make(map[byte]int)

	for i := 0; i < len(astr); i++ {
		if _, ok := tmap[astr[i]]; !ok {
			tmap[astr[i]] = 1
		} else {
			return false
		}
	}
	return true
}

// 使用位运算呢

func isUnique2(astr string) bool {
	num := 0
	for _, v := range astr {
		moveBit := v - 'a'
		if num&(1<<moveBit) != 0 { // 证明相同的位置出现过
			return false
		} else { // 亦或
			num = num | (1 << moveBit)
		}
	}
	return true
}
