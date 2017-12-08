package main

import "fmt"

/*
method不只是只能作用在struct上面, 以定义在任何你自定义的类型、内置类型、struct等各种类型上面。
struct只是自定义类型里面一种比较特殊的类型而已，还有其他自定义类型申明，可以通过如下这样的申明来实现。
	type typeName typeLiteral
*/

type ages int
type money float32
type months map[string]int

func main() {
	m := months{"January": 31,
		"February": 28}
	fmt.Println(m)
}
