package main

import (
	"fmt"
	"runtime"
	"sync"
)

// Go运行时的调度器使用GOMAXPROCS参数来确定需要使用多少个OS线程来同时执行Go代码。默认值是机器上的CPU核心数。
// 例如在一个8核心的机器上，调度器会把Go代码同时调度到8个OS线程上（GOMAXPROCS是m:n调度中的n）。
// Go语言中可以通过runtime.GOMAXPROCS()函数设置当前程序并发时占用的CPU逻辑核心数。
// Go1.5版本之前，默认使用的是单核心执行。Go1.5版本之后，默认使用全部的CPU逻辑核心数

var wg sync.WaitGroup

func a() {
	for i:=0;i<10;i++ {
		fmt.Println("A:", i)
	}
	wg.Done()        // a函数执行完毕，计数器减一
}

func b() {
	for i:=0;i<10;i++{
		fmt.Println("B:", i)
	}
	wg.Done()        // b函数执行完毕，计数器减一
}

func main() {
	runtime.GOMAXPROCS(1)   // 只使用1个CPU核心，此时只能先执行完a或b函数，无法并发执行。若使用多个CPU核心，可能得到下面结果
	/*
	B: 0
	B: 1
	B: 2
	B: 3
	B: 4
	B: 5
	B: 6
	A: 0
	A: 1
	A: 2
	A: 3
	A: 4
	A: 5
	A: 6
	A: 7
	A: 8
	A: 9
	B: 7
	B: 8
	B: 9
	*/
	wg.Add(2)   // 因为要实现两个goroutine(函数)，初识计数器设置为2
	go a()
	go b()
	wg.Wait()
}
