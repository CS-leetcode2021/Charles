package daily

import "strings"

/**
 *  @ClassName:408_validIPAddress
 *  @Description:TODO
 *  @Author:jackey
 *  @Create:2021/8/20 下午10:20
 */

// 100/92
func validIPAddress(IP string) string {
	n := len(IP)
	if n == 0 {
		return "Neither"
	}
	for i := 0; i < n; i++ {
		if IP[i] == '.' {
			sts := strings.Split(IP, ".")
			m := len(sts)
			if m != 4 {
				return "Neither"
			}
			for j := 0; j < 4; j++ {
				tag := Judge4(sts[j])
				if !tag {
					return "Neither"
				}
			}
			return "IPv4"
		}
	}

	sts := strings.Split(IP, ":")
	m := len(sts)
	if m != 8 {
		return "Neither"
	}
	for i := 0; i < 8; i++ {
		tag := Judge6(sts[i])
		if !tag {
			return "Neither"
		}
	}
	return "IPv6"
}

func Judge4(s string) bool {
	if len(s) == 0 || len(s) > 3 {
		return false
	}
	if len(s) == 1 {
		return true
	}
	res := 0
	opt := 100
	for i := 0; i < len(s); i++ {
		if i == 0 && s[i] == '0' {
			return false
		}
		tmp := int(s[i] - '0')
		res += tmp * opt
		opt /= 10
	}

	if res >= 256 {
		return false
	}
	return true
}

func Judge6(s string) bool {
	if len(s) == 0 || len(s) > 4 {
		return false
	}

	for i := 0; i < len(s); i++ {
		if s[i] > 'f' && s[i] <= 'z' {
			return false
		}
		if s[i] > 'F' && s[i] <= 'Z' {
			return false
		}
	}
	return true

}
