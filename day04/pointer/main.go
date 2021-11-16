package main

import "fmt"

// Go语言中的指针不能进行偏移和运算，因此Go语言中的指针操作非常简单，我们只需要记住两个符号：&（取地址）和*（根据地址取值）

func main() {
	// 取变量指针的语法如下：
	a := 10
	b := &a    // b存储变量a的地址
	fmt.Printf("a:%d ptr:%p\n", a, &a)   // a:10 ptr:0xc000132008
	fmt.Printf("b:%p type:%T\n", b, b)   // b:0xc000132008 type:*int
	fmt.Println(&b)                             // 0xc00012c018

	c := *b   // 指针取值（根据指针去内存取值）
	fmt.Printf("type of c:%T\n", c)      // int
	fmt.Printf("value of c:%v\n", c)     // 10

	// 取地址操作符&和取值操作符*是一对互补操作符，&取出地址，*根据地址取出地址指向的值。
	// 变量、指针地址、指针变量、取地址、取值的相互关系和特性如下：
	// 1.对变量进行取地址（&）操作，可以获得这个变量的指针变量。
	// 2.指针变量的值是指针地址。
	// 3.对指针变量进行取值（*）操作，可以获得指针变量指向的原变量的值。

	// 指针传值示例：
	d := 10
	modify1(d)
	fmt.Println(d)   // 10
	modify2(&d)
	fmt.Println(d)   // 100
}


func modify1(x int) {
	x = 100
}

func modify2(x *int) {
	*x = 100
}