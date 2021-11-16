package main

import (
	"fmt"
	"reflect"
)

type myInt int64

// ############# 变量的内在机制 ##################
// Go语言中变量是分为两部分的：1.类型信息：预先定义好的元信息。  2.值信息：程序运行过程中可动态变化的。


// ############# reflect包 #################
// 在Go语言中反射的相关功能由内置的reflect包提供，任意接口值在反射中都可以理解为由reflect.Type和reflect.Value两部分组成，
// 并且reflect包提供了reflect.TypeOf和reflect.ValueOf两个函数来获取任意对象的Value和Type

// ############# TypeOf ################
// 使用reflect.TypeOf()函数可以获得任意值的类型对象（reflect.Type），程序通过类型对象可以访问任意值的类型信息。
func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("type:%v\n", v)
}

// ############# type name和type kind ###############
// Go语言的反射中像数组、切片、Map、指针等类型的变量，它们的.Name()都是返回空
func reflectType1(x interface{}) {
	t := reflect.TypeOf(x)
	fmt.Printf("type:%v  kind:%v\n", t.Name(), t.Kind())
}

// ############# ValueOf ################
// reflect.ValueOf()返回的是reflect.Value类型，其中包含了原始值的值信息。reflect.Value与原始值之间可以互相转换。
// 通过反射获取值
func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind()
	switch k {
	case reflect.Int64:
		// v.Int()从反射中获取整型的原始值，然后通过int64()强制类型转换
		fmt.Printf("type is int64, value is %d\n", int64(v.Int()))
	case reflect.Float32:
		// v.Float()从反射中获取浮点型的原始值，然后通过float32()强制类型转换
		fmt.Printf("type is float32, value is %f\n", float32(v.Float()))
	case reflect.Float64:
		// v.Float()从反射中获取浮点型的原始值，然后通过float64()强制类型转换
		fmt.Printf("type is float64, value is %f\n", float64(v.Float()))
	}
}

// 通过反射设置变量的值
// 想要在函数中通过反射修改变量的值，需要注意函数参数传递的是值拷贝，必须传递变量地址才能修改变量值。而反射中使用专有的Elem()方法来获取指针对应的值。
func reflectSetValue1(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(200) //修改的是副本，reflect包会引发panic
	}
}
func reflectSetValue2(x interface{}) {
	v := reflect.ValueOf(x)
	// 反射中使用 Elem()方法获取指针对应的值
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200)
	}
}


func main(){
	var a float32 = 3.14
	reflectType(a)    // type:float32
	var b int64 = 100
	reflectType(b)    // type:int64

	// ###########################################################################
	var c *float32  // 指针
	var d myInt		// 自定义类型
	var e rune		// 类型别名
	reflectType1(c)   // type:  kind:ptr
	reflectType1(d)   // type:myInt  kind:int64
	reflectType1(e)   // type:int32  kind:int32

	type person struct {
		name string
		age int
	}
	type book struct {title string}
	var f = person{
		name:"洪笳淏",
		age:23,
	}
	var g = book{"《情商》"}
	reflectType1(f)   // type:person  kind:struct
	reflectType1(g)   // type:book  kind:struct

	// ###########################################################################
	var h float32 = 3.14
	var i int64 = 100
	reflectValue(h)   // type is float32, value is 3.140000
	reflectValue(i)   // type is int64, value is 100
	// 将int类型的原始值转换为reflect.Value类型
	j := reflect.ValueOf(10)
	fmt.Printf("type j :%T\n", j) // type j :reflect.Value

	// ###########################################################################
	var k int64 = 100
	// reflectSetValue1(k) //panic: reflect: reflect.Value.SetInt using unaddressable value
	reflectSetValue2(&k)
	fmt.Println(k)   // 200
}
