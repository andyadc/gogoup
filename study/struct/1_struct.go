package main

import "fmt"

// ********strcut********
/*
Go语言中，也和C或者其他语言一样，我们可以声明新的类型，作为其它类型的属性或字段的容器。
我们可以创建一个自定义类型person代表一个人的实体。这个实体拥有属性：姓名和年龄。这样的类型我们称之struct。
	type Person struct {
		name string
		age int
	}

	var P person  // P现在就是person类型的变量了
	P.name = "andyadc"  // 赋值"Astaxie"给P的name属性.
	P.age = 25  // 赋值"25"给变量P的age属性
	fmt.Printf("The person's name is %s", P.name)  // 访问P的name属性.

除了上面这种P的声明使用之外，还有另外几种声明使用方式：
1.按照顺序提供初始化值
	p1 := Person{"Tom", 19}
2.通过field:value的方式初始化，这样可以任意顺序
	p2 := Person{name:"Jack", age:17}
3.当然也可以通过new函数分配一个指针，此处P的类型为*person
	p3 := new(Person)


struct的匿名字段

可以重载通过匿名字段继承的一些字段

*/

type Skills []string

type Person struct {
	name string
	age  int
}

type Student struct {
	Person     // 匿名字段，那么默认Student就包含了Human的所有字段
	Skills     // 匿名字段，自定义的类型string slice
	int        // 内置类型作为匿名字段
	speciality string
}

func main() {
	mark := Student{}
	mark.int = 1

	fmt.Println(mark)

	mark.Person = Person{"Pola", 12}

	fmt.Println(mark.name)
}

// 比较两个人的年龄，返回年龄大的那个人，并且返回年龄差
// struct也是传值的
func Older(p1, p2 Person) (Person, int) {
	d := p1.age - p2.age
	if d > 0 {
		return p1, d
	}
	return p2, p2.age - p1.age
}
