package daily

import (
	"strings"
)

/**
 *  @ClassName:71_simplifyPath
 *  @Description:TODO
 *  @Author:jackey
 *  @Create:2021/8/20 下午8:45
 */

func simplifyPath(path string) string {
	// 有一个内置函数
	// path.Clean()
	n := len(path)
	if n == 0 {
		return ""
	}

	sts := strings.Split(path, "/")

	stark := make([]string, 0)
	for i := 0; i < len(sts); i++ {
		if sts[i] == " " || sts[i] == "." {
			continue
		} else if sts[i] == ".." {
			if len(stark) > 0 {
				stark = stark[:len(stark)-1]
			}
		}else {
			stark = append(stark,sts[i])
		}
	}
	res := ""
	if len(stark) == 0 {
		return "/"
	}
	for i := 0; i < len(stark); i++ {
		res += "/"
		res += stark[i]
	}
	return res
}