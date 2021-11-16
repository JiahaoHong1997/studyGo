package main

import "fmt"

// Go语言中接口（interface）是一种类型，一种抽象的类型。
// interface是一组method的集合，是duck-type programming的一种体现。接口做的事情就像是定义一个协议（规则），
// 只要一台机器有洗衣服和甩干的功能，我就称它为洗衣机。不关心属性（数据），只关心行为（方法）。

// Go语言提倡面向接口编程。
// 每个接口由数个方法组成，接口的定义格式如下：
// type 接口类型名 interface{
//    方法名1( 参数列表1 ) 返回值列表1
//    方法名2( 参数列表2 ) 返回值列表2
//    …
// }


// ############## 实现接口的条件 #################
// 一个对象只要全部实现了接口中的方法，那么就实现了这个接口。换句话说，接口就是一个需要实现的方法列表。
// 定义一个Sayer接口,只要实现了say()这个方法的类型都可以称为sayer类型
type Sayer interface {
	say()
}

// 定义dog和cat两个结构体：
type dog struct{}
type cat struct{}

// dog实现了Sayer接口
func (d dog) say() {
	fmt.Println("汪汪汪！")
}

// cat实现了Sayer接口
func (c cat) say() {
	fmt.Println("喵喵喵！")
}

func main() {
	var x Sayer // 声明一个Sayer类型的变量
	a := cat{}
	b := dog{}
	x = a       // 可以把cat实例直接赋值给x
	x.say()     // 喵喵喵！
	x = b       // 把dog实例直接赋值给x
	x.say()     // 汪汪汪！
}
