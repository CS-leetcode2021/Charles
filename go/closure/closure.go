package main

import (
	"fmt"
	"time"
)

/**
 *  @ClassName:closure
 *  @Description:闭包
 *  @Author:jackey
 *  @Create:2021/8/9 下午3:45
 */

func main() {
	go foo6()
	foo6Chan <- 1
	foo6Chan <- 2
	foo6Chan <- 3
	foo6Chan <- 5
	time.Sleep(time.Second * 1)
	fmt.Println("-----------------------------------")
	// 纳秒
	foo6Chan <- 11
	time.Sleep(time.Duration(1) * time.Nanosecond)
	foo6Chan <- 12
	time.Sleep(time.Duration(1) * time.Nanosecond)
	foo6Chan <- 13
	time.Sleep(time.Duration(1) * time.Nanosecond)
	foo6Chan <- 15
	time.Sleep(time.Second * 1)

	fmt.Println("-----------------------------------")
	// 微秒
	foo6Chan <- 21
	time.Sleep(time.Duration(1) * time.Microsecond)
	foo6Chan <- 22
	time.Sleep(time.Duration(1) * time.Microsecond)
	foo6Chan <- 23
	time.Sleep(time.Duration(1) * time.Microsecond)
	foo6Chan <- 25
	time.Sleep(time.Second * 1)

	fmt.Println("-----------------------------------")

	// 毫秒
	foo6Chan <- 31
	time.Sleep(time.Duration(1) * time.Millisecond)
	foo6Chan <- 32
	time.Sleep(time.Duration(1) * time.Millisecond)
	foo6Chan <- 33
	time.Sleep(time.Duration(1) * time.Millisecond)
	foo6Chan <- 35
	time.Sleep(time.Second * 1)

	fmt.Println("-----------------------------------")
	foo6Chan <- 41
	time.Sleep(time.Duration(1) * time.Second)
	foo6Chan <- 42
	time.Sleep(time.Duration(1) * time.Second)
	foo6Chan <- 43
	time.Sleep(time.Duration(1) * time.Second)
	foo6Chan <- 45
	time.Sleep(time.Second * 1)

	close(foo6Chan)

	foo9()
	time.Sleep(time.Second*1)
}

func main03() {
	x := 133

	f1 := foo1(&x)
	f2 := foo2(x)

	f1()
	f2()
	f1()
	f2()

	fmt.Println("--------------------")

	x = 233
	f1()
	f2()
	f1()
	f2()

	fmt.Println("--------------------")

	foo1(&x)()
	foo2(x)()
	foo1(&x)()
	foo2(x)()
	foo2(x)()

	fmt.Println("-------------------")
	f7s := foo7(11)

	for _, f7 := range f7s {
		f7()
	}

	fmt.Println("--------------------")
	f8 := foo8()
	f8()

}

func main02() {

	foo3()

	fmt.Println("------------------------------")

	foo4()

	fmt.Println("------------------------------")
	foo5()

	time.Sleep(time.Second * 2)

	fmt.Println("------------------------------")

}

func main01() {

	x := 133

	for i := 0; i < 3; i++ {
		f1 := foo1(&x)
		fmt.Printf("main: val = %d \n", x)
		f1()
		fmt.Printf("main: val = %d \n", x)
		fmt.Println("----------------------")
		f2 := foo2(x)
		fmt.Printf("main: val = %d \n", x)
		f2()
		fmt.Printf("main: val = %d \n", x)
		fmt.Println("----------------------")
	}

	fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!test2")
	x = 233

	for i := 0; i < 3; i++ {
		f1 := foo1(&x)
		fmt.Printf("main: val = %d \n", x)
		f1()
		fmt.Printf("main: val = %d \n", x)
		fmt.Println("----------------------")
		f2 := foo2(x)
		fmt.Printf("main: val = %d \n", x)
		f2()
		fmt.Printf("main: val = %d \n", x)
		fmt.Println("----------------------")
	}

}

// foo1 case 1：
func foo1(x *int) func() {
	return func() {
		*x = *x + 1
		fmt.Printf("fool val = %d \n", *x)
	}
}

// case 2: foo2
func foo2(x int) func() {
	return func() {
		x = x + 1
		fmt.Printf("foo2 val = %d \n", x)
	}
}

// case 3:

func foo3() {
	values := []int{1, 2, 3, 5}
	for _, val := range values {
		fmt.Printf("foo3 val = %d \n", val)
	}
}

//case 4:

func show(v interface{}) {
	fmt.Printf("foo4 val = %v \n", v)
}

func foo4() {
	values := []int{1, 2, 3, 5}
	for _, val := range values {
		go show(val)
	}
}

// case 5:

func foo5() {
	values := []int{1, 2, 3, 5}
	for _, val := range values {
		// 匿名函数.在这里也是闭包
		// 在我们调用go func() { xxx }()的时候，只要没有真正开始执行这段代码，那它还只是一段函数声明。而在这段匿名函数被执行的时候，才是内部变量寻找真正赋值的时候。
		go func() {
			fmt.Printf("foo5 val = %v \n", val)
		}()
	}
}

// case 6:

var foo6Chan = make(chan int, 10)

func foo6() {
	for val := range foo6Chan {
		go func() {
			fmt.Printf("foo6 val = %d \n", val)
		}()
	}
}

// case 7:

func foo7(x int) []func() {
	var fs []func()

	values := []int{1, 2, 3, 5}
	for _, val := range values {
		fs = append(fs, func() {
			fmt.Printf("foo7 val = %d \n", x+val)
		})
	}

	return fs
}

// case 8:
// 闭包的延迟绑定
// 闭包的神奇支出，他会保存相关引用的环境，也就是说，val这个变量在闭包内的生命周期得到了保证，因此在执行这个闭包的时候
// 会去外部环境寻找最新的数值
// 本来局部变量会在函数结束退出时销毁，但是现在不会。
func foo8() func() {
	x := 1
	f := func() {
		fmt.Printf("foo8 val = %d \n", x)
	}

	x = 11
	return f
}

func foo9() {
	for i := 1; i < 10; i++ {
		curTime := time.Now().UnixNano()
		go func(t1 int64) {
			t2 := time.Now().UnixNano()
			fmt.Printf("foo8 ts = %d us \n", t2-t1)
		}(curTime)
	}
}