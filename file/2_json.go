package main

import (
	"encoding/json"
	"fmt"
)

/**

{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}

struct tag
`json:"servers`

Go类型和JSON类型的对应关系如下：
	bool 代表 JSON booleans,
	float64 代表 JSON numbers,
	string 代表 JSON strings,
	nil 代表 JSON null.

*/

type ServerSlice struct {
	Servers []Server `json:"servers"`
}

type Server struct {
	ID         int    `json:"-"`
	Address    string `json:"adress,omitempty"`
	ServerName string `json:"serverName"`
	ServerIP   string `json:"serverIP"`
}

func main() {
	//readJson()
	//parseJson()
	writeJson()
}

/**
针对JSON的输出，我们在定义struct tag的时候需要注意的几点是:
	字段的tag是"-"，那么这个字段不会输出到JSON
	tag中带有自定义名称，那么这个自定义名称会出现在JSON的字段名中，例如上面例子中serverName
	tag中如果带有"omitempty"选项，那么如果该字段值为空，就不会输出到JSON串中
	如果字段类型是bool, string, int, int64等，而tag中带有",string"选项，那么这个字段在输出到JSON的时候会把该字段对应的值转换成JSON字符串

Marshal函数只有在转换成功的时候才会返回数据，在转换的过程中我们需要注意几点：
	JSON对象只支持string作为key，所以要编码一个map，那么必须是map[string]T这种类型(T是Go语言中任意的类型)
	Channel, complex和function是不能被编码成JSON的
	嵌套的数据是不能编码的，不然会让JSON编码进入死循环
	指针在编码的时候会输出指针指向的内容，而空指针会输出null
*/
func writeJson() {
	var ss ServerSlice
	ss.Servers = append(ss.Servers, Server{ID: 1, ServerName: `a`, ServerIP: `123`})
	ss.Servers = append(ss.Servers, Server{ID: 2, ServerName: `b`, ServerIP: `321`, Address: `shanghai`})

	b, err := json.Marshal(ss)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(string(b))
}

func parseJson() {
	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	var f interface{}
	json.Unmarshal(b, &f)
	fmt.Println(f)

	// 断言
	m := f.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int", vv)
		case float64:
			fmt.Println(k, "is float64", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}
}

func readJson() {
	var ss ServerSlice
	str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	json.Unmarshal([]byte(str), &ss)
	fmt.Println(ss)
	fmt.Println(len(ss.Servers))
}
