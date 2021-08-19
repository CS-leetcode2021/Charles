package testTimers

/**
 *  @ClassName:text
 *  @Description:TODO
 *  @Author:jackey
 *  @Create:2021/8/18 下午9:33
 */

// 初期：全局共用一个互斥锁，对计时器的影响非常大
// 中期：分片四叉堆，将全局的四叉堆分割成了 64 个更小的四叉堆
//  如果当前机器上的处理器 P 的个数超过了 64，多个处理器上的计时器就可能存储在同一个桶中。每一个计时器桶都由一个运行 runtime.timerproc:76f4fd8 函数的 Goroutine 处理。

// 现在：在最新版本的实现中，计时器桶已经被移除，所有的计时器都以最小四叉堆的形式存储在处理器 runtime.p 中。
