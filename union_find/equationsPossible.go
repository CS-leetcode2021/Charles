package union_find

/**
 *  @ClassName:equationsPossible
 *  @Description:leetcode 990. 等式方程的可满足性
 *  @Author:jackey
 *  @Create:2021/4/1 下午9:52
 */

func equationsPossible(equations []string) bool {
	count, parent, size := uf(26)
	for _, str := range equations {
		if str[1] == '=' {
			x := str[0]
			y := str[3]

			union(int(x-'a'),int(y-'a'),count,parent,size)
		}
	}

	for _, str := range equations {
		if str[1] == '!' {
			x := str[0]
			y := str[3]

			if connected(int(x-'a'), int(y-'b'),parent) {
				return false
			}
		}
	}


	return true
}
