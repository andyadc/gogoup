package main

/*
Go的import还支持如下两种方式来加载自己写的模块:
1.相对路径
	import “./model” //当前文件同一目录的model目录，但是不建议这种方式来import

2.绝对路径
	import “shorturl/model” //加载$GOPATH/src/shorturl/model模块

一些特殊的import
1.点操作
	import(
		 . "fmt"
	)
	这个点操作的含义就是这个包导入之后在你调用这个包的函数时，你可以省略前缀的包名，
	也就是前面你调用的fmt.Println("hello world")可以省略的写成Println("hello world")

2.别名操作
	import (
		f "fmt"
	)
	把包命名成另一个我们用起来容易记忆的名字,
	别名操作的话调用包函数时前缀变成了我们的前缀，即f.Println("hello world")

3._ 操作
	import (
	    "database/sql"
	    _ "github.com/ziutek/mymysql/godrv"
	)
	_ 操作其实是引入该包，而不直接使用包里面的函数，而是调用了该包里面的init函数。

*/
import (
	f "fmt"
)

func main() {
	f.Println("import")
}
