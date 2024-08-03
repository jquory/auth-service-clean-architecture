package entities

import (
	"github.com/google/uuid"
	"time"
)

type Role struct {
	Uuid      uuid.UUID `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"column:role_name;type:varchar(100)" json:"roleName"`
	GroupId   int64     `gorm:"column:group_id;type:int" json:"groupId"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Role) TableName() string {
	return "role"
}
