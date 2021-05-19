package string

/**
 *  @ClassName:	string.go
 *  @Description:TODO
 *  @Author:	Jack
 *  @Create:2021/3/9 下午7:24
 */

func removeDuplicates(s string) string {
	stack := []byte{}

	for i := range s {
		if len(stack) > 0 && stack[len(stack)-1] == s[i] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}

	return string(stack)
}
