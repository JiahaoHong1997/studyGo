package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

func main() {
	// map是一种无序的基于key-value的数据结构，Go语言中的map是引用类型，必须初始化才能使用。
	// Go语言中 map的定义语法如下：
	// map[KeyType]ValueType

	// map类型的变量默认初始值为nil，需要使用make()函数来分配内存。语法为：
	// make(map[KeyType]ValueType, [cap])
	// 其中cap表示map的容量，该参数虽然不是必须的，但是我们应该在初始化map的时候就为其指定一个合适的容量。

	// map类型的变量默认初始值为nil，需要使用make()函数来分配内存。语法为：
	scoreMap := make(map[string]int, 8)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	fmt.Println(scoreMap)
	fmt.Println(scoreMap["小明"])
	fmt.Printf("type of scoreMap:%T\n", scoreMap)

	// map也支持在声明的时候填充元素，例如：
	userInfo := map[string]string{
		"username": "洪笳淏",
		"password": "hjh1314lvwxy",
	}
	fmt.Println(userInfo)

	// ############### 判断map中键是否存在 ####################
	value, ok := scoreMap["张三"]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("查无此人")
	}

	// ############## map的遍历 #################
	// Go语言中使用for range遍历
	// 遍历key和value
	scoreMap["娜扎"] = 60
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}

	// 只想遍历key的时候，可以按下面的写法：
	for k := range scoreMap {
		fmt.Println(k)
	}

	// ############## 删除键值对 ################
	// 使用delete()内建函数从map中删除一组键值对，delete()函数的格式如下：
	// delete(map, key)
	delete(scoreMap, "小明") //将小明:100从map中删除
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}

	// ############# 按照指定顺序遍历map ################
	var scoreMap1 = make(map[string]int, 200)

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i) //生成stu开头的字符串
		value := rand.Intn(100)          //生成0~99的随机整数
		scoreMap1[key] = value
	}
	// 取出map中的所有key存入切片keys
	var keys = make([]string, 0, 200)
	for key := range scoreMap1 {
		keys = append(keys, key)
	}
	// 对切片进行排序
	sort.Strings(keys)
	//按照排序后的key遍历map
	for _, key := range keys {
		fmt.Println(key, scoreMap1[key])
	}

	// ############## 元素为map类型的切片 ################
	var mapSlice = make([]map[string]string, 3) // 切片的元素为map
	// [nil nil nil]
	for index, value := range mapSlice {
		fmt.Printf("index:%v  value:%v\n", index, value)
	}
	fmt.Println("after init")
	// 对切片中的map元素进行初始化
	mapSlice[0] = make(map[string]string, 10) // 单独初始化切片元素
	mapSlice[0]["name"] = "洪笳淏"
	mapSlice[0]["password"] = "hjh1314lvwxy"
	mapSlice[0]["address"] = "太慈桥"
	for index, value := range mapSlice {
		fmt.Printf("index:%v  value:%v\n", index, value)
	}

	// ############## 值为切片类型的map #################
	var sliceMap = make(map[string][]string, 3)
	fmt.Println(sliceMap)
	fmt.Println("after init")
	key := "中国"
	v, ok := sliceMap[key]
	if !ok {
		v = make([]string, 0, 2) // 完成对切片的初始化
	}
	v = append(v, "北京", "上海")
	sliceMap[key] = v
	fmt.Println(sliceMap)

	// ############### 写一个程序，统计一个字符串中每个单词出现的次数。比如：”how do you do”中how=1 do=2 you=1。 ###############
	// 0.定义一个map[string]int
	var sMap = make(map[string]int, 50)
	s := "how do you do"
	// 1.字符串中有哪些单词
	s1 := strings.Split(s, " ")            // 将字符串转换成由单词组成的切片
	fmt.Printf("s1:%v  type:%T\n", s1, s1) // s1:[how do you do]  type:[]string
	// 2.遍历单词做统计
	for _, v1 := range s1 {
		_, ok := sMap[v1] // 该单词是否在map中出现过
		if ok {
			sMap[v1]++ // 如果出现过，对应的值加1
		} else {
			sMap[v1] = 1 // 否则置为1
		}
	}
	for k1, v1 := range sMap {
		fmt.Printf("key:%s  value:%d\n", k1, v1)
	}

	type Map map[string][]int
	m := make(Map)
	q := []int{1, 2}                 // 切片容量为2
	fmt.Printf("%d %p\n", cap(q), q) // 2 0xc0000b2250
	q = append(q, 3)                 // 切片容量变为4，地址改变
	fmt.Printf("%+v\n", q)           // [1 2 3]
	m["q1mi"] = q
	fmt.Printf("%+v\n", m["q1mi"])            // [1 2 3]
	q = append(q[:1], q[2:]...)               // 删除元素后，仅改变了切片长度，切片容量并没有减少
	fmt.Printf("%+v  %d  %p\n", q, cap(q), q) // [1 3]   4   0xc0000b4040
	fmt.Printf("%+v\n", m["q1mi"])            // 【1 3 3】


	// map引用赋值
	r := make(map[string]int, 7)
	r["hjh"] = 1
	r["wxy"] = 2
	u := r
	fmt.Println(r)  // map[hjh:1 wxy:2]
	fmt.Println(u)  // map[hjh:1 wxy:2]
	l := make(map[string]int, 7)
	l["zzz"] = 3
	r = l
	fmt.Println(r)  // map[zzz:3]
	fmt.Println(l)  // map[zzz:3]
	fmt.Println(u)  // map[hjh:1 wxy:2]

	// slice引用赋值
	o := []int{}
	o = append(o, 1,2,3,4,5)
	y := o
	fmt.Printf("add:%p, value:%v\n", &o, o) // add:0xc0000a6440, value:[1 2 3 4 5]
	fmt.Printf("add:%p, value:%v\n", &y, y) // add:0xc0000a6460, value:[1 2 3 4 5]
	e := []int{4,5,6,7,8}
	o = e
	fmt.Printf("add:%p, value:%v\n", &o, o) // add:0xc0000a6440, value:[4 5 6 7 8]
	fmt.Printf("add:%p, value:%v\n", &y, y) // add:0xc0000a6460, value:[1 2 3 4 5]
	fmt.Printf("add:%p, value:%v\n", &e, e) // add:0xc0000a64c0, value:[4 5 6 7 8]
}

