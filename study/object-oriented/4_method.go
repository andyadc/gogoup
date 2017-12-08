package main

import "fmt"

/*
method继承
method也是可以继承的。如果匿名字段实现了一个method，那么包含这个匿名字段的struct也能调用该method。

method重写

*/

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human
	school string
}

type Employee struct {
	Human   //匿名字段
	company string
}

func (h *Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

// method重写
func (e Employee) SayHi() {
	fmt.Printf("Hi, I am %s, come from %s\n", e.name, e.company)
}

func main() {
	jack := Student{Human{"Jack", 23, "110"}, "Harvard"}
	tom := Employee{Human{"Tome", 43, "112"}, "Google"}

	jack.SayHi()
	tom.SayHi()
}
