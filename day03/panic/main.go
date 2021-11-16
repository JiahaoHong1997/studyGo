package main

import "fmt"

// Go语言中目前（Go1.12）是没有异常机制，但是使用panic/recover模式来处理错误。 panic可以在任何地方引发，但recover只有在defer调用的函数中有效。

func main() {
	funcA()
	funcB()
	funcC()
}


func funcA() {
	fmt.Println("func A")
}

func funcB() {
	defer func() {
		// recover()必须搭配defer使用。
		// defer一定要在可能引发panic的语句之前定义。
		err := recover()
		//如果程序出出现了panic错误,可以通过recover恢复过来
		if err != nil {
			fmt.Println("recover in B")
		}
	}()
	panic("panic in B")
}

func funcC() {
	fmt.Println("func C")
}