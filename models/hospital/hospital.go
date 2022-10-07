package hospital

type Hospital struct {
	ID           uint   `gorm:"type:int;primaryKey"`
	Code         string `gorm:"column:code;type:varchar(30);not null;comment '医院代码'"`
	LevelText    string `gorm:"column:levelText;type:varchar(40);not null;comment '医院等级'"`
	Maintain     bool   `gorm:"column:maintain;type:varchar(40);not null;comment 'maintain'"`
	Name         string `gorm:"column:name;type:varchar(100);not null;comment '医院名称'"`
	OpenTimeText string `gorm:"column:openTimeText;type:varchar(40);not null;comment '开放时间'"`
	Picture      string `gorm:"column:picture;type:varchar(150);comment '医院图标'"`
}
