package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/**
 *  @ClassName:7_errRecord
 *  @Description:TODO
 *  @Author:jackey
 *  @Create:2021/7/20 下午7:23
 */
type cnt struct {
	index int
	cnt   int
}

// 4ms 956kb
func main() {
	input := bufio.NewScanner(os.Stdin)
	m := map[string]*cnt{} // 映射
	index := 0
	for input.Scan() { // 判断是否还有输入
		index++
		str := strings.Split(input.Text(), " ") // 扫描第一行输入
		path := strings.Split(str[0], `\`)
		buf := bytes.NewBufferString(path[len(path)-1]) // 只要最后一个文件名
		buf.WriteString(" ")
		buf.WriteString(str[1])
		//strs := buf.String()
		//fmt.Println(strs)
		if _, ok := m[buf.String()]; ok {
			m[buf.String()].cnt += 1
		} else {
			m[buf.String()] = &cnt{index, 1}
		}
	}

	// 将map转化为数组
	keyarr := make([]string, len(m))
	i := 0
	for k, _ := range m {
		keyarr[i] = k
		i++
	}

	// 自定义排序法则
	sort.Slice(keyarr, func(i, j int) bool {
		if m[keyarr[i]].cnt > m[keyarr[j]].cnt {
			return true
		} else if m[keyarr[i]].cnt < m[keyarr[j]].cnt {
			return false
		} else {
			return m[keyarr[i]].index < m[keyarr[j]].index
		}
	})


	for i := 0; i < 8 && i < len(keyarr); i++ {
		s := strings.Split(keyarr[i], " ")
		c := m[keyarr[i]].cnt
		cs := strconv.Itoa(c)

		if len(s[0]) > 16 {
			st := (s[0])[len(s[0])-16:]
			fmt.Println(st + " " + s[1] + " " + cs)
		} else {
			fmt.Println(keyarr[i] + " " + cs)
		}
	}
}
