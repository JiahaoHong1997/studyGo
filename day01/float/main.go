package main

import "fmt"

func main(){
	 f1 := 1.23456
	 fmt.Printf("%T \n", f1) // float64,默认go语言中的float类型是float64
	 f2 := float32(1.23456)
	 fmt.Printf("%T \n", f2) // float32
	 // f1 = f2  //float32位的值不能直接赋值给float64类型的变量
}
