package main

import (
	"fmt"
	m "github.com/JiahaoHong1997/studygo/day05/package" // m是包的别名,后面是包的路径
)

func main() {
	fmt.Println(m.Add(100, 200))  // 300
	fmt.Println(m.Mode)                 // 1
}