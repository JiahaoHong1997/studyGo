package main

import "fmt"

func main(){
	// 十进制
	var a int = 10
	fmt.Printf("%d \n", a) // 十进制 10
	fmt.Printf("%b \n", a) // 二进制 1010

	// 八进制, 以0开头
	var b int = 077
	fmt.Printf("%o \n", b) // 八进制 77
	fmt.Printf("%d \n", b) // 十进制 63

	// 十六进制，以0x开头
	var c int = 0xff
	fmt.Printf("%x \n", c) // 十六进制 ff
	fmt.Printf("%d \n", c) // 十进制 255
	fmt.Printf("%T \n", c) // %T 用于查看变量类型

	// 声明int8类型变量
	d := int8(9)
	fmt.Printf("%T \n", d) // int8

}
