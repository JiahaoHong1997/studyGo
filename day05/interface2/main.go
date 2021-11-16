package main

import "fmt"

// ########## 空接口的应用 #############
// 空接口作为函数的参数
func show(a interface{}) {
	fmt.Printf("type:%T   value:%v\n", a, a)
}

func main() {
	// ############### 空接口 ####################
	// 空接口是指没有定义任何方法的接口。因此任何类型都实现了空接口。
	// 空接口类型的变量可以存储任意类型的变量。

	// 空接口的应用
	// 1.空接口类型作为函数的参数
	// 2.空接口类型可以作为map的value
	// 3.
	var x interface{}
	show(x)      // type:<nil>   value:<nil>
	s := "Hello 笳淏"
	x = s
	fmt.Printf("type:%T  value:%v\n", x, x)   // type:string  value:Hello 笳淏
	i := 100
	x = i
	fmt.Printf("type:%T  value:%v\n", x, x)   // type:int  value:100
	b := true
	x = b
	fmt.Printf("type:%T  value:%v\n", x, x)   // type:bool  value:true

	// ########## 类型断言 #############
	// x.(T)    T为要判断的空接口的数据类型，猜的类型不对时，ok=false,ret=T的类型的零值
	ret, ok := x.(string)
	if ok {
		fmt.Println(ret)
	}else {
		fmt.Println("断言类型失败!", ret)
	}

	// 使用switch进行类型断言
	switch t := x.(type) {
	case string:
		fmt.Printf("是字符串,value:%v\n", t)
	case bool:
		fmt.Printf("是布尔,value:%v\n", t)
	case int:
		fmt.Printf("是int,value:%v\n", t)
	default:
		fmt.Printf("猜不到了,value:%v\n", t)
	}



	// 空接口作为map值
	var studentInfo = make(map[string]interface{})
	studentInfo["name"] = "沙河娜扎"
	studentInfo["age"] = 18
	studentInfo["married"] = false
	fmt.Println(studentInfo)    // map[age:18 married:false name:沙河娜扎]

}
