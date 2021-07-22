package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var A, B int
	input := bufio.NewScanner(os.Stdin)
	bs := make([]byte,40000*4)

	input.Buffer(bs,len(bs))
	for input.Scan() {

		s := strings.Split(input.Text(), " ")

		A, _ = strconv.Atoi(s[0])
		B, _ = strconv.Atoi(s[1])

		fmt.Println(A+B)

	}
}
