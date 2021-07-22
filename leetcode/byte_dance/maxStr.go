package main

import "fmt"

/**
 *  @ClassName:maxStr.go
 *  @Description:最长公共子序列 acwing-897、leetcode-1143
 *  @Author:jackey
 *  @Create:2021/6/15 上午11:17
 */

func main() {
	m, n := 0,0
	
	fmt.Scan(&m,&n)
	
	mStr, nStr := " "," "

	fmt.Scanln(&mStr)
	fmt.Scanln(&nStr)


	fmt.Println(mStr,nStr)

	// 创建二维数组
	res := make([][]int,m+1)
	for i := 0; i <= m; i++ {
		res[i] = make([]int,n+1)
	}

	// dp
	for i := 1; i <= m; i++ {
		for j := 1; j <= n ; j++ {
			res[i][j] = MaxStr(res[i-1][j],res[i][j-1])
			if mStr[i-1] == nStr[j-1] {
				res[i][j] = MaxStr(res[i][j],res[i-1][j-1]+1)
			}
		}
	}

	fmt.Println(res[m][n])

}

func MaxStr(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// leetcode
func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1),len(text2)

	// 创建二维数组
	res := make([][]int,m+1)
	for i := 0; i <= m; i++ {
		res[i] = make([]int,n+1)
	}

	// dp
	for i := 1; i <= m; i++ {
		for j := 1; j <= n ; j++ {
			res[i][j] = MaxStr(res[i-1][j],res[i][j-1])
			if text1[i-1] == text2[j-1] {
				res[i][j] = MaxStr(res[i][j],res[i-1][j-1]+1)
			}
		}
	}

	return res[m][n]
}