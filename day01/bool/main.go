package main

import "fmt"

//bool类型的值不能和别的类型的值互相转化
func main(){
	b1 := true
	var b2 bool // 默认是false
	fmt.Printf("%T \n", b1)
	fmt.Printf("%T value:%v\n", b2, b2)  // %v表示打印变量的值
}
