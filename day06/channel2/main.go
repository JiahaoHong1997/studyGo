package main

import "fmt"

func recv(c chan int) {
	ret := <- c
	fmt.Println("接收成功", ret)
}


func counter(out chan<- int){
	for i := 0; i< 100 ;i++{
		out <- i
	}
	close(out)
}

// chan<- int是一个只写单向通道（只能对其写入int类型值），可以对其执行发送操作但是不能执行接收操作；
// <-chan int是一个只读单向通道（只能从其读取int类型值），可以对其执行接收操作但是不能执行发送操作。
// 在函数传参及任何赋值操作中可以将双向通道转换为单向通道，但反过来是不可以的
func squarer(out chan<- int, in <-chan int){
	for i := range in {
		out <- i * i
	}
	close(out)
}

func printer(in <-chan int) {
	for i := range in {
		fmt.Println(i)
	}
}

func main() {
	// 无缓冲通道上的发送操作会阻塞，直到另一个goroutine在该通道上执行接收操作，这时值才能发送成功，两个goroutine将继续执行。
	// 相反，如果接收操作先执行，接收方的goroutine将阻塞，直到另一个goroutine在该通道上发送一个值。
	// 使用无缓冲通道进行通信将导致发送和接收的goroutine同步化。因此，无缓冲通道也被称为同步通道。

	ch := make(chan int)
	go recv(ch)   // 启用goroutine从通道接收值
	ch <- 10
	fmt.Println("发送成功")


	// #################### 有缓冲的通道 ########################
	/*
	解决上面问题的方法还有一种就是使用有缓冲区的通道。我们可以在使用make函数初始化通道的时候为其指定通道的容量。
	只要通道的容量大于零，那么该通道就是有缓冲的通道，通道的容量表示通道中能存放元素的数量。
	就像你小区的快递柜只有那么个多格子，格子满了就装不下了，就阻塞了，等到别人取走一个快递员就能往里面放一个。
	我们可以使用内置的len函数获取通道内元素的数量，使用cap函数获取通道的容量，虽然我们很少会这么做。
	*/

	ch1 := make(chan int, 1)   // 创建一个容量为1的有缓冲区通道
	ch1 <- 10
	fmt.Println("发送成功")

	// for range从通道循环取值
	/*
	当向通道中发送完数据时，我们可以通过close函数来关闭通道。
	当通道被关闭时，再往该通道发送值会引发panic，从该通道取值的操作会先取完通道中的值，再然后取到的值一直都是对应类型的零值。
	 */
	ch2 := make(chan int)
	ch3 := make(chan int)
	// 开启goroutine将0~100的数发送到ch2中
	go func() {
		for i := 0; i<100;i++{
			ch2 <- i
		}
		close(ch2)
	}()
	// 开启goroutine从ch12接收值，并将该值的平方发送到ch3中
	go func() {
		for {
			i, ok := <-ch2  // 通道关闭后再取值ok=false
			if !ok {
				break
			}
			ch3 <- i * i
		}
		close(ch3)
	}()
	// 在主goroutine中从ch3中接收值打印
	for i := range ch3 {   // 通道关闭后会退出for range循环
		fmt.Println(i)
	}
	// 从上面的例子中我们看到有两种方式在接收值的时候判断该通道是否被关闭，不过我们通常使用的是for range的方式。
	// 使用for range遍历通道，当通道被关闭的时候就会退出for range


	// ############## 单向通道 ################
	// 有的时候我们会将通道作为参数在多个任务函数间传递，很多时候我们在不同的任务函数中使用通道都会对其进行限制，比如限制通道在函数中只能发送或只能接收
	ch4 := make(chan int)
	ch5 := make(chan int)
	go counter(ch4)
	go squarer(ch5, ch4)
	printer(ch5)

}
