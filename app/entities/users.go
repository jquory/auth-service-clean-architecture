package entities

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id             int32     `gorm:"autoIncrement;column:seq;type:int"`
	Uuid           uuid.UUID `gorm:"primaryKey;column:id;type:uuid;default:gen_random_uuid()" json:"id"`
	Username       string    `gorm:"column:username;type:varchar(100)" json:"username"`
	Password       string    `gorm:"column:password;type:varchar(255)"`
	FullName       *string   `gorm:"column:full_name;type:varchar(255)" json:"fullName"`
	Email          *string   `gorm:"column:email;type:varchar(200);unique" json:"email"` // using pointer string to define nullable field
	PhoneNumber    *string   `gorm:"column:phone_number;type:varchar(100)" json:"phoneNumber"`
	ProfilePicture *string   `gorm:"column:profile_picture;type:varchar(200)" json:"profilePicture"`
	RoleGroupId    uuid.UUID `gorm:"column:role_group_id;type:int" json:"roleGroupId"`
	IsActive       bool      `gorm:"column:is_active;type:boolean" json:"isActive"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (User) TableName() string {
	return "users"
}
