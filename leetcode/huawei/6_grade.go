package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/**
 *  @ClassName:6_grade
 *  @Description:TODO
 *  @Author:jackey
 *  @Create:2021/7/20 下午6:05
 */

// 86 3932
func main() {
	for {
		var m, n int

		_, err := fmt.Scan(&n, &m)
		if err != nil {
			break
		}

		var scores []int
		for i := 0; i < n; i++ {
			var x int
			if _, err := fmt.Scan(&x); err != nil {
				panic(err)
			}
			scores = append(scores, x)
		}

		for i := 0; i < m; i++ {
			var op string
			var a, b int
			if _, err := fmt.Scan(&op, &a, &b); err != nil {
				panic(err)
			}
			switch op {
			case "Q":
				handQuery(scores, a, b)
			case "U":
				handUpdate(scores, a, b)
			}
		}
	}
}

func handQuery(scores []int, i, j int) {
	if i > j { // 考虑大于的问题
		i, j = j, i
	}
	i, j = i-1, j-1
	max := scores[i]
	for k := i + 1; k <= j; k++ {
		if scores[k] > max {
			max = scores[k]
		}
	}
	fmt.Println(max)
}

func handUpdate(scores []int, i, j int) {
	scores[i-1] = j
}

// 481ms 6176kb
func main2() {
	var N int
	var M int
	for {
		str1, str2 := "", ""
		n, _ := fmt.Scanln(&str1, &str2)

		if n == 0 {
			os.Exit(0)
		}
		N, _ = strconv.Atoi(str1)
		M, _ = strconv.Atoi(str2)
		studentG := make([]int, N)

		for i := 0; i < N; i++ {
			fmt.Scan(&studentG[i])
		}

		op := make([]string, M)
		opA := make([]int, M)
		opB := make([]int, M)
		for i := 0; i < M; i++ {
			fmt.Scanln(&op[i], &opA[i], &opB[i])
		}

		for i := 0; i < M; i++ {
			if op[i] == "Q" {
				start := opA[i]
				end := opB[i]
				if start > end {
					start, end = end, start
				}
				tmp := make([]int, end-start+1)
				copy(tmp, studentG[start-1:end])
				sort.Ints(tmp)
				fmt.Println(tmp[len(tmp)-1])
			} else { // B为成绩

				Aid := opA[i]
				GB := opB[i]
				studentG[Aid-1] = GB
			}
		}
	}
}


func main3(){
	var N,M int
	input:=bufio.NewScanner(os.Stdin)
	bs:=make([]byte,40000*4)
	input.Buffer(bs,len(bs))
	for input.Scan(){
		s:=strings.Split(input.Text()," ")
		N,_=strconv.Atoi(s[0])
		M,_=strconv.Atoi(s[1])
		scores:=make([]int,N)
		//fmt.Println(N,M)
		input.Scan()
		s=strings.Split(input.Text()," ")

		for i:=0;i<N;i++{
			scores[i],_=strconv.Atoi(s[i])
		}
		for i:=0;i<M;i++{
			input.Scan()
			//fmt.Println(input.Text())
			str:=strings.Split(input.Text()," ")
			if str[0]=="Q"{
				n1,_:=strconv.Atoi(str[1])
				n1--
				n2,_:=strconv.Atoi(str[2])
				n2--
				if n1>n2{
					n2,n1=n1,n2
				}
				max:=scores[n1]
				for j:=n1+1;j<=n2&&j<N;j++{
					if scores[j]>max{
						max=scores[j]
					}
				}
				fmt.Println(max)
			}else{
				n1,_:=strconv.Atoi(str[1])
				n2,_:=strconv.Atoi(str[2])
				scores[n1-1]=n2
			}
		}
	}

}