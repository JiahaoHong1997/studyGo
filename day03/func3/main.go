package main

import (
	"fmt"
	"strings"
)

// 匿名函数和闭包
func main(){
	// 将匿名函数保存到变量
	// 匿名函数因为没有函数名，所以没办法像普通函数那样调用，所以匿名函数需要保存到某个变量或者作为立即执行函数:
	add := func(x, y int) {
		fmt.Println(x + y)
	}
	add(10, 20) // 通过变量调用匿名函数

	//自执行函数：匿名函数定义完加()直接执行
	func(x, y int) {
		fmt.Println(x + y)
	}(10, 20)

	// 闭包指的是一个函数和与其相关的引用环境组合而成的实体，简单来说，闭包=函数+引用环境
	// 例1:
	var f = adder()
	fmt.Println(f(10))  // 10
	fmt.Println(f(20))  // 30
	fmt.Println(f(30))  // 60

	f1 := adder()
	fmt.Println(f1(40))  // 40
	fmt.Println(f1(50))  // 90

	// 例2:
	f2 := adder1(10)   // x被初始化为10
	fmt.Println(f2(10))   // 20
	fmt.Println(f2(20))   // 40
	fmt.Println(f2(30))   // 70

	f3 := adder1(20)   // x被初始化为20
	fmt.Println(f3(40))   // 60
	fmt.Println(f3(50))   // 110

	// 例3:
	jpgFunc := makeSuffixFunc(".jpg")
	txtFunc := makeSuffixFunc(".txt")
	fmt.Println(jpgFunc("test"))   // test.jpg
	fmt.Println(txtFunc("test"))   // test.txt

	// 例4:
	f4, f5 := calc(10)   // f4是加法操作，f5是减法操作，他们共同继承环境变量base=10
	fmt.Println(f4(1), f5(2)) // 11, 9
	fmt.Println(f4(3), f5(4)) // 12, 8
	fmt.Println(f4(5), f5(6)) // 13, 7
}


func adder() func(int) int {
	var x int   // x默认初始化为0
	return func(y int) int {
		x += y
		return x
	}
}

func adder1(x int) func(int) int {  // x作为adder1的环境变量，对于赋值给同一个变量的函数，其值一直继承
	return func(y int) int {
		x += y
		return x
	}
}

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}