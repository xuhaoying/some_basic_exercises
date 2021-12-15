package main

import "fmt"

func plus(a int, b int) int {
	//求和，接受两个int，并返回int
	return a + b

}

func vals(a int, b int) (int, int) {
	//返回两个int
	return a + b, a - b

}

func variadic(nums ...int) {
	//可变参数，使用任意数目的int
	fmt.Println(nums, "")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	a, b := vals(3, 7)
	fmt.Println(a)
	fmt.Println(b)
	//go里面变量如果不使用会报错，用` _ `接收多余的参数
	_, c := vals(4, 6)
	fmt.Println(c)

	variadic(1, 2)
	variadic(1, 2, 3, 4)

	nums := []int{1, 2, 3, 4, 6, 9}
	//多个值作为变参 func(slice...)
	variadic(nums...)
}
