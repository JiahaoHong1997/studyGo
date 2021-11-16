package main

import "fmt"

// const pi = 3.1415926    //单独声明

const (
	pi       = 3.1415926
	r        = 3
	statusOK = 200
	notfound = 404
) // 批量声明

const (
	n1 = 200
	n2
	n3
) // 在批量声明常量的时候，若某一行没有写值，则默认和上一行值相同

// iota
// iota是go语言的常量计数器，只能在常量的表达式中式使用。iota在const关键字出现时将重置为0。const中每新增一行常量声明将使iota计数一次。
//（iota可理解为const语句块中的行索引）使用iota能简化定义，在定义枚举时很有用
const (
	a1 = iota // 0
	a2        // 1 (默认等于iota，由于iota增加，a2值也增加)
	a3        // 2
	a4        // 3
)

const ( // 新的const关键字出现，iota重置为0
	b1 = iota   // 0
	b2          // 1
	_           // 2  相当于把这个变量丢弃了，不能打印 _
	b3          // 3
)

// 插队
const (
	c1 = iota // 0
	c2 = 100
	c3        // 100
	c4 = iota  // 3
)

// 多个变量声明在一行
const(
	d1, d2 = iota + 1, iota + 2  // d1 = 1 , d2 = 2
	d3, d4 = iota + 1, iota + 2  // d3 = 2 , d4 = 3
)

// 定义数量级
const(
	_ = iota
	KB = 1 << (10 * iota)  // 2^10
	MB = 1 << (10 * iota)  // 2^20
	GB = 1 << (10 * iota)  // 2^30
	TB = 1 << (10 * iota)  // 2^40
	PB = 1 << (10 * iota)  // 2^50
)

func main() {
	//pi = 123    错误定义，常量定义之后就不可改变
	fmt.Println("n1:", n1)
	fmt.Println("n2:", n2)
	fmt.Println("n3:", n3)

	fmt.Println("a1:", a1)
	fmt.Println("a2:", a2)
	fmt.Println("a3:", a3)
	fmt.Println("a4:", a4)

	fmt.Println("b1:", b1)
	fmt.Println("b2:", b2)
	fmt.Println("b3:", b3)

	fmt.Println("c1:", c1)
	fmt.Println("c2:", c2)
	fmt.Println("c3:", c3)
	fmt.Println("c4:", c4)

	fmt.Println("d1:", d1)
	fmt.Println("d2:", d2)
	fmt.Println("d3:", d3)
	fmt.Println("d4:", d4)

	fmt.Println("KB:", KB)
	fmt.Println("MB:", MB)
	fmt.Println("GB:", GB)
	fmt.Println("TB:", TB)
	fmt.Println("PB:", PB)
}
