package main

import "fmt"

/*
channels
goroutine运行在相同的地址空间，因此访问共享内存必须做好同步。那么goroutine之间如何进行数据的通信呢，
Go提供了一个很好的通信机制channel。channel可以与Unix shell 中的双向管道做类比：可以通过它发送或者接收值。
这些值只能是特定的类型：channel类型。
定义一个channel时，也需要定义发送到channel的值的类型。注意，必须使用make 创建channel：
	ci := make(chan int)
	cs := make(chan string)
	cf := make(chan interface{})

channel通过操作符<-来接收和发送数据
	ch <- v    // 发送v到channel ch.
	v := <- ch  // 从ch中接收数据，并赋值给v

默认情况下，channel接收和发送数据都是阻塞的，除非另一端已经准备好，
这样就使得Goroutines同步变的更加的简单，而不需要显式的lock。

所谓阻塞，也就是如果读取（value := <-ch）它将会被阻塞，直到有数据接收。
其次，任何发送（ch<-5）将会被阻塞，直到数据被读出。

Buffered Channels
上面我们介绍了默认的非缓存类型的channel，不过Go也允许指定channel的缓冲大小，
很简单，就是channel可以存储多少元素。ch:= make(chan bool, 4)，创建了可以存储4个元素的bool 型channel。
在这个channel 中，前4个元素可以无阻塞的写入。当写入第5个元素时，代码将会阻塞，
直到其他goroutine从channel 中读取一些元素，腾出空间。
	ch := make(chan type, value)

当 value = 0 时，channel 是无缓冲阻塞读写的，
当value > 0 时，channel 有缓冲、是非阻塞的，直到写满 value 个元素才阻塞写入。

*/

func bufferedChannels() {
	c := make(chan int, 1) // 修改2为1就报错，修改2为3可以正常运行
	c <- 1
	c <- 2
	fmt.Println(<-c)
	fmt.Println(<-c)

	/*
		修改为1报如下错误:
		fatal error: all goroutines are asleep - deadlock!
	*/
}

func main() {
	a := []int{19, 2, 3, 20, 4, 5, 9, 12}
	c := make(chan int)
	d := make(chan int)

	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], d)

	x, y := <-c, <-d // receive from c
	fmt.Println(x, y, x+y)

	bufferedChannels()
}

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum
}
