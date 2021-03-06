package main

import "fmt"

// ********** array **********
/*
array就是数组，它的定义方式如下：
var arr [n]type

在[n]type中，n表示数组的长度，type表示存储元素的类型。通过[]来进行读取或赋值：
	var arr [3]int // 声明了一个int类型的数组
	arr[1] = 1
	arr[0] = 3 // 数组下标是从0开始的

	fmt.Println(arr) // 未赋值元素，默认返回0

***由于长度也是数组类型的一部分，因此[3]int与[4]int是不同的类型，数组也就不能改变长度。
数组之间的赋值是值的赋值，即当把一个数组作为参数传入函数的时候，传入的其实是该数组的副本，而不是它的指针。

数组可以使用另一种 := 来声明
	a := [3]int{1, 2, 3}   // 声明了一个长度为3的int数组
	b := [10]int{1, 2, 3}  // 声明了一个长度为10的int数组，其中前三个元素初始化为1、2、3，其它默认为0
	c := [...]int{1, 2, 3} // 可以省略长度而采用`...`的方式，Go会自动根据元素个数来计算长度

Go支持嵌套数组，即多维数组。
	// 声明了一个二维数组，该数组以两个数组作为元素，其中每个数组中又有4个int类型的元素
	doubleArray := [2][4]int{[4]int{1, 2, 3, 4}, [4]int{5, 6, 7, 8}}
	// 上面的声明可以简化，直接忽略内部的类型
	easyArray := [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}


*/

func main() {
	// 声明了一个二维数组，该数组以两个数组作为元素，其中每个数组中又有4个int类型的元素
	doubleArray := [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
	// 上面的声明可以简化，直接忽略内部的类型
	easyArray := [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
	fmt.Println(doubleArray)
	fmt.Println(easyArray)

	array_point()
	array_copy()
}

func array_point() {
	// 声明包含5个元素的指向整数的数组
	// 用整型指针初始化索引为0和3的数组元素
	array := [5]*int{0: new(int), 3: new(int)}

	// 为索引为0的元素赋值
	*array[0] = 10
	fmt.Println(array)
}

// 把同样类型的一个数组赋值给另外一个数组
func array_copy() {
	var array1 [5]string
	array2 := [5]string{"red", "orange", "yellow"}

	array1 = array2

	fmt.Println(array1)
	fmt.Println(array2)
}
