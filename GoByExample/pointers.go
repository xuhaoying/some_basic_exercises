package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	//*int 指针，地址对应的值变为0
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i) //值传参，i没有被改变
	fmt.Println("zeroval:", i)

	zeroptr(&i) //传入i的内存地址，更改地址对应的值，i值变化
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)
}
