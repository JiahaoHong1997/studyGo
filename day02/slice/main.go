package main

import (
	"fmt"
	"sort"
)

func main() {
	// 因为数组的长度是固定的并且数组长度属于类型的一部分，所以数组有很多的局限性。
	// 切片（Slice）是一个拥有相同类型元素的可变长度的序列。它是基于数组类型做的一层封装。它非常灵活，支持自动扩容。
	// 切片是一个引用类型，它的内部结构包含地址、长度和容量。切片一般用于快速地操作一块数据集合。
	// 声明切片类型的基本语法如下：
	// var name  []T
	// 切片是一个引用类型，必须初始化后才能使用
	var a []string              //声明一个字符串切片但并未初始化
	var b = []int{}             //声明一个整型切片并初始化
	var c = []bool{false, true} //声明一个布尔切片并初始化
	//var d = []bool{false, true} //声明一个布尔切片并初始化
	fmt.Println(a)        //[]
	fmt.Println(b)        //[]
	fmt.Println(c)        //[false true]
	fmt.Println(a == nil) //true
	fmt.Println(b == nil) //false
	fmt.Println(c == nil) //false
	// fmt.Println(c == d)   //切片是引用类型，不支持直接比较，只能和nil比较

	// 切片拥有自己的长度和容量，我们可以通过使用内置的len()函数求长度，使用内置的cap()函数求切片的容量。

	// 切片表达式
	// 简单切片表达式
	// 切片表达式中的low和high表示一个索引范围（左包含，右不包含），也就是下面代码中从数组a中选出1<=索引值<4的元素组成切片s，
	// 得到的切片长度=high-low，容量等于得到的切片的底层数组的容量。
	// 基于数组定义切片
	d := [5]int{1, 2, 3, 4, 5}
	s := d[1:3] // 左闭右开,s的底层数组为{2,3,4,5},所以容量为4，容量长度从切片首部开始计算
	fmt.Printf("%T\n", s)
	fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s)) // s:[2 3] len(s):2 cap(s):4

	// make函数构造切片
	e := make([]int, 5, 10) // make(切片类型, 切片长度(可省略), 切片容量(可省略))
	fmt.Printf("d:%v, 类型：%T\n", e, e)

	// nil
	var f []int         // 声明int类型切片
	var g = []int{}     // 声明并且初始化
	h := make([]int, 0) // 声明并且初始化
	if f == nil {       // true
		fmt.Println("f是一个nil")
	}
	if g == nil { // false
		fmt.Println("g是一个nil")
	}
	if h == nil { // false
		fmt.Println("h是一个nil")
	}

	fmt.Printf("f:%v, len(f):%d, cap(f):%d\n", f, len(f), cap(f))
	fmt.Printf("g:%v, len(g):%d, cap(g):%d\n", g, len(g), cap(g))
	fmt.Printf("h:%v, len(h):%d, cap(h):%d\n", h, len(h), cap(h))
	// 所以想要知道一个切片是否为空，不能通过是否等于nil来判断，必须通过查看len()来知道
	// 所以要判断一个切片是否是空的，要是用len(s) == 0来判断，不应该使用s == nil来判断。

	// #################### 切片的赋值拷贝 ######################
	i := make([]int, 3) // [0, 0, 0]
	j := i
	j[0] = 100
	fmt.Printf("i:%v\n", i) // [100, 0, 0]
	fmt.Printf("j:%v\n", j) // [100, 0, 0]
	// i和j切片的值均发生了变化，因为他们的底层数组是一样的，他们都只相当于指向底层数组的指针

	// #################### 切片的遍历 #########################
	// 通过索引来遍历
	k := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(k); i++ {
		fmt.Println(i, k[i])
	}
	fmt.Println()
	// for range 遍历
	for index, value := range k{
		fmt.Println(index, value)
	}

	// ################## append函数来扩容 ##########################
	//切片每次扩容后，指向的地址空间都会发生变化
	//往一个切片中不断添加元素的过程，类似于公司搬家，公司发展初期，资金紧张，人员很少，所以只需要很小的房间即可容纳所有的员工，
	//随着业务的拓展和收入的增加就需要扩充工位，但是办公地的大小是固定的，无法改变，因此公司只能选择搬家，每次搬家就需要将所有的人员转移到新的办公点。
	//员工和工位就是切片中的元素。
	//办公地就是分配好的内存。
	//搬家就是重新分配内存。
	//无论搬多少次家，公司名称始终不会变，代表外部使用切片的变量名不会修改。
	//由于搬家后地址发生变化，因此内存“地址”也会有修改。
	var l []int  // 此时只是声明了一个切片，还没有申请内存，切片要初始化之后才能使用
	//l[0] = 100    没有初始化，会爆超出内存的错误
	l = append(l, 10)  // 一定要赋值
	fmt.Println(l)

	for i:=0;i<10;i++{
		l = append(l, i)
		fmt.Printf("l:%v  len(l):%d  cap(l):%d  prt:%p\n", l , len(l), cap(l), l)
	}

	// append一次可追加多个元素
	l = append(l, 11,12,13,14,15)
	fmt.Println(l)

	m := []int{16,17,18,19,20}
	l = append(l, m...)  // 将切片m加入到l中
	fmt.Println(l)

	// ################### 切片的copy ########################
	// copy后的切片指向新的数组地址
	n := []int{1, 2, 3, 4, 5}
	o := make([]int, 5, 5)
	p := o
	copy(o, n)
	o[0] = 100
	fmt.Println(n)  // [0, 1, 2, 3, 4, 5]
	fmt.Println(o)  // [100, 1, 2, 3, 4, 5]
	fmt.Println(p)  // [100, 1, 2, 3, 4, 5]

	// ################### 切片删除元素 ########################
	// 删除索引为index的元素：append(q[:index], q[index+1:]...)
	q := []string{"北京", "上海", "广州", "深圳"}
	q = append(q[0:3], q[4:]...)  // 删除深圳
	fmt.Println(q)   // [北京 上海 广州]

	// ################## 在切片中某一位置添加元素/切片 ####################
	q = append(q[:2], append([]string{"成都"}, q[2:]...)...)  // 在索引为2处添加元素"成都"
	fmt.Println(q)   // [北京 上海 成都 广州]

	q = append(q[:2], append([]string{"厦门", "重庆"}, q[2:]...)...)  // 在索引为2处添加元素"厦门"，"重庆"
	fmt.Println(q)   // [北京 上海 厦门 重庆 成都 广州]


	var r = make([]string, 5, 10)
	for i := 0; i < 10; i++ {
		r = append(r, fmt.Sprintf("%v", i))
	}
	fmt.Println(r)  // [     0 1 2 3 4 5 6 7 8 9]

	// 使用内置的sort包对数组var s = [...]int{3, 7, 8, 9, 1}进行排序
	var t = [...]int{3,7,8,9,1}
	sort.Ints(t[:])   // t[:]得到的是一个切片，指向底层数组t
	fmt.Println(t)

}
