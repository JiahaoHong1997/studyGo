package main

import "fmt"

// 使用值接收者实现接口和使用指针接收者实现接口的区别

type mover interface {
	move()
}

type mover1 interface {
	move1()
}

type sayer interface {
	say()
}

// 接口的嵌套：定义一个实现move和say方法的接口
type animal interface {
	move1()
	say()
}

type person struct {
	name string
	age int8
}

// 使用值接收者实现接口:类型的值和类型的指针都能够保存到接口变量中
func (p person) move(){
	fmt.Printf("%s在跑\n", p.name)
}

// 使用指针接收者实现接口:只有类型的指针能够保存在接口变量中
func (p *person) move1(){
	fmt.Printf("%s在跑\n", p.name)
}

func (p *person) say(){
	fmt.Printf("%s在叫\n", p.name)
}

func main() {
	var m mover
	p1 := person{      // person类型的值
		"洪笳淏",
		23,
	}
	m = p1
	m.move()           // 洪笳淏在跑
	fmt.Println(m)     // {洪笳淏 18}

	p2 := &person{     // p2是person类型的指针
		"吴星怡",
		18,
	}
	m = p2
	m.move()          // 吴星怡在跑
	fmt.Println(m)    // &{吴星怡 18}


	var m1 mover1
	//m1 = p1           // 无法赋值，因为p1是person的值类型，没有实现mover1接口
	//m1.move1()
	m1 = p2
	m1.move1()          // 吴星怡在跑
	fmt.Println(m1)     // &{吴星怡 18}

	var s sayer       // 定义一个sayer类型的变量
	s = p2
	s.say()             // 吴星怡在叫
	fmt.Println(s)      // &{吴星怡 18}


	var a animal     // 定义一个animal类型的变量
	a = p2
	a.move1()        // 吴星怡在跑
	a.say()          // 吴星怡在叫
	fmt.Println(a)   // &{吴星怡 18}

}
