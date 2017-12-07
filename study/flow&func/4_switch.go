package main

import "fmt"

/*
语法如下:
	switch sExpr {
	case expr1:
		some instructions
	case expr2:
		some other instructions
	case expr3:
		some other instructions
	default:
		other code
	}

***Go里面switch默认相当于每个case最后带有break，匹配成功后不会自动向下执行其他case，而是跳出整个switch,
但是可以使用fallthrough强制执行后面的case代码。

fallthrough强制执行后面的case代码是不关心case条件是否匹配

*/

func main() {
	i := 1
	switch i {
	case 1:
		fmt.Println("i is equal to 1")
		fallthrough
	case 2:
		fmt.Println("i is equal to 2")
		fallthrough
	case 3, 4, 5: // 可以把很多值聚合在了一个case里面
		fmt.Println("i is equal to 3, 4 or 5")
		fallthrough
	default:
		fmt.Println("All I know is that i is an integer")
	}
}
