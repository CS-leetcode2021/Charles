package __DfsBfs

/**
 *  @ClassName:127_ladderLength
 *  @Description:单次接龙1
 *  @Author:jackey
 *  @Create:2021/7/26 下午7:51
 */

// 1、从beginWord开始，每次变换一个字母
// 2、变换后的字母到wordList中去查找是否存在，可以使用map记录，效率高
// 3、如果存在则进入队列，使用bfs，因为是让你返回最短的路径问题
// 4、入队以后继续按照bfs进行遍历，直至队列长度为0

func ladderLength(beginWord string, endWord string, wordList []string) int {
	// 创建wordMap
	wordMap := getWordMap(wordList, beginWord)
	// 创建队列
	que := []string{beginWord}
	// 标记长度
	depth := 0

	for len(que) > 0 {
		depth++
		// 当前长度
		curLen := len(que)

		// 出队列
		for i := 0; i < curLen; i++ {
			// 出队列操作
			curWord := que[0]
			que = que[1:]

			// 需要对每一个字母进行26中变换，再判断是否又符合wordlist的进队列
			candidates := getCandidates(curWord)

			for _, candidate := range candidates {
				if _, ok := wordMap[candidate]; ok {
					// 查看是否是结束字符
					if candidate == endWord {
						return depth+1
					}
					delete(wordMap,candidate)
					que = append(que,candidate)
				}
			}
		}
	}
	return 0
}

// 创建map
func getWordMap(wordList []string, beginWord string) map[string]int {
	wordMap := make(map[string]int)
	for _, str := range wordList {
		if i, ok := wordMap[str]; !ok {
			if str != beginWord {
				wordMap[str] = i
			}
		}
	}
	return wordMap
}

// 创建所有的变换形式
func getCandidates(word string) []string {
	var res []string

	for i := 0; i < 26; i++ {
		for j := 0; j < len(word); j++ {
			if word[j] != byte(int('a')+i) {
				// 每次改变一个字母 // 时间成本是不是特别大？
				res = append(res, word[:j]+string(int('a')+i)+word[j+1:])
			}
		}
	}
	return res
}
