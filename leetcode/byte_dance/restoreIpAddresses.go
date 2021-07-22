package main

import "strconv"

/**
 *  @ClassName:restoreIpAddresses
 *  @Description:字节 93 复原IP地址；解决方法：回溯；回溯会穷举所有的节点，通常用于解决找出所有的可能的组合
 *  【https://www.cis.upenn.edu/~matuszek/cit594-2012/Pages/backtracking.html】
 *  @Author:jackey
 *  @Create:2021/6/21 下午8:51
 */

const SegCount = 4

var (
	ans []string
	segments []int
)

func restoreIP(s string) []string {
	segments = make([]int,SegCount)
	ans = []string{}
	dfs(s,0,0)
	return ans
}

func dfs(s string, segID, segStart int)  {
	// 如果找到了4段IP地址，并且遍历完了所有的字符串，那么就是一种答案，将答案进行拼接即可

	if segID == SegCount {	// 等于4段ID
		if segStart == len(s) {
			ipAddr := ""
			for i := 0; i < SegCount; i++ {	// 进行IP拼接
				ipAddr += strconv.Itoa(segments[i])
				if i != SegCount-1 {
					ipAddr += "."
				}
			}
			ans = append(ans,ipAddr)	// 所有的结果
		}
		return
	}

	// 还未寻到第四段，便已经是遍历结束
	if segStart == len(s) {
		return
	}

	if s[segStart] == '0' {	// 如果是0,那么该段便是0
		segments[segID] = 0
		dfs(s,segID+1,segStart+1)
	}
	// 一般情况，枚举每一种可能性并递归
	addr := 0
	for segEnd := segStart; segEnd < len(s); segEnd++ {
		addr = addr * 10 + int(s[segEnd] - '0')
		if addr > 0 && addr <= 0xFF {
			segments[segID] = addr
			dfs(s, segID + 1, segEnd + 1)
		} else {
			break
		}
	}

}