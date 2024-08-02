package entities

import "time"

type Menu struct {
	Id           int64  `gorm:"autoIncrement;primaryKey"`
	MenuId       string `gorm:"column:menu_id;type:varchar(100);" json:"id"`
	Name         string `gorm:"column:name;type:varchar(200);" json:"menuName"`
	Url          string `gorm:"column:url;type:varchar(500);" json:"menuUrl"`
	Icon         string `gorm:"type:varchar(100);column:icon" json:"menuIcon"`
	ParentMenuId int64  `gorm:"column:parent_menu_id;type:int;" json:"menuParentId"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (Menu) TableName() string {
	return "menu"
}
