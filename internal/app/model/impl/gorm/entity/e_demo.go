package entity

// Demo demo实体
type Demo struct {
	Model
	Code    string  `gorm:"column:code;size:50;index;default:'';not null;"`  // 编号
	Name    string  `gorm:"column:name;size:100;index;default:'';not null;"` // 名称
	Memo    *string `gorm:"column:memo;size:200;"`                           // 备注
	Status  int     `gorm:"column:status;index;default:0;not null;"`         // 状态(1:启用 2:停用)
	Creator string  `gorm:"column:creator;size:36;"`                         // 创建者
}

// TableName 表名
func (a Demo) TableName() string {
	return a.Model.TableName("demo")
}
