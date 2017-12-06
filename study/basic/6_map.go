package main

// ********** map **********

/*
map也就是Python中字典的概念，它的格式为map[keyType]valueType
	// 声明一个key是字符串，值为int的字典,这种方式的声明需要在使用之前使用make初始化
	var numbers map[string]int
	numbers = make(map[string]int)
	numbers["a"] = 1
	numbers["b"] = 10
	numbers["c"] = 123
	fmt.Println(numbers)

	// 另一种map的声明方式
	names := make(map[string]int)
	names["jack"] = 23
	names["andy"] = 223
	names["green"] = 123


使用map过程中需要注意的几点：
	map是无序的，每次打印出来的map都会不一样，它不能通过index获取，而必须通过key获取
	map的长度是不固定的，也就是和slice一样，也是一种引用类型
	内置的len函数同样适用于map，返回map拥有的key的数量
	map的值可以很方便的修改，通过numbers["one"]=11可以很容易的把key为one的字典值改为11
	map和其他基本型别不同，它不是thread-safe，在多个go-routine存取时，必须使用mutex lock机制

map的初始化可以通过key:val的方式初始化值，同时map内置有判断是否存在key的方式
通过delete删除map的元素：
	// 初始化一个字典
	rating := map[string]int{`Java`: 1, `C`: 2, `Python`: 3, `Go`: 4, "C#":5}
	// map有两个返回值，第二个返回值，如果不存在key，那么ok为false，如果存在ok为true
	csharpRating, ok := rating["C#"]

	fmt.Println(rating)
	fmt.Println(ok)
	fmt.Println(csharpRating)

	if ok {
		// 删除key为C#的元素
		delete(rating, "C#")
	}
	fmt.Println(rating)

map也是一种引用类型，如果两个map同时指向一个底层，那么一个改变，另一个也相应的改变：
	m := make(map[string]string)
	m["hello"] = "Bonjour"
	n := m
	n["hello"] = `v`


*/
