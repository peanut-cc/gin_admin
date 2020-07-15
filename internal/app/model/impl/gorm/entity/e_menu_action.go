package entity

// MenuAction 菜单动作实体
type MenuAction struct {
	Model
	MenuID string `gorm:"column:menu_id;size:36;index;default:'';not null;"` // 菜单ID
	Code   string `gorm:"column:code;size:100;default:'';not null;"`         // 动作编号
	Name   string `gorm:"column:name;size:100;default:'';not null;"`         // 动作名称
}

// TableName 表名
func (a MenuAction) TableName() string {
	return a.Model.TableName("menu_action")
}
