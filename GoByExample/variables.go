package main

import "fmt"

func main() {
	var a string = "initial"
	fmt.Println(a)

	var b, c int = 1, 2 //声明多个变量
	fmt.Println(b, c)

	var d = true //自动推断已经初始化的变量类型
	fmt.Println(d)

	var e int //没有给出初始值，初始化为零值
	fmt.Println(e)

	f := "short" //声明并初始化变量
	fmt.Println(f)
}
