package main

import "fmt"

// 数组必须指定存放的元素类型和容量(长度)
// 数组的长度是存放的元素类型的一部分
func main() {
	var a1 [3]bool
	var a2 [4]bool

	fmt.Printf("a1:%T, a2:%T\n", a1, a2)

	//############### 数组的初始化 ###################
	// 如果不初始化，默认为0(false)值
	// 1. 初始化方式1
	fmt.Println(a1)
	a1 = [3]bool{true, true, true}
	fmt.Println(a1)

	// 2. 初始化方式2
	// 根据初始值自动判断数组的长度是多少
	b := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(b)

	// 3. 使用指定索引值的方式来初始化数组
	c := [...]int{1: 1, 3: 5} // 将索引1和3的位置分别初始化为1和5
	fmt.Println(c)
	fmt.Printf("type of a:%T\n", c)

	//############## 遍历数组的两种方法 ####################
	var d = [...]string{"北京", "上海", "深圳"}
	// 方法1:for循环遍历
	for i := 0; i < len(d); i++ {
		fmt.Println(d[i])
	}

	// 方法2:for range遍历
	for index, value := range d{
		fmt.Println(index, value)
	}

	// ############# 多维数组 ##################
	e := [3][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	fmt.Println(e)   // 打印整个数组
	fmt.Println(e[2][1])  // 支持索引取值:重庆

	// ####二维数组的遍历####
	// 方法1: for range 循环遍历
	for _, v1 := range e{
		for _, v2 := range v1{
			fmt.Printf("%s\t", v2)
		}
	}
	fmt.Println()

	// 方法2: for循环遍历
	for i:=0;i<len(e);i++{
		for j:=0;j<len(e[i]);j++{
			fmt.Printf("%s\t", e[i][j])
		}
	}
	// 注意： 多维数组只有第一层可以使用...来让编译器推导数组长度。例如：
	e1 := [...][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}

	//e2 := [3][...]string{    // 错误写法
	//	{"北京", "上海"},
	//	{"广州", "深圳"},
	//	{"成都", "重庆"},
	//}
	fmt.Println(e1)
	//fmt.Println(e2)

	// ############## 数组是值类型 #####################
	// 数组是值类型，赋值和传参会复制整个数组。因此改变副本的值，不会改变本身的值。
	f1 := [3]int{10, 20, 30}
	modifyArray(f1) //在modify中修改的是a的副本x
	fmt.Println(f1) //[10 20 30]   不改变原始数值

	f2 := [3][2]int{
		{1, 1},
		{1, 1},
		{1, 1},
	}
	modifyArray2(f2)  //在modify中修改的是b的副本x
	fmt.Println(f2)   //[[1 1] [1 1] [1 1]]  不改变原始数值

	// 注意：
	//数组支持 “==“、”!=” 操作符，因为内存总是被初始化过的。
	//[n]*T表示指针数组，*[n]T表示数组指针 。

	// 	求数组所有元素的和
	var g = [...]int{1, 3, 5, 7, 8}
	var sum int = 0
	for _, v := range g{
		sum += v
	}
	fmt.Printf("数组g的和为%d\n", sum)

	// 找出数组中和为指定值的两个元素的下标
	var trg int = 8
	var i int = 0
	var j int = len(g)-1
	var ret [10][2]int
	var n int = 0
	for {
		if i == j{
			break
		}else if g[i] + g[j] > trg{
			j--
		}else if g[i] + g[j] == trg{
			ret[n][0] = i
			ret[n][1] = j
			n++
			i++
		}else{
			i++
		}
	}
	for _, v1 := range ret{
		if v1[0] == v1[1]{
			break
		}
		fmt.Println(v1)
	}
}

func modifyArray(x [3]int){
	x[0] = 100
}

func modifyArray2(x [3][2]int){
	x[2][0] = 100
}