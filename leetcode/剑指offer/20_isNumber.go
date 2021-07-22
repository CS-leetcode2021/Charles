package main

// 剑指offer-20 表示数值的字符串

func isNumber(s string) bool {
	if len(s) == 0 {
		return false
	}
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			s = s[i:]
			break
		}
		if i == len(s)-1{
			return false
		}
	}
	numFlag := false
	dotFlag := false
	eFlag := false

	space := false
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' &&!space {
			numFlag = true
		} else if s[i] == '.' && !dotFlag && !eFlag &&!space {
			dotFlag = true
		} else if (s[i] == 'e' || s[i] == 'E') && !eFlag && numFlag &&!space {
			eFlag = true
			numFlag = false
		}else if (s[i] == '+' || s[i] == '-') && (i == 0 || s[i-1] == 'e' || s[i-1] == 'E'&&!space ) {
			continue
		} else if s[i] == ' ' &&  numFlag {
			space = true
			if i == len(s)-1 {
				return true
			}
		}else {
			return false
		}
	}
	return numFlag
}