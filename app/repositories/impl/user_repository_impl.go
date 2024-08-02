package impl

import (
	"auth-service/app/entities"
	"auth-service/app/repositories"
	"context"
	"errors"
	"gorm.io/gorm"
)

func NewUserRepositoryImpl(DB *gorm.DB) repositories.UserRepository {
	return &userRepositoryImpl{}
}

type userRepositoryImpl struct {
	*gorm.DB
}

func (u userRepositoryImpl) Authentication(ctx context.Context, username string) (entities.User, error) {
	var userResult entities.User
	result := u.DB.WithContext(ctx).
		Joins("inner join role on role.group_id = users.role_group_id").
		Preload("UserRoles").
		Where("users.username = ? and users.is_active = ?", username, true).
		Find(&userResult)
	if result.RowsAffected == 0 {
		return entities.User{}, errors.New("user not found")
	}

	return userResult, nil
}

func (u userRepositoryImpl) Create(user repositories.User) {

}

func (u userRepositoryImpl) DeleteAll() {
	
}
