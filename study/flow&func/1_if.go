package main

/*
Go里面if条件判断语句中不需要括号
	if x > 10 {
		fmt.Println("x is greater than 10")
	} else {
		fmt.Println("x is less than 10")
	}

Go的if还有一个强大的地方就是条件判断语句里面允许声明一个变量，这个变量的作用域只能在该条件逻辑块内，其他地方就不起作用了
	if x := cal(); x > 10 {
		fmt.Println("x is greater than 10")
	} else {
		fmt.Println("x is less than 10")
	}
	fmt.Println(x) // //这个地方如果这样调用就编译出错了，因为x是条件里面的变量: unresolved reference

多个条件的时候
	if i := cal(); i == 3 {

	} else if i == 2 {

	} else {

	}

*/

func main() {
}
