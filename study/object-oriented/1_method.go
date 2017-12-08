package main

import (
	"fmt"
	"math"
)

/*
函数的另一种形态，带有接收者的函数，我们称为method

"A method is a function with an implicit first argument, called a receiver."

method的语法如下：
func (r ReceiverType) funcName(parameters) (results)

在使用method的时候重要注意几点
	虽然method的名字一模一样，但是如果接收者不一样，那么method就不一样
	method里面可以访问接收者的字段
	调用method通过.访问，就像struct里面访问字段一样

默认此处方法的Receiver是以值传递，而非引用传递，是的，Receiver还可以是指针,
两者的差别在于, 指针作为Receiver会对实例对象的内容发生操作,而普通类型作为Receiver仅仅是以副本作为操作对象,并不对原实例对象发生操作。


*/

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

// func
func area(r Rectangle) float64 {
	return r.height * r.width
}

// method
func (r Rectangle) area() float64 {
	return r.width * r.height
}

// method
func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func main() {
	r1 := Rectangle{10.5, 90.2}
	r2 := Rectangle{8.5, 103.2}

	fmt.Println(area(r1))
	fmt.Println(area(r2))

	fmt.Println(r1.area())
	fmt.Println(r2.area())

	c1 := Circle{23.312}
	fmt.Println(c1.area())
}
