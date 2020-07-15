package entity

// MenuActionResource 菜单动作关联资源实体
type MenuActionResource struct {
	Model
	ActionID string `gorm:"column:action_id;size:36;index;default:'';not null;"` // 菜单动作ID
	Method   string `gorm:"column:method;size:100;default:'';not null;"`         // 资源请求方式(支持正则)
	Path     string `gorm:"column:path;size:100;default:'';not null;"`           // 资源请求路径（支持/:id匹配）
}

// TableName 表名
func (a MenuActionResource) TableName() string {
	return a.Model.TableName("menu_action_resource")
}
