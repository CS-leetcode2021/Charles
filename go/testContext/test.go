package main

import (
	"context"
	"fmt"
	"time"
)

/**
 *  @ClassName:test
 *  @Description:TODO
 *  @Author:jackey
 *  @Create:2021/8/18 下午2:30
 */

// Go 语言中的 context.Context 的主要作用还是在多个 Goroutine 组成的树中同步取消信号以减少对资源的消耗和占用，虽然它也有传值的功能，但是这个功能我们还是很少用到。
func main() {
	// 创建一个过期时间为1s的上下文
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	go handle(ctx, 1500*time.Millisecond)
	select {
	case <-ctx.Done():
		fmt.Println("main", ctx.Err())
	}

	time.Sleep(1*time.Second)
}

// duration持续时间,过期时间大于处理时间，有足够多的时间处理该请求
// 设置1500程序会因为上下文过期而终止
// 多个 Goroutine 同时订阅 ctx.Done() 管道中的消息，一旦接收到取消信号就立刻停止当前正在执行的工作
func handle(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Println("handle", ctx.Err())
	case <-time.After(duration):
		fmt.Println("process request with", duration)
	}
}

// 不难发现 context.emptyCtx 通过空方法实现了 context.Context 接口中的所有方法，它没有任何功能。
//