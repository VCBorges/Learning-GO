package users

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateUser(
	db *gorm.DB,
	data *UserCreateInput,
) (*User, error) {
	_, err := GetUserByEmail(data.Email, db)
	if err == nil {
		return &User{}, errors.New("email in use")
	}

	if data.Email == "" {
		return &User{}, errors.New("user email cannot be empty")
	}
	user := User{
		Id:        uuid.New(),
		Email:     data.Email,
		FirstName: data.FirstName,
		Password:  data.Password,
		CreatedAt: time.Now(),
	}
	return &user, db.Create(user).Error
}

func GetUserByEmail(
	email string,
	db *gorm.DB,
) (*User, error) {
	var user User

	err := db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserById(
	id uuid.UUID,
	db *gorm.DB,
) (*User, error) {
	var user User

	err := db.Where("id = ?", id.String()).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
