package main

import (
	"errors"
	"fmt"
)

//**************内置基础类型****************
/*
Boolean
在Go中，布尔值的类型为bool，值是true或false，默认为false。

	var isActive bool                   // 全局变量声明
	var enabled, disabled = true, false // 忽略类型的声明
	func testBool() {
		var available bool // 一般声明
		valid := true      // 简短声明
		available = true   // 赋值操作
	}

数值类型
整数类型有无符号和带符号两种。Go同时支持int和uint，这两种类型的长度相同，但具体长度取决于不同编译器的实现。
Go里面也有直接定义好位数的类型：rune, int8, int16, int32, int64和byte, uint8, uint16, uint32, uint64。其中rune是int32的别称，byte是uint8的别称。
*这些类型的变量之间不允许互相赋值或操作，不然会在编译时引起编译器报错。
*尽管int的长度是32 bit, 但int 与 int32并不可以互用
浮点数的类型有float32和float64两种（没有float类型），默认是float64。

Go还支持复数。它的默认类型是complex128（64位实数+64位虚数）。
如果需要小一些的，也有complex64(32位实数+32位虚数)。
复数的形式为RE + IMi，其中RE是实数部分，IM是虚数部分，而最后的i是虚数单位。
	var c complex128 = 4 + 5i

字符串
Go中的字符串都是采用UTF-8字符集编码。
字符串是用一对双引号（""）或反引号（` `）括起来定义，它的类型是string。
	var frenchHello string // 声明变量为字符串的一般方法
	var emptyString  string = `` // 声明了一个字符串变量，初始化为空字符串

	func test()  {
		no, yes, maybe := "no", "yes", `maybe` // 简短声明，同时声明多个变量
		china := "china" // 简短声明
		frenchHello = "bonjour" // 常规赋值
	}

在Go中字符串是不可变的. 如下面的代码编译时会报错：cannot assign to s[0]
	var s string = "hello"
	s[0] = 'a'

如果真的想要修改怎么办呢？下面的代码可以实现：
	s := `hello`
	b := []byte(s) // 将字符串 s 转换为 []byte 类型
	b[1] = 'p'
	s2 := string(b) // 再转换回 string 类型
	fmt.Println(s2)

Go中可以使用+操作符来连接两个字符串：
	s := "hello"
	m := " world"
	c := s + m
	fmt.Println(c)

如果要声明一个多行的字符串怎么办？可以通过`来声明：
	m := `hello
			world`
	fmt.Println(m)
` 括起的字符串为Raw字符串，即字符串在代码中的形式就是打印时的形式，它没有字符转义，换行也将原样输出。

错误类型
Go内置有一个error类型，专门用来处理错误信息，Go的package里面还专门有一个包errors来处理错误：
	err := errors.New("emit macho dwarf: elf header corrupted")
	if err != nil {
		fmt.Println(err)
	}


*/

func main() {
	err := errors.New("emit macho dwarf: elf header corrupted")
	if err != nil {
		fmt.Println(err)
	}
}
