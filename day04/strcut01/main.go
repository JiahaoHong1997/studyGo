package main

import (
	"fmt"
	"unsafe"
)

// 类型定义
type NewInt int

// 类型别名
type MyInt = int

// 结构体
type person struct {
	name string
	city string
	age int8
}

// 构造函数
func newPerson(name, city string, age int8) *person {
	return &person{
		name:name,
		city:city,
		age:age,
	}
}

func main() {
	// ########### 自定义类型 ################
	// 将MyInt定义为int类型
	//type MyInt int    // 通过type关键字的定义，MyInt就是一种新的类型，它具有int的特性

	// ########### 类型别名 #################
	// 之前见过的rune和byte就是类型别名，他们的定义如下：
	// type byte = uint8
	// type rune = int32

	// ############ 类型定义和类型别名的区别 #################
	var a NewInt
	var b MyInt

	fmt.Printf("type of a:%T\n", a) //type of a:main.NewInt
	fmt.Printf("type of b:%T\n", b) //type of b:int

	// ############ 结构体 ################
	var p1 person
	p1.name = "洪笳淏"
	p1.city = "贵阳"
	p1.age = 23
	fmt.Printf("p1=%v\n", p1)   // p1={洪笳淏 贵阳 23}
	fmt.Printf("p1=%#v\n", p1)    // p1=main.person{name:"洪笳淏", city:"贵阳", age:23}

	// ########### 匿名结构体 #############
	// 在定义一些临时数据结构等场景下还可以使用匿名结构体。
	var user struct {Name string; Age int}
	user.Name = "洪笳淏"
	user.Age = 23
	fmt.Printf("%#v\n", user)

	// ############ 创建指针类型结构体 ################
	var p2 = new(person)   // new创建指向person类型的指针，没有初始化的结构体，其成员变量都是对应其类型的零值。
	fmt.Printf("%T\n", p2)    // *main.person
	fmt.Printf("p2=%#v\n", p2)    // p2=&main.person{name:"", city:"", age:0}

	// ############ 取结构体的地址实例化 ##################
	// 使用&对结构体进行取地址操作相当于对该结构体类型进行了一次new实例化操作
	p3 := &person{}
	fmt.Printf("%T\n", p3)     //*main.person
	fmt.Printf("p3=%#v\n", p3)    //p3=&main.person{name:"", city:"", age:0}
	p3.name = "洪笳淏"   // p3.name = "洪笳淏"其实在底层是(*p3).name = "洪笳淏"
	p3.age = 23
	p3.city = "贵阳"
	fmt.Printf("p3=%#v\n", p3)   //p3=&main.person{name:"七米", city:"成都", age:30}

	// ############### 结构体初始化 #################
	type man struct {
		name string
		age int8
		city string
	}

	var p4 man    // 没有初始化的结构体，其成员变量都是对应其类型的零值。
	fmt.Printf("p4=%#v\n", p4)   // p4=main.man{name:"", age:0, city:""}

	// 使用键值对初始化
	p5 := man{
		name:"洪笳淏",
		age:23,
		city:"贵阳",
	}
	fmt.Printf("p5=%#v\n", p5) // p5=main.man{name:"洪笳淏", age:23, city:"贵阳"}

	// 也可以对结构体指针进行键值对初始化，例如：
	p6 := &man{
		name: "小王子",
		city: "北京",
		age:  18,
	}
	fmt.Printf("p6=%#v\n", p6) // p6=&main.man{name:"小王子", age:18, city:"北京"}

	// 当某些字段没有初始值的时候，该字段可以不写。此时，没有指定初始值的字段的值就是该字段类型的零值。

	p7 := &man{
		city: "北京",
	}
	fmt.Printf("p7=%#v\n", p7) // p7=&main.man{name:"", age:0, city:"北京"}

	// 初始化结构体的时候可以简写，也就是初始化的时候不写键，直接写值：
	p8 := &man{
		"洪笳淏",
		23,
		"贵阳",
	}
	fmt.Printf("p8=%#v\n", p8) // p8=&main.man{name:"洪笳淏", age:23, city:"贵阳"}
	// 使用这种格式初始化时，需要注意：
	// 1.必须初始化结构体的所有字段。
	// 2.初始值的填充顺序必须与字段在结构体中的声明顺序一致。
	// 3.该方式不能和键值初始化方式混用。


	// ################# 结构体内存布局 #######################
	// 结构体占用一块连续的内存
	type test struct {
		a int8
		b int8
		c int8
		d int8
	}
	n := test{
		1, 2, 3, 4,
	}
	fmt.Printf("n.a %p\n", &n.a)  // n.a 0xc0000b2100
	fmt.Printf("n.b %p\n", &n.b)  // n.b 0xc0000b2101
	fmt.Printf("n.c %p\n", &n.c)  // n.c 0xc0000b2102
	fmt.Printf("n.d %p\n", &n.d)  // n.d 0xc0000b2103

	// ############## 空结构体 ################
	var v struct{}
	fmt.Println(unsafe.Sizeof(v))  // 0    空结构体是不占用空间的

	// 面试题
	type student struct {
		name string
		age  int
	}
	m := make(map[string]*student)
	stus := []student{
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八", age: 9000},
		{name: "小变态", age: 1},
	}

	for _, stu := range stus {
		fmt.Println(stu)
		m[stu.name] = &stu
		e := &stu // 总是指向末尾元素，注：只有当切片元素是引用类型时，才能用指针指向其末尾元素
		fmt.Printf("e:%p  stus:%p \n", e, &stus)
	}
	fmt.Println(m)
	for k, v := range m {
		fmt.Println(*v)
		fmt.Println(k, "=>", v.name)
	}

	// ############## 构造函数 #################
	// Go语言的结构体没有构造函数，我们可以自己实现。 例如，下方的代码就实现了一个person的构造函数。
	// 因为struct是值类型，如果结构体比较复杂的话，值拷贝性能开销会比较大，所以该构造函数返回的是结构体指针类型。
	p9 := newPerson("洪笳淏", "贵阳", 23)
	fmt.Printf("%#v\n", p9)   // &main.person{name:"洪笳淏", city:"贵阳", age:23}

}
