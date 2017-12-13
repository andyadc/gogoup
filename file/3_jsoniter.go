package main

/*

http://jsoniter.com/index.cn.html
*/
import (
	"fmt"
	"github.com/json-iterator/go"
)

func main() {

	str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	var data interface{}

	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	json.Unmarshal([]byte(str), &data)

	fmt.Println(data)

	val := []byte(`{"ID":1,"Name":"Reds","Colors":["Crimson","Red","Ruby","Maroon"]}`)
	res := jsoniter.Get(val, "Colors", 1).ToString()
	fmt.Println(res)

	m := map[string]interface{}{
		"3": 3,
		"1": 1,
		"2": 2,
	}
	b, _ := json.Marshal(m)
	fmt.Println(string(b))
}
