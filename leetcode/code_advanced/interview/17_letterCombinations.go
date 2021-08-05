package interview

import "fmt"

/**
 *  @ClassName:17_letterCombinations
 *  @Description:电话号码的字母组合
 *  @Author:jackey
 *  @Create:2021/8/5 下午9:11
 */

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	tmap := map[int]string{
		2: "abc",
		3: "def",
		4: "ghi",
		5: "jkl",
		6: "mno",
		7: "pqrs",
		8: "tuv",
		9: "wxyz",
	}

	path := make([]byte, 0)
	res := make([]string, 0)

	backtracking17(digits, path, &res, &tmap, 0)
	return res
}

func backtracking17(digits string, path []byte, res *[]string, tmap *map[int]string, k int) {
	if len(path) == len(digits) {
		// 当前path符合要求
		tmp := make([]byte, len(digits))
		copy(tmp, path)
		*res = append(*res, string(tmp))
		return
	}

	nums := int(digits[k]-'0')
	strs := (*tmap)[nums]
	for j := 0; j < len(strs); j++ {
		path = append(path, strs[j])
		backtracking17(digits, path, res, tmap, k+1)
		path = path[:len(path)-1]
	}
	return
}

func main() {
	tmp := ""
	res := letterCombinations(tmp)
	fmt.Println(res)
}