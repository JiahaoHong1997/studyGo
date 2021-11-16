package main

import "fmt"

func sayHello() {
	fmt.Println("Hello 洪笳淏！")
}

// 定义一个接收string类型的name参数
func sayHello2(name string) {
	fmt.Println("Hello" + name)
}

// 定义接收多个参数的函数
func intSum(x int, y int) int{
	ret := x + y
	return ret
}
// 或者写为
func intSum2(x int, y int) (ret int){
	ret = x + y
	return
}

// 函数接收可变参数,在函数变量名后面加...表示可变参数
// 可变参数在函数体中是切片类型
func intSum3(a ...int) {
	fmt.Println(a)  // 【10 20】
	fmt.Printf("%T\n", a)   // []int
}

func intSum4(a ...int) int{
	ret := 0
	for _, v := range a {
		ret = ret + v
	}
	return ret
}

// 固定参数和可变参数同时出现时，可变参数要放在最后
func intSum5(a int, b ...int) int {
	ret := a
	for _,v := range b {
		ret = ret + v
	}
	return ret
}

// 定义具有多个返回值的函数
func calc(a int, b int) (sum int, sub int){
	sum = a + b
	sub = a - b
	return
}

func main() {
	// 函数调用
	sayHello()
	sayHello2("洪笳淏")
	r := intSum(10,20)
	fmt.Println(r)
	r2 := intSum2(10,20)
	fmt.Println(r2)
	intSum3(10,20)
	r3 := intSum4(10,20)
	fmt.Println(r3)
	r4 := intSum5(10, 1,2,3,4,5,6,7,8,9)  // a=10, b=[1 2 3 4 5 6 7 8 9]
	fmt.Println(r4)
	r5 := intSum5(10)  // a=10, b=[]
	fmt.Println(r5)
	x, y := calc(100,200)
	fmt.Println(x, y)
}
