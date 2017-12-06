package main

import "fmt"

//***************常量*****************
/*
在Go程序中，常量可定义为数值、布尔值或字符串等类型。

	const constName = value
	如果需要，也可以明确指定常量的类型：
	const Pi float32 = 3.1415926

	const i = 10000
	const MaxThread = 10
	const prefix = "adc"

在Go语言中，同时声明多个常量、变量，或者导入多个包时，可采用分组的方式进行声明。

	const(
		i = 100
		pi = 3.1415
		prefix = "Go_"
	)


iota枚举
Go里面有一个关键字iota，这个关键字用来声明enum的时候采用，它默认开始值是0，const中每增加一行加1

除非被显式设置为其它值或iota，每个const分组的第一个常量被默认设置为它的0值，
第二及后续的常量被默认设置为它前面那个常量的值，如果前面那个常量的值是iota，则它也被设置为iota。
*/

const Pi float32 = 3.1415926

const (
	x = iota // x = 0
	y = iota // y = 1
	z = iota // z = 2
	w        // 常量声明省略值时，默认和之前一个值的字面相同。这里隐式地说w = iota，因此w = 3。其实上面y和z可同样不用"= iota"
)

const v = iota // 每遇到一个const关键字，iota就会重置，此时v = 0

const (
	h, i, j = iota, iota, iota // h = 0, i = 0, j = 0 iota在同一行则值相同
)

const (
	a       = iota // a = 0
	b       = "B"
	c       = iota             // c = 2
	d, e, f = iota, iota, iota // d = 3, e = 3, f = 3
	g       = iota             // g = 4
)

func main() {
	fmt.Println("CONST")
	fmt.Println(Pi)

	fmt.Println(x, y, z, w)
	fmt.Println(v)
	fmt.Println(h, i, j)
	fmt.Println(a, b, c, d, e, f, g)
}
