package main

import "fmt"

func main() {
	// 基本格式
	for i := 0; i < 10; i++ {
		fmt.Print(i)
	}

	var i = 5
	for i < 10 {
		fmt.Println(i)
		i++
	}

	// 无限循环
	//for{
	//	fmt.Println("123")
	//}

	// for range,返回索引和值
	s := "Hello洪笳淏"
	for i, v := range s {
		fmt.Printf("%d %c\n", i, v)
	}

	// break
	for i := 0; i < 10; i++ {
		if i == 5 {
			break // 跳出for循环
		}
		fmt.Println(i)
	}
	fmt.Println("over")

	// continue
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue // 跳出本轮for循环，直接开始下一轮循环
		}
		fmt.Println(i)
	}
	fmt.Println("over")
}
