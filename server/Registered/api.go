package registered

import "fmt"

// 筛选出需要查询的

// 传入医院的所有的数据进行查询
func Check(hasCode, firstDeptCode, secondDeptCode string) {
	// 通过传入数据医院,
	params := fmt.Sprintf(`
			{
				"hostCode":%v, 
				"firstDeptCode":%v,
				"secondDeptCode":%v,
				"week":1
			}
	`, hasCode, firstDeptCode, secondDeptCode)

}
