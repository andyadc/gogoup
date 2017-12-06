package main

// ********** make / new *********

/*
make用于内建类型（map、slice 和channel）的内存分配。
new用于各种类型的内存分配。

NEW
内建函数new本质上说跟其它语言中的同名函数功能一样：new(T)分配了零值填充的T类型的内存空间，并且返回其地址，即一个*T类型的值。
用Go的术语说，它返回了一个指针，指向新分配的类型T的零值。有一点非常重要：
*new返回指针。

MAKE
内建函数make(T, args)与new(T)有着不同的功能，make只能创建slice、map和channel，并且返回一个有初始值(非零)的T类型，而不是*T。
本质来讲，导致这三个类型有所不同的原因是指向数据结构的引用在使用前必须被初始化。例如，一个slice，是一个包含指向数据（内部array）的指针、长度和容量的三项描述符；在这些项目被初始化之前，slice为nil。
对于slice、map和channel来说，make初始化了内部的数据结构，填充适当的值。
*make返回初始化后的（非零）值。


零值
关于“零值”，所指并非是空值，而是一种“变量未填充前”的默认值，通常为0。 此处罗列 部分类型 的 “零值”
int     0
int8    0
int32   0
int64   0
uint    0x0
rune    0 // rune的实际类型是 int32
byte    0x0 // byte的实际类型是 uint8
float32 0 //长度为 4 byte
float64 0 //长度为 8 byte
bool    false
string  ""


*/

func main() {
}
