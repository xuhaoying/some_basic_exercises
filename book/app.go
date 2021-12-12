package main

import (
	//"some_basic_exercises/book/cat"
	"fmt"
	"math"
)

func main() {
	//cat.Work()
	var m, n int
	fmt.Print("请输入第一个操作数：")
	fmt.Scan((&m))
	fmt.Print("请输入第二个操作数：")
	fmt.Scan(&n)

	r1 := m + n
	r2 := m - n
	r3 := m * n
	r4 := m / n
	r5 := m % n
	r6 := math.Pow(float64(m), float64(n))

	fmt.Printf("%d + %d = %d\n", m, n, r1)
	fmt.Printf("%d - %d = %d\n", m, n, r2)
	fmt.Printf("%d * %d = %d\n", m, n, r3)
	fmt.Printf("%d / %d = %d\n", m, n, r4)
	fmt.Printf("%d %% %d = %d\n", m, n, r5)
	fmt.Printf("%.0f ** %d = %.2f\n", float64(m), int(float64(n)), r6)
	//panic: runtime error: integer divide by zero

}
