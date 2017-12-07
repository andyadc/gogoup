package main

import "fmt"

/*
Go里面最强大的一个控制逻辑就是for，它即可以用来循环读取数据，又可以当作while来控制逻辑，还能迭代操作。
语法如下：
	for expression1; expression2; expression3 {
		//...
	}

	sum := 0
	for i := 0; i < 100; i ++ {
		sum += i
	}
	fmt.Println("sum: ", sum)

	// 忽略expression1和expression3：
	total := 1
	for ; total < 1000;  {
		total += total
	}
	fmt.Println("sum: ", total)

	// 相当于while
	num := 1
	for num < 1000  {
		num += num
	}
	fmt.Println("sum: ", num)

	// 循环里面有两个关键操作break和continue	,break操作是跳出当前循环，continue是跳过本次循环。
	for i := 10; i > 0; i-- {
		if i == 5 {
			// break
			continue
		}
		fmt.Println(i)
	}

	// for配合range可以用于读取slice和map的数据：
	params := map[string]int{"A": 90, "B": 80, "C": 70}
	for k, v := range params {
		fmt.Println(k, v)
	}

	// 可以使用_来丢弃不需要的返回值
	for _, v := range params {
		fmt.Println(v)
	}

*/

func main() {
	// calculate()

	params := map[string]int{"A": 90, "B": 80, "C": 70}
	for k, v := range params {
		fmt.Println(k, v)
	}

}

func calculate() {
	sum := 0
	for i := 0; i < 100; i++ {
		sum += i
	}
	fmt.Println("sum: ", sum)

	// 忽略expression1和expression3：
	total := 1
	for total < 1000 {
		total += total
	}
	fmt.Println("sum: ", total)

	// 相当于while
	num := 1
	for num < 1000 {
		num += num
	}
	fmt.Println("sum: ", num)
}
