package main

import (
	"sort"
)

/**
 *  @ClassName:38_permutation
 *  @Description:剑指offer-38 字符串排列
 *  @Author:jackey
 *  @Create:2021/7/19 下午3:02
 */

// 回溯
func permutation(s string)(ans []string){
	t := []byte(s)

	sort.Slice(t, func(i, j int) bool {
		return t[i] < t[j]
	})

	n := len(s)
	perm := make([]byte,0,n)
	vis := make([]bool,n)
	var backtrack func(int)
	backtrack = func(i int) {
		if i == n {
			ans = append(ans, string(perm))
			return
		}
		for j, b := range vis {
			if b || j > 0 && !vis[j-1] && t[j-1] == t[j] {
				continue
			}
			vis[j] = true
			perm = append(perm, t[j])
			backtrack(i + 1)
			perm = perm[:len(perm)-1]
			vis[j] = false
		}
	}
	backtrack(0)
	return
}
