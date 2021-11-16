package main

import "fmt"

// new是一个内置的函数，它的函数签名如下：
// func new(Type) *Type
// Type表示类型，new函数只接受一个参数，这个参数是一个类型
// *Type表示类型指针，new函数返回一个指向该类型内存地址的指针。

func main() {
	a := new(int)  // new函数返回指向类型的指针
	b := new(bool)
	fmt.Printf("%T\n", a) // *int
	fmt.Printf("%T\n", b) // *bool
	fmt.Println(*a)       // 0
	fmt.Println(*b)       // false

	// var a *int只是声明了一个指针变量a但是没有初始化，指针作为引用类型需要初始化后才会拥有内存空间，才可以给它赋值
	var c *int
	c = new(int)
	*c = 10
	fmt.Println(*c)    // 10


	// ################ make #####################
	// make也是用于内存分配的，区别于new，它只用于slice、map以及chan的内存创建，而且它返回的类型就是这三个类型本身，而不是他们的指针类型，
	// 因为这三种类型就是引用类型，所以就没有必要返回他们的指针了。

	// new与make的区别
	// 1.二者都是用来做内存分配的;
	// 2.make只用于slice、map以及channel的初始化，返回的还是这三个引用类型本身;
	// 3.而new用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针。
}
