package main

import (
	"fmt"
	"time"
)

// ########################## 并发与并行的区别 ##################################
// 并发:同一时段内执行多个任务(用微信和两个人同时聊天)
// 并行:同一时刻执行多个任务(你和你朋友都在用微信和女朋友聊天)


// ########################## goroutine ######################################
// goroutine的概念类似于线程，但 goroutine是由Go的运行时（runtime）调度和管理的。Go程序会智能地将 goroutine 中的任务合理地分配给每个CPU
// 在Go语言编程中你不需要去自己写进程、线程、协程，你的技能包里只有一个技能–goroutine，当你需要让某个任务并发执行的时候，你只需要把这个任务包装成一个函数，
// 开启一个goroutine去执行这个函数就可以了，就是这么简单粗暴。一个goroutine必定对应一个函数，可以创建多个goroutine去执行相同的函数


// ########################## 启动单个goroutine ################################
// 启动goroutine的方式非常简单，只需要在调用的函数（普通函数和匿名函数）前面加上一个go关键字。
func hello() {
	fmt.Println("Hello Goroutine!")
}

func hello1() {
	fmt.Println("Hello Goroutine again!")
}

func main() {
	// 这个示例中hello函数和下面的语句是串行的，执行的结果是打印完Hello Goroutine!后打印main goroutine done!。
	hello()                                    // Hello Goroutine!
	fmt.Println("main goroutine done!")   // main goroutine done!

	fmt.Println("#################################################################################################")

	// 接下来我们在调用hello函数前面加上关键字go，也就是启动一个goroutine去执行hello这个函数。
	// 这一次的执行结果只打印了main goroutine done!，并没有打印Hello Goroutine!
	// 在程序启动时，Go程序就会为main()函数创建一个默认的goroutine。
	//
	// 当main()函数返回的时候该goroutine就结束了，所有在main()函数中启动的goroutine会一同结束，main函数所在的goroutine就像是权利的游戏中的夜王，
	// 其他的goroutine都是异鬼，夜王一死它转化的那些异鬼也就全部GG了。
	go hello()                                 // 启动另一个goroutine去执行hello函数
	fmt.Println("main goroutine done!")

	fmt.Println("#################################################################################################")

	// 要想办法让main函数等一等hello函数，最简单粗暴的方式就是time.Sleep了
	// 为什么会先打印main goroutine done!是因为我们在创建新的goroutine的时候需要花费一些时间，而此时main函数所在的goroutine是继续执行的
	go hello1()
	fmt.Println("main goroutine done!")
	time.Sleep(time.Second)
}
