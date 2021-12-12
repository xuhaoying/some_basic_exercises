package main // 如果不是main,要先build才能直接运行

import (
	"fmt"
)

func main() {
	var stockcode = 123
	var enddate = "2020-12-31"
	var url = "Code=%d&endDate=%s"
	var target_url = fmt.Sprintf(url, stockcode, enddate)
	fmt.Println(target_url)
}
