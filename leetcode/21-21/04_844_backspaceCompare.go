package _1_21

/**
 *  @ClassName:04_844_backspaceCompare
 *  @Description:比较含退格的字符串
 *  @Author:jackey
 *  @Create:2021/8/10 下午9:15
 */

// 拷贝一下进行比较
func backspaceCompare(s string, t string) bool {

	str1, str2 := make([]byte, 0), make([]byte, 0)

	for i := 0; i < len(s); i++ {
		if s[i] != '#' {
			str1 = append(str1, s[i])
		} else {
			if len(str1) > 0 {
				str1 = str1[:len(str1)-1]
			}
			// len(str1) == 0不做处理
		}
	}

	for i := 0; i < len(t); i++ {
		if t[i] != '#' {
			str2 = append(str2, t[i])
		} else {
			if len(str2) > 0 {
				str2 = str2[:len(str2)-1]
			}
		}
	}
	if string(str1) == string(str2) {
		return true
	}
	return false

}

// 优化 function

func backspaceCompare2(s string, t string) bool {

	str1 := opt(s)
	str2 := opt(t)
	if str1 == str2 {
		return true
	}
	return false

}

func opt(s string) string {
	strs := make([]byte, 0)

	for i := 0; i < len(s); i++ {
		if s[i] != '#' {
			strs = append(strs, s[i])
		} else {
			if len(strs) > 0 {
				strs = strs[:len(strs)-1]
			}
			// len(str1) == 0不做处理
		}
	}
	return string(strs)
}

// 优化， 直接操作
func build(str string) string {
	s := []byte{}
	for i := range str {
		if str[i] != '#' {
			s = append(s, str[i])
		} else if len(s) > 0 {
			s = s[:len(s)-1]
		}
	}
	return string(s)
}

func backspaceCompare3(s, t string) bool {
	return build(s) == build(t)
}

// 双指针
func backspaceCompare4(s, t string) bool {
	skipS, skipT := 0, 0
	i, j := len(s)-1, len(t)-1
	for i >= 0 || j >= 0 {
		for i >= 0 {
			if s[i] == '#' {
				skipS++
				i--
			} else if skipS > 0 {
				skipS--
				i--
			} else {
				break
			}
		}
		for j >= 0 {
			if t[j] == '#' {
				skipT++
				j--
			} else if skipT > 0 {
				skipT--
				j--
			} else {
				break
			}
		}
		if i >= 0 && j >= 0 {
			if s[i] != t[j] {
				return false
			}
		} else if i >= 0 || j >= 0 {
			return false
		}
		i--
		j--
	}
	return true
}
