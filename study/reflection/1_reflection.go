package main

import (
	"fmt"
	"reflect"
)

// ********* Reflection *********
/*
Go语言实现了反射，所谓反射就是能检查程序在运行时的状态。

*/

func main() {
	i := 108
	t := reflect.TypeOf(i)  // 得到类型的元数据,通过t我们能获取类型定义里面的所有元素
	v := reflect.ValueOf(i) // 得到实际的值，通过v我们获取存储在里面的值，还可以去改变值

	fmt.Println(t)
	fmt.Println(v)

	a := 13.14
	rv := reflect.ValueOf(&a)
	ev := rv.Elem()
	ev.SetFloat(19.08)

	fmt.Println(a)

	//tag := t.Elem().Field(0).Tag
	//name := v.Elem().Field(0).String()
	//
	//fmt.Println(tag)
	//fmt.Println(name)
}
