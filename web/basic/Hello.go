package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
	http.HandleFunc("/", sayHello) // 设置访问的路由
	http.HandleFunc("/login", login)

	log.Println("Starting server...")
	err := http.ListenAndServe(":8888", nil) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

func login(w http.ResponseWriter, r *http.Request) {
	log.Printf("method", r.Method) // 获取请求的方法
	if r.Method == `GET` {
		t, err := template.ParseFiles("login.gtpl")
		if err != nil {
			log.Fatalf("Parse error: ", err)
		}
		log.Println(t.Execute(w, nil))
	} else {
		r.ParseForm()
		log.Println("username: ", r.Form["username"])
		log.Println("password: ", r.Form["password"])
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

	// 设置cookie
	expiration := time.Now()
	expiration = expiration.AddDate(0, 0, 1)
	cookie := http.Cookie{Name: `go-cookie-hello`, Value: `Hello Go`, Expires: expiration, HttpOnly: true}
	http.SetCookie(w, &cookie)

	fmt.Fprint(w, "Hello Go") //这个写入到w的是输出到客户端的
}
