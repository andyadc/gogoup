package main

import (
	"fmt"
	"os"
)

//  ******** func *********
/*
函数是Go里面的核心设计，它通过关键字func来声明，它的格式如下：
	func funcName(input1 type1, input2 type2) (output1 type1, output2 type2) {
		//这里是处理逻辑代码
		//返回多个值
		return value1, value2
	}

上面的代码我们看出
	关键字func用来声明一个函数funcName
	函数可以有一个或者多个参数，每个参数后面带有类型，通过,分隔
	函数可以返回多个值
	上面返回值声明了两个变量output1和output2，如果你不想声明也可以，直接就两个类型
	如果只有一个返回值且不声明返回值变量，那么你可以省略 包括返回值 的括号
	如果没有返回值，那么就直接省略最后的返回信息
	如果有返回值， 那么必须在函数的外层添加return语句

变参
func myfunc(arg ...int) {}

传值与传指针
当我们传一个参数值到被调用函数里面时，实际上是传了这个值的一份copy，当在被调用函数中修改参数值的时候，调用函数中相应实参不会发生任何变化，因为数值变化只作用在copy上。

变量在内存中是存放于一定地址上的，修改变量实际是修改变量地址处的内存。
只有add1函数知道x变量所在的地址，才能修改x变量的值。
所以我们需要将x所在地址&x传入函数，并将函数的参数的类型由int改为*int，即改为指针类型，才能在函数中修改x变量的值。
此时参数仍然是按copy传递的，只是copy的是一个指针。

传指针有什么好处呢?
	传指针使得多个函数能操作同一个对象。
	传指针比较轻量级 (8bytes),只是传内存地址，我们可以用指针传递体积大的结构体。如果用参数值传递的话, 在每次copy上面就会花费相对较多的系统开销（内存和时间）。所以当你要传递大的结构体的时候，用指针是一个明智的选择。
	Go语言中channel，slice，map这三种类型的实现机制类似指针，所以可以直接传递，而不用取地址后传递指针。（注：若函数需改变slice的长度，则仍需要取地址传递指针）

defer
延迟（defer）语句，可以在函数中添加多个defer语句。
当函数执行到最后时，这些defer语句会按照逆序执行，最后该函数返回。

函数作为值、类型
在Go中函数也是一种变量，我们可以通过type来定义它，它的类型就是所有拥有相同的参数，相同的返回值的一种类型
	type typeName func(input1 inputType1 , input2 inputType2 [, ...]) (result1 resultType1 [, ...])

Panic
	是一个内建函数，可以中断原有的控制流程，进入一个令人恐慌的流程中。
	当函数F调用panic，函数F的执行被中断，但是F中的延迟函数会正常执行，然后F返回到调用它的地方。
	在调用的地方，F的行为就像调用了panic。这一过程继续向上，直到发生panic的goroutine中所有调用的函数返回，
	此时程序退出。恐慌可以直接调用panic产生。也可以由运行时错误产生，例如访问越界的数组。

Recover
	是一个内建的函数，可以让进入令人恐慌的流程中的goroutine恢复过来。recover仅在延迟函数中有效。
	在正常的执行过程中，调用recover会返回nil，并且没有其它任何效果。
	如果当前的goroutine陷入恐慌，调用recover可以捕获到panic的输入值，并且恢复正常的执行。

main函数和init函数
Go里面有两个保留的函数：init函数（能够应用于所有的package）和main函数（只能应用于package main）。
这两个函数在定义时不能有任何的参数和返回值;
一个package里面可以写任意多个init函数;

Go程序会自动调用init()和main()，所以你不需要在任何地方调用这两个函数。
每个package中的init函数都是可选的，但package main就必须包含一个main函数。

初始化执行过程
程序的初始化和执行都起始于main包。
如果main包还导入了其它的包，那么就会在编译时将它们依次导入。
有时一个包会被多个包同时导入，那么它只会被导入一次（例如很多包可能都会用到fmt包，但它只会被导入一次，因为没有必要导入多次）。
当一个包被导入时，如果该包还导入了其它的包，那么会先将其它包导入进来，
然后再对这些包中的包级常量和变量进行初始化，接着执行init函数（如果有的话），依次类推。
等所有被导入的包都加载完毕了，就会开始对main包中的包级常量和变量进行初始化，
然后执行main包中的init函数（如果存在的话），最后执行main函数。


*/

func main() {
	// testDefer()

	//fmt.Println(isOdd(1))
	//fmt.Println(isEven(1))

	slice := []int{1, 3, 6, 7, 9}
	fmt.Println(slice)

	m := filter(slice, isOdd)
	fmt.Println(m)

	//testPanic()

	// panic recover
	f := testPanic
	b := throwsPanic(f)
	fmt.Println(b)
}

// init
func init() {
	fmt.Println("init...")
}

func testPanic() {
	user := os.Getenv("USER")
	fmt.Println(user)
	if user == `` {
		panic("no value for $USER")
	}
}

func throwsPanic(f func()) (b bool) {
	defer func() {
		if x := recover(); x != nil {
			b = true
		}
	}()
	f() // 执行函数f，如果f中出现了panic，那么就可以恢复回来
	return
}

/*
函数当做值和类型在我们写一些通用接口的时候非常有用，通过上面例子我们看到testInt这个类型是一个函数类型，
然后两个filter函数的参数和返回值与testInt类型是一样的，
但是我们可以实现很多种的逻辑，这样使得我们的程序变得非常的灵活。
*/

type testInt func(int) bool // 声明了一个函数类型

// 声明的函数类型在这个地方当做了一个参数
func filter(slice []int, f testInt) []int {
	var result []int

	for _, v := range slice {
		if f(v) {
			result = append(result, v)
		}
	}

	return result
}

func isOdd(i int) bool {
	return i%2 != 0
}

func isEven(i int) bool {
	return i%2 == 0
}

// defer
func testDefer() {
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
}

/*
	a, b := 1, 1

	a1 := add1(a)
	fmt.Println(a)
	fmt.Println(a1)

	b1 := add2(&b)
	fmt.Println(b)
	fmt.Println(b1)
*/

// 传值
func add1(a int) int {
	a = a + 1
	return a
}

// 传指针
func add2(a *int) int {
	*a = *a + 1
	return *a
}

// 变参
func sum(arg ...int) {
	for _, v := range arg {
		fmt.Println(v)
	}
}

// 多个返回值
func SumAndProduct(a, b int) (int, int) {
	return a + b, a * b
}

func SumAndProduct2(A, B int) (add int, Multiplied int) {
	add = A + B
	Multiplied = A * B
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
