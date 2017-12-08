package main

import (
	"fmt"
	"strconv"
)

/*
空interface
空interface(interface{})不包含任何的method，正因为如此，所有的类型都实现了空interface。

空interface对于描述起不到任何的作用(因为它不包含任何的method），
但是空interface在我们需要存储任意类型的数值的时候相当有用，因为它可以存储任意类型的数值。
	// 定义a为空接口
	var a interface{}
	var i int = 5
	s := "Hello world"
	// a可以存储任意类型的数值
	a = i
	a = s
一个函数把interface{}作为参数，那么他可以接受任意类型的值作为参数，
如果一个函数返回interface{},那么也就可以返回任意类型的值。

interface的变量里面可以存储任意类型的数值(该类型实现了interface)。
那么我们怎么反向知道这个变量里面实际保存了的是哪个类型的对象呢？目前常用的有两种方法：
Comma-ok断言
	Go语言里面有一个语法，可以直接判断是否是该类型的变量： value, ok = element.(T)，
	这里value就是变量的值，ok是一个bool类型，element是interface变量，T是断言的类型。

switch测试

*/

var a interface{}
var s = `kong`

type ELement interface{}
type List []ELement

type People struct {
	name string
	age  int
}

func (p People) String() string {
	return "(name: " + p.name + " - age: " + strconv.Itoa(p.age) + " years)"
}

func main() {
	list := make(List, 3)
	list[0] = 1
	list[1] = "Hello"
	list[2] = People{"Andy", 19}

	//useCommaOk(list)
	useSwitch(list)

}

// Comma-ok断言
func useCommaOk(list []ELement) {
	for index, element := range list {
		if value, ok := element.(int); ok {
			fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
		} else if value, ok := element.(string); ok {
			fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
		} else if value, ok := element.(People); ok {
			fmt.Printf("list[%d] is a User and its value is %s\n", index, value)
		} else {
			fmt.Printf("list[%d] is of a different type\n", index)
		}
	}
}

// switch测试
/*
***有一点需要强调的是：`element.(type)`语法不能在switch外的任何逻辑里面使用，
如果你要在switch外面判断一个类型就使用`comma-ok`。
*/
func useSwitch(list []ELement) {
	for index, element := range list {
		switch value := element.(type) {
		case int:
			fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
		case string:
			fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
		case People:
			fmt.Printf("list[%d] is a User and its value is %s\n", index, value)
		default:
			fmt.Printf("list[%d] is of a different type\n", index)
		}
	}
}
