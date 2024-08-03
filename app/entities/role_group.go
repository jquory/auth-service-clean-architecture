package entities

import (
	"github.com/google/uuid"
	"time"
)

type RoleGroup struct {
	Uuid      uuid.UUID `gorm:"primaryKey;column:id;type:uuid;default:gen_random_uuid()" json:"id"`
	Name      string    `gorm:"column:name;type:varchar(255)" json:"roleGroupName"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (RoleGroup) TableName() string {
	return "role_group"
}
