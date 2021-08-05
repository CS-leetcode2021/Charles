package interview

import "fmt"

/**
 *  @ClassName:tmp_print
 *  @Description:go语言N个线程交替打印100
 *  @Author:jackey
 *  @Create:2021/8/5 下午10:19
 */

func print() {
	chanNum := 3
	chanQueue := make([]chan int, chanNum) // channel

	res := 0

	exitChan := make(chan bool) // 退出标志

	for i := 0; i < chanNum; i++ {
		chanQueue[i] = make(chan int)
		if i == chanNum-1 {
			go func(i int) {
				chanQueue[i] <- 1
			}(i)
		}
	}

	for i := 0; i < chanNum; i++ {
		lastChan, curChan := make(chan int), make(chan int)

		if i == 0 {
			lastChan = chanQueue[chanNum-1]
		} else {
			lastChan = chanQueue[i-1]
		}

		curChan = chanQueue[i]

		go func(i int, lastChan, curChan chan int) {
			for {
				if res > 100 {
					exitChan <- true
				}
				<-lastChan
				fmt.Printf("thread%d: %d \n", i, res)
				res++
				curChan <- 1
			}
		}(i, lastChan, curChan)

	}
	<-exitChan
	fmt.Println("done")

}

func main() {
	print()
}