package main

import "fmt"

func main01()  {
	x := 34
	//
	fmt.Printf("%b\n",x)
	//
	//x &= 0x1
	//
	//fmt.Printf("%b\n",x)

	x = b(x)
	fmt.Printf("%b\n",x)
}

func b(x int) int {
	val := 0

	for i := 64; i != 0; i-- {
		val = (val << 1) | (x & 0x1)
		x >>= 1
	}
	return val
}

func main() {
	sum := 0

	for i := 0; i < 10; i++ {
		if (i & 1 )!= 0 {
			fmt.Println(i)
			fmt.Printf("%b\n",i)
			continue
		}
		sum +=  i
	}

	fmt.Println(sum)
}