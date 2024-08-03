package repositories

import (
	"auth-service/app/entities"
	"context"
	"github.com/google/uuid"
)

type User struct {
	Id             int32
	Uuid           uuid.UUID
	Username       string
	Password       string
	FullName       *string
	Email          *string
	PhoneNumber    *string
	ProfilePicture *string
	RoleGroupId    uuid.UUID
}

type UserRepository interface {
	Authentication(ctx context.Context, username string) (entities.User, error)
	Create(user User)
	GetSingleUserById(id uuid.UUID) (entities.User, error)
	DeleteSingle(id uuid.UUID) (string, error)
	DeleteAll()
}
