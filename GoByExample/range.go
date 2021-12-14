package main

import "fmt"

func main() {

	//range 迭代各种数据结构
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
		fmt.Println("i: %s -- num:%s", i, num)
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Println("%s -> %s", k, v)
	}

	for i, c := range "go" {
		//i 起始字节位置，0，1之类; c 编码
		fmt.Println(i, c)
	}
	//0 103
	//1 111
}
