package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println(IsIP(`196.123.1.3`))
}

func IsIP(ip string) bool {
	if m, _ := regexp.MatchString(`((25[0-5]|2[0-4]\d|((1\d{2})|([1-9]?\d)))\.){3}(25[0-5]|2[0-4]\d|((1\d{2})|([1-9]?\d)))`, ip); m {
		return true
	}
	return false
}
