package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", sayHello) // 设置访问的路由

	log.Println("Starting server...")
	err := http.ListenAndServe(":8888", nil) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

func sayHello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // 解析参数，默认是不会解析的
	log.Println(r.Form)
	log.Println("path", r.URL.Path)
	log.Println("scheme", r.URL.Scheme)

	for k, v := range r.Form {
		log.Println("key: ", k, ", val: ", strings.Join(v, ``))
	}

	fmt.Fprint(w, "Hello Go") //这个写入到w的是输出到客户端的
}
