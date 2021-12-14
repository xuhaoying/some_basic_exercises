package main

import "fmt"

func main() {
	//slice的类型仅由它所包含的元素决定，不像数组还需要元素的个数
	//要创建一个长度非零的空slice，需要使用内建的方法make
	s := make([]string, 3) //创建一个长度3的string类型的slice。初始化为零值
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f") //return a new slice
	fmt.Println("apd:", s)

	c := make([]string, len(s)) //the same length
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:5]
	fmt.Println("s[2:5]:", l)
	fmt.Println("s[:5]:", s[:5])
	fmt.Println("s[2:]:", s[2:])

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}
