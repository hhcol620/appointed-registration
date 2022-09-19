package listdepartment

// 声明存储数据结构体
// 大类
type KeyDepartment struct {
	Name string
	Code string
}

// 小类
type ValueDepartment struct {
	Code      string
	Dept1Code string
	HotDept   bool
	Name      string
}
