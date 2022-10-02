package address

/**
* 代码描述: 行政区的代号
* 作者:小大白兔
* 创建时间:2022/10/01 20:30:13
 */
type AllAddress struct {
	ID           uint   `gorm:"type:int;primaryKey"`
	AddressCode  string `gorm:"column:addressCode;type:varchar(6);not null;comment'行政区代码'"`
	AddressValue string `gorm:"column:addressValue;type:varchar(20);not null;comment'行政区名称'"`
}

/**
* 代码描述: 医院等级
* 作者:小大白兔
* 创建时间:2022/10/01 20:44:56
 */
type AllGrade struct {
	ID         uint   `gorm:"type:int;primaryKey"`
	GradeCode  uint   `gorm:"column:gradeCode;type:int;not null;comment'医院级别'"`
	GradeValue string `gorm:"column:gradeValue;type:varchar(20);not null;comment '级别名字'"`
}
