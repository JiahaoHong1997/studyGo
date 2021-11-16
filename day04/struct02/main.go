package main

import "fmt"

// 方法和接收者
// 方法的定义格式如下：
// func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) {
// 函数体
// }


// Person结构体
type Person struct{
	name string
	age int8
}

// newPerson构造函数
func newPerson(name string, age int8) *Person{
	return &Person{
		name:name,
		age:age,
	}
}

// Dream Person做梦的方法
func (p Person) Dream() {    // 方法与函数的区别是，函数不属于任何类型，方法属于特定的类型。
	fmt.Printf("%s的梦想是学好Go语言！\n", p.name)
}

// ############# 指针类型的接收者 ###################
// SetAge 设置p的年龄
// 使用指针接收者
func (p *Person) SetAge(newAge int8) {
	p.age = newAge
}

// 什么时候应该使用指针类型接收者
// 1.需要修改接收者中的值
// 2.接收者是拷贝代价比较大的大对象
// 3.保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。

// ############# 值类型的接收者 ###################
// SetAge2 设置p的年龄
// 使用值接收者
func (p Person) SetAge2(newAge int8) {   // 在值类型接收者的方法中可以获取接收者的成员值，但修改操作只是针对副本，无法修改接收者变量本身。
	p.age = newAge
	fmt.Printf("副本中的age值为%d\n", p.age)
}


// ############## 任意类型添加方法 #################
// 非本地类型不能定义方法，也就是说我们不能给别的包的类型定义方法。
//MyInt 将int定义为自定义MyInt类型
type MyInt int

//SayHello 为MyInt添加一个SayHello的方法
func (m MyInt) SayHello() {
	fmt.Println("Hello, 我是一个int。")
}

// ############## 结构体的匿名字段 ##################
// 结构体允许其成员字段在声明时没有字段名而只有类型，这种没有名字的字段就称为匿名字段。
// 这里匿名字段的说法并不代表没有字段名，而是默认会采用类型名作为字段名，结构体要求字段名称必须唯一，因此一个结构体中同种类型的匿名字段只能有一个。
type man struct {
	string
	int
}

// ############### 嵌套结构体 ######################
// 一个结构体中可以嵌套包含另一个结构体或结构体指针
//Address 地址结构体
type Address struct {
	Province string
	City     string
}

//User 用户结构体
type User struct {
	Name    string
	Gender  string
	Address Address
}

// ################# 嵌套匿名字段 ###################
// 上面user结构体中嵌套的Address结构体也可以采用匿名字段的方式
//Address 地址结构体
type Address1 struct {
	Province string
	City     string
}

//User 用户结构体
type User1 struct {
	Name    string
	Gender  string
	Address1 //匿名字段
}

// ################# 嵌套结构体的字段名冲 #####################
//Address 地址结构体
type Address2 struct {
	Province   string
	City       string
	CreateTime string
}

//Email 邮箱结构体
type Email struct {
	Account    string
	CreateTime string
}

//User 用户结构体
type User2 struct {
	Name   string
	Gender string
	Address2
	Email
}


func main() {
	p1 := newPerson("洪笳淏", 23)
	p1.Dream()
	fmt.Println(p1.age) // 23

	p1.SetAge(30)
	fmt.Println(p1.age)   // 30

	p1.SetAge2(40)
	fmt.Println(p1.age)   // 30

	var m1 MyInt
	m1.SayHello() //Hello, 我是一个int。
	m1 = 100
	fmt.Printf("%#v  %T\n", m1, m1) //100  main.MyInt

	p2 := man{
		"洪笳淏",
		23,
	}
	fmt.Printf("%#v\n", p2)   // main.man{string:"洪笳淏", int:23}
	fmt.Println(p2.string, p2.int)   // 洪笳淏  23

	user1 := User{
		"洪笳淏",
		"男",
		Address{
			"贵州",
			"贵阳",
		},
	}
	fmt.Printf("%#v\n", user1)  // main.User{Name:"洪笳淏", Gender:"男", Address:main.Address{Province:"贵州", City:"贵阳"}}

	var user2 User1
	user2.Name = "小王子"
	user2.Gender = "男"
	user2.Address1.Province = "山东"  // 匿名字段默认使用类型名作为字段名
	user2.Address1.City = "威海"      // 匿名字段可以省略
	fmt.Printf("user2=%#v\n", user2) //user2=main.User{Name:"小王子", Gender:"男", Address:main.Address{Province:"山东", City:"威海"}}

	var user3 User2
	user3.Name = "吴星怡"
	user3.Gender = "女"
	user3.Address2.CreateTime = "2000" //指定Address结构体中的CreateTime
	user3.Email.CreateTime = "2000"   //指定Email结构体中的CreateTime
	fmt.Printf("user3=%#v\n", user3)
	// user3=main.User2{Name:"吴星怡", Gender:"女", Address2:main.Address2{Province:"", City:"", CreateTime:"2000"}, Email:main.Email{Account:"", CreateTime:"2000"}}
}
