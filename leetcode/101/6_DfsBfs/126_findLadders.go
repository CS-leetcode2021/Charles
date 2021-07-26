package __DfsBfs

/**
 *  @ClassName:126_findLadders
 *  @Description:单次接龙2
 *  @Author:jackey
 *  @Create:2021/7/26 下午5:17
 */

// bfs||双端bfs优化
// 需要图的结合！！！未完待续....

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	res := make([][]string, 0)
	// 创建map
	wordMap := make(map[string]bool)
	for _, w := range wordList {
		wordMap[w] = true
	}

	if !wordMap[endWord] {
		return res
	}

	// 创建queue，这里为什么是二位数组？？？
	queue := make([][]string, 0)
	queue = append(queue, []string{beginWord})

	// queueLen is used to track how many slices in queue are in the same level
	queueLen := 1

	// visited
	levelMap := make(map[string]bool)
	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]
		lastWord := path[len(path)-1]

		for i := 0; i < len(lastWord); i++ {
			for c := 'a'; c <= 'z'; c++ { // 26个字母变换，每次只变换一次
				nextWord := lastWord[:i] + string(c) + lastWord[i+1:]
				if nextWord == endWord {
					path = append(path, endWord)
					res = append(res, path)
					continue
				}
				if wordMap[nextWord] {
					levelMap[nextWord] = true
					newPath := make([]string, len(path))

					copy(newPath, path)
					newPath = append(newPath, nextWord)
					queue = append(queue, newPath)
				}
			}
		}
		queueLen--

		if queueLen == 0 {
			if len(res) > 0 {
				return res
			}

			for k, _ := range levelMap {
				delete(wordMap,k)
			}

			levelMap = make(map[string]bool)
			queueLen = len(queue)
		}
	}
	return res
}
