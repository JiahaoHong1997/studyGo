package main

import "fmt"

// if条件判断
func main() {
	age := 19
	if age > 18 {
		fmt.Println("澳门首家线上赌场上线了！")
	} else{
		fmt.Println("改写作业了！")
	}

	if age > 35{
		fmt.Println("人到中年")
	}else if age>18{
		fmt.Print("青年")
	}else{
		fmt.Println("好好学习！")
	}

	// age1变量此时只在下面这个if判断语句中作用，出了if语句则找不到age1（局部变量）
	if age1 := 19; age1 > 18{
		fmt.Println("澳门首家线上赌场上线了！")
	}else{
		fmt.Println("改写作业了！")
	}
}
