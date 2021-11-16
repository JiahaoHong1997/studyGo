package main

import (
	"encoding/json"
	"fmt"
)

// ################## 结构体的“继承 ######################
// Animal
type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s会动！\n", a.name)
}

// Dog
type Dog struct {
	Feet    int8
	*Animal // 通过嵌套匿名结构体实现继承
}

func (d *Dog) wang() {
	fmt.Printf("%s会汪汪汪~\n", d.name)
}

// ################# 结构体字段的可见性 #####################
// 结构体中字段大写开头表示可公开访问，小写表示私有（仅在定义当前结构体的包中可访问）。

// ################# 结构体与JSON序列化 #####################
// Student
type Student struct {
	ID     int
	Gender string
	Name   string
}

// Class
type Class struct {
	Title    string
	Students []*Student
}

// ################### 结构体标签（Tag） ######################
// Tag在结构体字段的后方定义，由一对反引号包裹起来，当转化成json包时替代该字段
//Student1 学生
type Student1 struct {
	ID     int    `json:"id"` //通过指定tag实现json序列化该字段时的key
	Gender string //json序列化是默认使用字段名作为key
	name   string //私有不能被json包访问
}

// ######### slice和map这两种数据类型都包含了指向底层数据的指针，因此我们在需要复制它们时要特别注意 ###############
type Person struct {
	name   string
	age    int8
	dreams []string
}

func (p *Person) SetDreams(dreams []string) {   // 直接传入slice进行赋值
	p.dreams = dreams
}

func (p *Person) SetDreams1(dreams []string) {  // 在方法中使用传入的slice的拷贝进行结构体赋值
	p.dreams = make([]string, len(dreams))
	copy(p.dreams, dreams)
}

func main() {
	d1 := Dog{
		4,
		&Animal{ //注意嵌套的是结构体指针
			"乐乐",
		},
	}
	d1.wang() // 乐乐会汪汪汪~
	d1.move() // 乐乐会动！

	c := &Class{
		"101",
		make([]*Student, 0, 200),
	}
	for i := 0; i < 10; i++ {
		stu := &Student{
			i,
			"男",
			fmt.Sprintf("stu%02d", i),
		}
		c.Students = append(c.Students, stu)
	}
	//JSON序列化：结构体-->JSON格式的字符串
	data, err := json.Marshal(c)
	if err != nil {
		fmt.Println("json marshal failed")
		return
	}
	fmt.Printf("json:%s\n", data)
	//JSON反序列化：JSON格式的字符串-->结构体
	str := `{"Title":"101","Students":[{"ID":0,"Gender":"男","Name":"stu00"},{"ID":1,"Gender":"男","Name":"stu01"},{"ID":2,"Gender":"男","Name":"stu02"},{"ID":3,"Gender":"男","Name":"stu03"},{"ID":4,"Gender":"男","Name":"stu04"},{"ID":5,"Gender":"男","Name":"stu05"},{"ID":6,"Gender":"男","Name":"stu06"},{"ID":7,"Gender":"男","Name":"stu07"},{"ID":8,"Gender":"男","Name":"stu08"},{"ID":9,"Gender":"男","Name":"stu09"}]}`
	c1 := &Class{}
	err = json.Unmarshal([]byte(str), c1)
	if err != nil {
		fmt.Println("json unmarshal failed!")
		return
	}
	fmt.Printf("%#v\n", c1)

	s1 := Student1{
		1,
		"男",
		"洪笳淏",
	}
	data1, err1 := json.Marshal(s1)
	if err1 != nil {
		fmt.Println("json marshal failed!")
		return
	}
	fmt.Printf("json str:%s\n", data1) //json str:{"id":1,"Gender":"男"}


	p1 := Person{name: "小王子", age: 18}
	data2 := []string{"吃饭", "睡觉", "打豆豆"}
	p1.SetDreams(data2)

	data2[1] = "不睡觉"  // 此时改变了p1.dreams的值
	fmt.Println(p1.dreams) // [吃饭 不睡觉 打豆豆]

	data3 := []string{"吃饭", "睡觉", "打豆豆"}
	p1.SetDreams1(data3)
	data3[1] = "不睡觉"  // 没改变p1.dreams的值
	fmt.Println(p1.dreams)  // [吃饭 睡觉 打豆豆]

}
