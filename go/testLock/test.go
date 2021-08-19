package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"time"
)

/**
 *  @ClassName:test
 *  @Description:TODO
 *  @Author:jackey
 *  @Create:2021/8/18 下午3:01
 */

// doc
// 互斥锁

type mutex struct {
	state int32  // 当前的状态
	sema  uint32 // 用于控制锁状态
}

// int32中的不同为分别表示了不同的状态
// 低三位
// mutexLocked-锁定状态
// mutexWoken-从正常模式被唤醒
// mutexStarving-饥饿状态-1.9模式引入的
// 其余位
// waitersCount-当前互斥锁上等待的Goroutine个数

var mu sync.Mutex

// 调用 sync.Mutex.Lock()进入加锁状态
// 如果当前锁的状态是0（没有竞争），则可以直接通过CAS进行加锁置位1（改进的方案是mcs lock）
// 如果当前锁状态为1，则尝试通过自旋等方式等待锁的释放
// 自旋判定：
// 1、判定当前的goroutine能否进入自旋
// 2、通过自旋等待互斥锁的释放
// 3、计算互斥锁的最新状态
// 4、更新互斥锁的状态并获取锁

// 互斥锁的加锁过程比较复杂，它涉及自旋、信号量以及调度等概念：
//
//如果互斥锁处于初始化状态，会通过置位 mutexLocked 加锁；
//如果互斥锁处于 mutexLocked 状态并且在普通模式下工作，会进入自旋，执行 30 次 PAUSE 指令消耗 CPU 时间等待锁的释放；
//如果当前 Goroutine 等待锁的时间超过了 1ms，互斥锁就会切换到饥饿模式；
//互斥锁在正常情况下会通过 runtime.sync_runtime_SemacquireMutex 将尝试获取锁的 Goroutine 切换至休眠状态，等待锁的持有者唤醒；
//如果当前 Goroutine 是互斥锁上的最后一个等待的协程或者等待的时间小于 1ms，那么它会将互斥锁切换回正常模式；
//互斥锁的解锁过程与之相比就比较简单，其代码行数不多、逻辑清晰，也比较容易理解：
//
//当互斥锁已经被解锁时，调用 sync.Mutex.Unlock 会直接抛出异常；
//当互斥锁处于饥饿模式时，将锁的所有权交给队列中的下一个等待者，等待者会负责设置 mutexLocked 标志位；
//当互斥锁处于普通模式时，如果没有 Goroutine 等待锁的释放或者已经有被唤醒的 Goroutine 获得了锁，会直接返回；在其他情况下会通过 sync.runtime_Semrelease 唤醒对应的 Goroutine；

// RWMutex
// 读写互斥锁 sync.RWMutex 是细粒度的互斥锁，它不限制资源的并发读，但是读写、写写操作无法并行执行。

type RW struct {
	w sync.RWMutex
}

// 获取写锁时会先阻塞写锁的获取，后阻塞读锁的获取，这种策略能够保证读操作不会被连续的写操作『饿死』
// 虽然读写互斥锁 sync.RWMutex 提供的功能比较复杂，但是因为它建立在 sync.Mutex 上，所以实现会简单很多。我们总结一下读锁和写锁的关系：
//
//调用 sync.RWMutex.Lock 尝试获取写锁时；
//每次 sync.RWMutex.RUnlock 都会将 readerCount 其减一，当它归零时该 Goroutine 会获得写锁；
//将 readerCount 减少 rwmutexMaxReaders 个数以阻塞后续的读操作；
//调用 sync.RWMutex.Unlock 释放写锁时，会先通知所有的读操作，然后才会释放持有的互斥锁；
//读写互斥锁在互斥锁之上提供了额外的更细粒度的控制，能够在读操作远远多于写操作时提升性能。

// WaitGroup
//  sync.WaitGroup.Add 和 sync.WaitGroup.Wait：

//func (wg *WaitGroup) Add(delta int) {
//	statep, semap := wg.state()
//	state := atomic.AddUint64(statep, uint64(delta)<<32)
//	v := int32(state >> 32)
//	w := uint32(state)
//	if v < 0 {
//		panic("sync: negative WaitGroup counter")
//	}
//	if v > 0 || w == 0 {
//		return
//	}
//	*statep = 0
//	for ; w != 0; w-- {
//		runtime_Semrelease(semap, false, 0)
//	}
//}

// 当调用计数器归零，即所有任务都执行完成时，才会通过 sync.runtime_Semrelease 唤醒处于等待状态的 Goroutine。
// 当 sync.WaitGroup 的计数器归零时，陷入睡眠状态的 Goroutine 会被唤醒，上述方法也会立刻返回。

// Once
// 保证在Go程序运行期间的某段代码，只会执行一次
func main1() {
	o := &sync.Once{}
	for i := 0; i < 10; i++ {
		o.Do(func() {
			fmt.Println("only once")
		})
	}
}

// cond
// 它可以让一组的 Goroutine 都在满足特定条件时被唤醒

var status int64

func main2() {
	c := sync.NewCond(&sync.Mutex{})
	for i := 0; i < 10; i++ {
		go listen(c,i)
	}

	time.Sleep(1 * time.Second)
	go broadcast(c)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
}

func broadcast(c *sync.Cond) {
	c.L.Lock()
	// 释放信号，同步量
	atomic.StoreInt64(&status, 1)
	// c.Signal()// 唤醒队列最前面的Goroutine
	c.Broadcast()	// 唤醒队列中全部的Goroutine
	c.L.Unlock()
}

func listen(c *sync.Cond,i int) {
	c.L.Lock()
	for atomic.LoadInt64(&status) != 1 {
		c.Wait()
	}
	fmt.Printf("listen %d \n",i)
	c.L.Unlock()
}
