package main

import "fmt"

// channel
/*
channel是一种类型，一种引用类型。声明通道类型的格式如下:
var 变量 chan 元素类型

例如：
var ch1 chan int  // 声明一个传递整型的通道
var ch2 chan bool // 声明一个传递布尔型的通道
var ch3 chan []int  // 声明一个传递int切片的通道

声明的通道后需要使用make函数初始化之后才能使用:
make(chan 元素类型, [缓冲大小])

例如:
ch4 := make(chan int)
ch5 := make(chan bool)
ch6 := make(chan []int)

*/

func main() {
	// 通道是引用类型，通道类型的空值是nil
	var ch chan int
	fmt.Println(ch) // <nil>

	// 声明的通道后需要使用make函数初始化之后才能使用
	/*
	我们使用ch := make(chan int)创建的是无缓冲的通道，无缓冲的通道只有在有人接收值的时候才能发送值。
	就像你住的小区没有快递柜和代收点，快递员给你打电话必须要把这个物品送到你的手中，简单来说就是无缓冲的通道必须有接收才能发送。
	下面的代码会阻塞在ch <- 10这一行代码形成死锁
	 */
	ch = make(chan int)

	// ################# channel操作 #####################
	// 通道有发送（send）、接收(receive）和关闭（close）三种操作。发送和接收都使用<-符号

	// 1.发送:将一个值发送到通道中
	ch <- 10 // 把10发送到ch中

	// 2.接收:从一个通道中接收值
	x := <-ch  // 从ch中接收值并赋值给变量x
	<-ch    // 从ch中接收值，忽略结果
	fmt.Println(x)

	// 3.关闭:通过调用内置的close函数来关闭通道
	close(ch)
	// 关于关闭通道需要注意的事情是，只有在通知接收方goroutine所有的数据都发送完毕的时候才需要关闭通道。
	// 通道是可以被垃圾回收机制回收的，它和关闭文件是不一样的，在结束操作之后关闭文件是必须要做的，但关闭通道不是必须的。
	/* 关闭后的通道有以下特点：
	1.对一个关闭的通道再发送值就会导致panic。
	2.对一个关闭的通道进行接收会一直获取值直到通道为空。
	3.对一个关闭的并且没有值的通道执行接收操作会得到对应类型的零值。
	4.关闭一个已经关闭的通道会导致panic。
	*/

	//

}
