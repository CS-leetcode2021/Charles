package __sort

import (
	"bytes"
	"sort"
)

/**
 *  @ClassName:451_frequencySort
 *  @Description:根据字符出现频率排序
 *  @Author:jackey
 *  @Create:2021/7/24 下午4:56
 */


// 有一个案例没通过
func frequencySort(s string) string {

	if len(s) <= 1 {
		return s
	}
	tMap := make(map[byte]int)
	tVis := make([][2]byte,0)
	for i := 0; i < len(s); i++ {
		if _ , ok := tMap[s[i]]; !ok{
			tMap[s[i]] = 1
			tVis = append(tVis,[2]byte{s[i],' '})
		}else {
			tMap[s[i]]++
		}
	}

	for i := 0; i < len(tVis); i++ {
		// int转化为byte是有截断存在的，在超大数据的情况下，byte只会选择低8位，会忽略高位的数据，产生误差
		tVis[i][1] = byte(tMap[tVis[i][0]])
	}

	sort.Slice(tVis, func(i, j int) bool {
		return tVis[i][1] > tVis[j][1]
	})

	res := ""

	for i := 0; i < len(tVis); i++ {
		count := int(tVis[i][1])
		for count != 0 {
			res += string(tVis[i][0])
			count--
		}
	}

	return res
}



type ByteNode struct {
	Key   byte
	Count int
}

func frequencySort3(s string) string {
	cache := map[byte]*ByteNode{}
	array := make([]*ByteNode, 0)
	for i := range s {
		if _, ok := cache[s[i]]; ok {
			node := cache[s[i]]
			node.Count++
		} else {
			node := &ByteNode{
				Key:   s[i],
				Count: 1,
			}
			cache[s[i]] = node
			array = append(array, node)
		}
	}
	sort.Slice(array, func(i, j int) bool {
		return array[i].Count > array[j].Count
	})
	ans := ""
	for _, node := range array {
		for node.Count > 0 {
			ans += string(node.Key)
			node.Count--
		}
	}
	return ans
}

// 官方

func frequencySort2(s string) string {
	cnt := map[byte]int{}
	for i := range s {
		cnt[s[i]]++
	}

	type pair struct {
		ch  byte
		cnt int
	}
	pairs := make([]pair, 0, len(cnt))
	for k, v := range cnt {
		pairs = append(pairs, pair{k, v})
	}
	sort.Slice(pairs, func(i, j int) bool { return pairs[i].cnt > pairs[j].cnt })

	ans := make([]byte, 0, len(s))
	for _, p := range pairs {
		// 重复N次
		ans = append(ans, bytes.Repeat([]byte{p.ch}, p.cnt)...)
	}
	return string(ans)
}