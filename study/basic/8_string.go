package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(strings.Contains(`andyadc`, `adc`))

	s := []string{"a", "b", "c"}
	fmt.Println(strings.Join(s, `, `))

	fmt.Println(strings.Repeat(`a`, 10))

	str := make([]byte, 0, 100)
	str = strconv.AppendInt(str, 4567, 10)
	str = strconv.AppendBool(str, false)

	fmt.Println(string(str))

	name := "安"
	for k, v := range []byte(name) {
		fmt.Println(k, v)
	}
}
