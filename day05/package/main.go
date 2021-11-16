package pkg1

import "fmt"

// 注意事项：

// 1.一个文件夹下面直接包含的文件只能归属一个package，同样一个package的文件不能在多个文件夹下。
// 2.包名可以不和文件夹的名字一样，包名不能包含 - 符号。
// 3.包名为main的包为应用程序的入口包，这种包编译后会得到一个可执行文件，而编译不包含main包的源代码则不会得到可执行文件。

// ############### 包变量可见性 ####################

var a = 100  // 首字母小写，外部包不可见，只能在当前包内使用

// 首字母大写外部包可见，可在其他包中使用
const Mode = 1

type person struct { // 首字母小写，外部包不可见，只能在当前包内使用
	name string
}

// 首字母大写，外部包可见，可在其他包中使用
func Add(x, y int) int {
	return x + y
}

func age() { // 首字母小写，外部包不可见，只能在当前包内使用
	var Age = 18 // 函数局部变量，外部包不可见，只能在当前函数内使用
	fmt.Println(Age)
}

// 结构体中的字段名和接口中的方法名如果首字母都是大写，外部包可以访问这些字段和方法。例如：
type Student struct {
	Name  string //可在包外访问的方法
	class string //仅限包内访问的字段
}

type Payer interface {
	init() //仅限包内访问的方法
	Pay()  //可在包外访问的方法
}