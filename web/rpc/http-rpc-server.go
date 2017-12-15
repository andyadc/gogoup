package main

/*
Go标准包中已经提供了对RPC的支持，而且支持三个级别的RPC：TCP、HTTP、JSONRPC。
但Go的RPC包是独一无二的RPC，它和传统的RPC系统不同，它只支持Go开发的服务器与客户端之间的交互，因为在内部，它们采用了Gob来编码。

Go RPC的函数只有符合下面的条件才能被远程访问，不然会被忽略，详细的要求如下：
	函数必须是导出的(首字母大写)
	必须有两个导出类型的参数，
	第一个参数是接收的参数，第二个参数是返回给客户端的参数，第二个参数必须是指针类型的
	函数还要有一个返回值error

举个例子，正确的RPC函数格式如下：
	func (t *T) MethodName(argType T1, replyType *T2) error

T、T1和T2类型必须能被encoding/gob包编解码。

任何的RPC都需要通过网络来传递数据，Go RPC可以利用HTTP和TCP来传递数据，利用HTTP的好处是可以直接复用net/http里面的一些函数。

*/

import (
	"errors"
	"fmt"
	"net/http"
	"net/rpc"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (arith *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (arith *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

/**
HTTP RPC
*/
func main() {
	arith := new(Arith)
	rpc.Register(arith) // 注册了一个Arith的RPC服务
	rpc.HandleHTTP()    // 把该服务注册到了HTTP协议上

	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		fmt.Println(err.Error())
	}

}
