package impl

import (
	"auth-service/app/common/exceptions"
	"auth-service/app/entities"
	"auth-service/app/repositories"
	"context"
	"errors"
	"github.com/google/uuid"
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

func (u userRepositoryImpl) Create(usr repositories.User) {
	user := entities.User{
		Username:       usr.Username,
		Password:       usr.Password,
		Email:          usr.Email,
		FullName:       usr.FullName,
		PhoneNumber:    usr.PhoneNumber,
		ProfilePicture: usr.ProfilePicture,
		RoleGroupId:    usr.RoleGroupId,
		IsActive:       false,
	}
	err := u.DB.Create(&user).Error
	exceptions.PanicLogging(err)
}

func (u userRepositoryImpl) GetSingleUserById(id uuid.UUID) (entities.User, error) {
	user := entities.User{Uuid: id}
	res := u.DB.Find(&user).Error
	if res == nil {
		return entities.User{}, errors.New("user not found")
	}

	return user, nil
}

func (u userRepositoryImpl) DeleteSingle(id uuid.UUID) (string, error) {
	user := entities.User{Uuid: id}
	res := u.DB.Delete(&user)
	if res.RowsAffected == 0 {
		return "", errors.New("user not found")
	}

	return "User with id {id} deleted", nil
}

func (u userRepositoryImpl) DeleteAll() {
	err := u.DB.Where("1=1").Delete(&entities.User{}).Error
	exceptions.PanicLogging(err)
}
