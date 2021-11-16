package main

import "fmt"

// 详见 https://sanyuesha.com/2017/07/23/go-defer/
func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}
func main() {
	fmt.Println(f1())  // 5
	fmt.Println(f2())  // 6
	fmt.Println(f3())  // 5
	fmt.Println(f4())  // 5

	main2()
}




// 规则1：当defer被声明时，其参数就会被实时解析
// 规则2：defer执行顺序为先进后出
// 规则3：defer可以读取有名返回值
func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main2() {
	x := 1
	y := 2
	defer calc("AA", x, calc("A", x, y))
	x = 10
	defer calc("BB", x, calc("B", x, y))
	y = 20
}
