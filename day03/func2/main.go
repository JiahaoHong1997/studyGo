package main

import (
	"errors"
	"fmt"
)

// 定义全局变量
var num int64 = 10

// 可以使用type关键字来定义一个函数类型
type calculation func(int, int) int

func testGlobalVar() {
	fmt.Printf("num=%d\n", num) //函数中可以访问全局变量num
}

func testNum() {
	num := 100
	fmt.Printf("num=%d\n", num)  // 函数中优先使用局部变量,所以此处num=100
}

// 语句块定义的变量
func testLocalVar(x, y int) {
	fmt.Println(x, y)
	if x > 0 {
		z := 100   //变量z只在if语句块生效
		fmt.Println(z)
	}
	// fmt.Println(z)  //此处无法使用变量z
}

func testLocalVar2() {
	for i := 0; i < 10; i++{
		fmt.Println(i)  //变量i只在当前for语句块中生效
	}
	//fmt.Println(i) //此处无法使用变量i
}

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

// 高阶函数分为函数作为参数和函数作为返回值两部分。
func calc(x, y int, op func(int, int) int) int{   // 函数可以作为参数：
	return op(x, y)
}

func do(s string) (func(int, int) int, error) {   // 函数也可以作为返回值
	switch s {
	case "+":
		return add, nil
	case "-":
		return sub, nil
	default:
		err := errors.New("无法识别的操作符")
		return nil, err
	}
}

func main(){
	// ############# 变量作用域 ##################
	testGlobalVar() // num=10
	testNum() // num=100
	testLocalVar(10, 20)
	testLocalVar2()

	var c calculation   // add和sub都能赋值给calculation类型的变量
	c = add             // 把add赋值给c
	fmt.Printf("type of c:%T\n", c)  // type of c:main.calculation
	fmt.Println(c(1, 2))  // 像调用add一样调用c

	f := add			// 将函数add赋值给变量f1
	fmt.Printf("type of f:%T\n", f) // type of f:func(int, int) int
	fmt.Println(f(10, 20))          // 像调用add一样调用f

	ret := calc(10, 20, add) // 函数可以作为参数
	fmt.Println(ret)   // 30

	option1, _ := do("+")   // option被赋值为函数add
	fmt.Printf("20+30=%d\n", option1(20, 30))  // 20+30=50
	option2, _ := do("-")
	fmt.Printf("20-30=%d\n", option2(20, 30))  // 20-30=-10
	option3, err := do("*")
	fmt.Printf("type of option1:%T\n", option1)  // type of option1:func(int, int) int
	fmt.Printf("type of option2:%T\n", option2)  // type of option1:func(int, int) int
	fmt.Printf("type of option3:%T\n", option3)  // type of option1:func(int, int) int
	fmt.Printf("type of err:%T\n", err)          // type of err:*errors.errorString
	//fmt.Printf("20*30=%d\n", option3(20, 30))    // panic: runtime error: invalid memory address or nil pointer dereference
	fmt.Printf("err:%v\n", err)
}
