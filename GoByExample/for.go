package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i += 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for { //一直执行，知道break或return跳出
		fmt.Println("loop")
		break
	}
}
