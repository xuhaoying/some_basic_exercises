package main

import "fmt"

func main() {
	//map 关联数据类型（哈希、字典）
	//创建一个空map
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13
	fmt.Println(m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)
	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	//键不存在时直接取值返回的是零值
	fmt.Println("del:", m["k2"])
	fmt.Println("del:", m)

	_, prs := m["k2"] //第二个参数返回false，可用来判断键不存在
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

}
