package users

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateUser(
	db *gorm.DB,
	data *UserCreateInput,
) (User, error) {
	user := User{
		Id:        uuid.New(),
		Email:     data.Email,
		FirstName: data.FirstName,
		Password:  data.Password,
		CreatedAt: time.Now(),
	}
	return user, db.Create(user).Error
}
