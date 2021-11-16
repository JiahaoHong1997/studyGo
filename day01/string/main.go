package main

import (
	"fmt"
	"strings"
)

func main(){
	// go语言中的字符串只能用""来包裹
	path := "\"~/go/src/github.com/JiahaoHong1997/day01/string\""  // \是转义字符
	fmt.Println(path)

	// 多行的字符串
	s := `
世情薄
人情恶
雨送黄昏花易落
	`
	fmt.Println(s)  // ``包裹的是什么格式，就怎么打印

	// 用``直接输出路径
	s2 := `~/go/src/github.com/JiahaoHong1997/day01/string`
	fmt.Println(s2)

	// 字符串相关操作
	fmt.Println(len(s2))  // 打印字符串长度

	// 字符串拼接
	name := "理想"
	world := "dsb"
	ss := name + world
	fmt.Println(name + world)
	fmt.Printf("%s \n", ss)
	ss1 := fmt.Sprintf("%s%s", name, world) // fmt.Sprintf并非打印操作，而是拼接操作
	fmt.Println(ss1)

	// 字符串分割
	ret := strings.Split(s2, "/")  // 以/为标志分割s2字符串
	fmt.Printf("ret:%T\n", ret)
	fmt.Println(ret)

	// 判断包含
	fmt.Println(strings.Contains(ss, "理性")) //false
	fmt.Println(strings.Contains(ss, "理想")) //true

	// 前缀和后缀判断
	fmt.Println(strings.HasPrefix(ss, "理想")) // 判断是否以理想开头 true
	fmt.Println(strings.HasSuffix(ss, "sb")) // 判断是否以sb结尾 true

	// 子串
	s3 := "abcdeb"
	fmt.Println(strings.Index(s3, "c")) // 判断c在字符串中出现的位置
	fmt.Println(strings.LastIndex(s3, "b")) // 判断b在字符串中最后出现的位置

	// 拼接
	fmt.Println(strings.Join(ret, "+")) // 用+把字符串拼接起来（ret是一个列表）

}
