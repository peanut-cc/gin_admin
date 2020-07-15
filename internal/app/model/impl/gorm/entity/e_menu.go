package entity

// Menu 菜单实体
type Menu struct {
	Model
	Name       string  `gorm:"column:name;size:50;index;default:'';not null;"` // 菜单名称
	Sequence   int     `gorm:"column:sequence;index;default:0;not null;"`      // 排序值
	Icon       *string `gorm:"column:icon;size:255;"`                          // 菜单图标
	Router     *string `gorm:"column:router;size:255;"`                        // 访问路由
	ParentID   *string `gorm:"column:parent_id;size:36;index;"`                // 父级内码
	ParentPath *string `gorm:"column:parent_path;size:518;index;"`             // 父级路径
	ShowStatus int     `gorm:"column:show_status;index;default:0;not null;"`   // 状态(1:显示 2:隐藏)
	Status     int     `gorm:"column:status;index;default:0;not null;"`        // 状态(1:启用 2:禁用)
	Memo       *string `gorm:"column:memo;size:1024;"`                         // 备注
	Creator    string  `gorm:"column:creator;size:36;"`                        // 创建人
}

// TableName 表名
func (a Menu) TableName() string {
	return a.Model.TableName("menu")
}
