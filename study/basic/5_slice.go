package main

import "fmt"

// ********** slice **********

/*
在初始定义数组时，我们并不知道需要多大的数组，因此我们就需要“动态数组”。在Go里面这种数据结构叫slice
slice并不是真正意义上的动态数组，而是一个引用类型。
slice总是指向一个底层array，slice的声明也可以像array一样，只是不需要长度。
	// 和声明array一样，只是少了长度
	var fslice []int

	// 声明一个slice，并初始化数据
	slice := []byte{'a', 'v', 'f', 'o'}

slice可以从一个数组或一个已经存在的slice中再次声明。
slice通过array[i:j]来获取，其中i是数组的开始位置，j是结束位置，但不包含array[j]，它的长度是j-i。
	// 声明一个含有10个元素元素类型为byte的数组
	var arr = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	// 声明两个含有byte的slice
	var a, b []byte

	// a指向数组的第1个元素开始，并到第7个元素结束，
	a = arr[1:7]
	// b是数组ar的另一个slice
	b = arr[1:]

***注意slice和数组在声明时的区别：声明数组时，方括号内写明了数组的长度或使用...自动计算长度，
而声明slice时，方括号内没有任何字符。
***slice是引用类型，所以当引用改变其中元素的值时，其它的所有引用都会改变该值。

slice有一些简便的操作
	slice的默认开始位置是0，ar[:n]等价于ar[0:n]
	slice的第二个序列默认是数组的长度，ar[n:]等价于ar[n:len(ar)]
	如果从一个数组里面直接获取slice，可以这样ar[:]，因为默认第一个序列是0，第二个是数组的长度，即等价于ar[0:len(ar)]

从概念上面来说slice像一个结构体，这个结构体包含了三个元素：
	一个指针，指向数组中slice指定的开始位置
	长度，即slice的长度
	最大长度，也就是slice开始位置到数组的最后位置的长度

slice有几个有用的内置函数：
	len 获取slice的长度
	cap 获取slice的最大容量
	append 向slice里面追加一个或者多个元素，然后返回一个和slice一样类型的slice
		c := append(a, 'x')
	copy 函数copy从源slice的src中复制元素到目标dst，并且返回复制的元素的个数
		d := []byte{'a', 'a', 'a'}
		i := copy(d, []byte{'x','y','z'})

***append函数会改变slice所引用的数组的内容，从而影响到引用同一数组的其它slice。
但当slice中没有剩余空间（即(cap-len) == 0）时，此时将动态分配新的数组空间。
返回的slice数组指针将指向这个空间，而原数组的内容将保持不变；其它引用此数组的slice则不受影响。

从Go1.2开始slice支持了三个参数的slice，之前我们一直采用这种方式在slice或者array基础上来获取一个slice
	var arr = [10]byte{}
	slice := arr[2:4]
	这个例子里面slice的容量是8，新版本里面可以指定这个容量
	slice = array[2:4:7]
	上面这个的容量就是7-2，即5。这样这个产生的新的slice就没办法访问最后的三个元素。

如果slice是这样的形式array[:i:j]，即第一个参数为空，默认值就是0。
*/

func main() {
	var arr = [10]byte{}
	slice := arr[2:4]
	fmt.Println(slice)

	slice2 := arr[2:4:5]
	fmt.Println(slice2)
}
