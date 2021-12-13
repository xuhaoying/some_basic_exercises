package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Println("write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend.")
	default:
		fmt.Println("it's a weekday.")
	}

	t := time.Now()
	switch {
	case t.Hour() > 12:
		fmt.Println("It's after noon") //执行完就退出了
	default:
		fmt.Println("It's before noon.")

	}

}
