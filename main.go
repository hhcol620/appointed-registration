package main

import "fmt"

func main() {
	// 通过传入数据医院,
	params := fmt.Sprintf(`
		{
			"hostCode":%v, 

		}
`, "lll")
	fmt.Println(params)
}
