package main

import "fmt"

//***************定义变量*****************
/*
使用var关键字是Go最基本的定义变量方式，与C语言不同的是Go把变量类型放在变量名后面：
	定义一个名称为“variableName”，类型为"type"的变量
	var variableName type

定义多个变量
	定义三个类型都是“type”的变量
	var vname1, vname2, vname3 type

定义变量并初始化值
	初始化“variableName”的变量为“value”值，类型是“type”
	var variableName type = value

同时初始化多个变量
	定义三个类型都是"type"的变量,并且分别初始化为相应的值vname1为v1，vname2为v2，vname3为v3
	var vname1, vname2, vname3 type = v1, v2, v3

也可以直接忽略类型声明
	定义三个变量，它们分别初始化为相应的值vname1为v1，vname2为v2，vname3为v3然后Go会根据其相应值的类型来帮你初始化它们
	var vname1, vname2, vname3 = v1, v2, v3

也可以继续简化
	定义三个变量，它们分别初始化为相应的值vname1为v1，vname2为v2，vname3为v3编译器会根据初始化的值自动推导出相应的类型
	vname1, vname2, vname3 := v1, v2, v3
	:=这个符号直接取代了var和type,这种形式叫做简短声明。不过它有一个限制，那就是它只能用在函数内部；
	在函数外部使用则会无法编译通过，所以一般用var方式来定义全局变量

_（下划线）是个特殊的变量名，任何赋予它的值都会被丢弃
	我们将值35赋予b，并同时丢弃34
	_, b := 34, 35

同时声明多个常量、变量，或者导入多个包时，可采用分组的方式进行声明。
	var (
		i int
		pi float32
		prefix string
	)

*Go对于已声明但未使用的变量会在编译阶段报错: 声明了i但未使用

*/

func main() {
	fmt.Println("VARIABLE")
}
