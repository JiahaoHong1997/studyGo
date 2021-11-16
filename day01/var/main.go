package main

import "fmt"

//批量声明变量
var (
	name string // ""
	age  int    // 0
	isOk bool   // false
)

func main() {
	name = "理想"
	age = 18
	isOk = true
	// Go语言中，变量声明后必须要使用，不使用就不能通过编译

	fmt.Printf("name:%s\n", name)
	fmt.Println("age:", age)
	fmt.Print(isOk)
	// fmt.Println()  有啥打啥，打印的输出末尾自动接换行符
	// fmt.Printf()   格式化打印(使用%加数据类型作为占位符)，不接换行符
	// fmt.Print()      有啥打啥，不接换行符
}
