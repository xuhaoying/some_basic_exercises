package main

import "fmt"

func intSeq() func() int {
	//闭包
	//返回一个在函数内定义的匿名函数，使用变量i
	i := 0
	return func() int {
		i += 1
		return i
	}

}

func main() {
	//将返回值赋给nextInt, 返回值是一个函数
	//函数包含i，每次调用都会更新i的值
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	//新的函数
	newInt2 := intSeq()
	fmt.Println(newInt2())
}
