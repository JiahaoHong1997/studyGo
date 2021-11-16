package main

import (
	"fmt"
	"sync"
)

// ##################### 启动多个goroutine ###########################
// 在Go语言中实现并发就是这样简单，我们还可以启动多个goroutine
// 这里使用了sync.WaitGroup来实现goroutine的同步
var wg sync.WaitGroup

func hello(i int) {
	defer wg.Done()   // goroutine结束就登记-1
	fmt.Println("Hello Goroutine!", i)
}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)    // 启动一个goroutine就登记+1
		go hello(i)
	}
	wg.Wait()   // 等待所有登记的goroutine都结束
	// 多次执行上面的代码，会发现每次打印的数字的顺序都不一致。这是因为10个goroutine是并发执行的，而goroutine的调度是随机的。
}
