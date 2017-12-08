package main

import "fmt"

/*

interface就是一组抽象方法的集合，它必须由其他非interface类型实现，而不能自我实现，
Go通过interface实现了duck-typing:即"当看到一只鸟走起来像鸭子、游泳起来像鸭子、叫起来也像鸭子，那么这只鸟就可以被称为鸭子"。

*/

type Person struct {
	name  string
	age   int
	phone string
}

type Man struct {
	Person
	company string
}

type Woman struct {
	Person
	address string
}

func (p Person) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", p.name, p.phone)
}

func (p Person) Sing(lyrics string) {
	fmt.Println("La la la la...", lyrics)
}

func (m Man) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", m.name,
		m.company, m.phone)
}

type User interface {
	SayHi()
	Sing(lyrics string)
}

func main() {
	mike := Man{Person{"Mike", 28, "112"}, "IBM"}
	lucy := Woman{Person{"Lucy", 28, "112"}, "Hollyword"}

	// 定义User类型的变量u
	var u User

	// u能存储Man
	u = mike
	u.SayHi()

	// u能存储Woman
	u = lucy
	u.Sing("ha li lu ya")

	fmt.Println(`----------------------------`)

	x := []User{mike, lucy}

	for _, v := range x {
		v.SayHi()
	}

}
