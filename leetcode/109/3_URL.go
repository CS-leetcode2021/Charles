package _09

/**
 *  @ClassName:3_URL
 *  @Description:URL化
 *  @Author:jackey
 *  @Create:2021/8/6 下午3:58
 */

func replaceSpaces(S string, length int) string {
	strs := make([]byte, 0)

	for i := 0; i < length; i++ {
		if S[i] == ' ' {
			tmp := []byte{'%', '2', '0'}
			strs = append(strs, tmp...)
		} else {
			strs = append(strs, S[i])
		}
	}
	return string(strs)
}
