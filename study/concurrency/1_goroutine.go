package main

import (
	"fmt"
	"runtime"
)

/*
goroutine是Go并行设计的核心。goroutine说到底其实就是协程，但是它比线程更小

goroutine是通过Go的runtime管理的一个线程管理器。goroutine通过go关键字实现了，其实就是一个普通的函数。
	go hello(a, b, c)


*/

func main() {
	go say("world")
	say("hello")
}

func say(s string) {
	for i := 0; i < 10; i++ {
		runtime.Gosched() // 让CPU把时间片让给别人,下次某个时候继续恢复执行该goroutine。
		fmt.Println(s)
	}
}
