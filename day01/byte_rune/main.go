package main

import "fmt"

// byte和rune类型，每个英文字符都是以1个byte存储，对于其他文字，表示为rune类型(3byte)
func main() {
	s := "Hello笳淏"
	n := len(s)    // len()求的是字节数量
	fmt.Println(n) // 11

	s1 := "Hello"
	n1 := len(s1)
	fmt.Println(n1) // 5,说明一个汉字占了3个字节

	k1, k2 := 0, 0
	for i := 0; i < len(s); i++ { //这样的语句是行不通的，因为汉字占了3个字节
		fmt.Printf("%c \n", s[i]) // 出现了乱码，不能正常打印
		k1++
	}

	for _, c := range s { // 从字符串中拿出具体字符
		fmt.Printf("%c \n", c) // 能正常打印
		k2++
	}

	fmt.Printf("字符串s中有%d个汉字\n", (k1-k2)/2)

	// 字符串修改
	s2 := "白萝卜"
	//s2[0] = "红"  // 错误，字符串是不能直接修改的变量类型
	s3 := []rune(s2) // 将s2强制转化成了一个rune切片 [白, 萝, 卜]
	fmt.Println(s3)  // 打印的是汉字的编码
	s3[0] = '红'
	fmt.Println(string(s3)) // 把s3强制转换成s3

	c1 := "红" // string
	c2 := '红' // rune(int32)
	fmt.Printf("%T %T \n", c1, c2)

	c3 := "H"  // string
	c4 := 'H'  // int32
	c5 := byte('H')  // byte(uint8)
	fmt.Printf("%T %T %T\n", c3, c4, c5) // string int32 uint8

	// 类型转换
	n2 := 10  // int
	var f float64
	f = float64(n2)
	fmt.Printf("f:%T %f\n", f, f)
}
