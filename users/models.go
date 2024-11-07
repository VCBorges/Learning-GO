package users

import (
	"time"
	"github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID  `gorm:"column:id;type:varchar(50);primaryKey"`
	Email     string     `gorm:"column:email;type:varchar(100);uniqueIndex;not null"`
	FirstName string     `gorm:"column:first_name;type:varchar(50);not null"`
	Password  string     `gorm:"column:password;type:varchar(255);not null"`
	CreatedAt time.Time  `gorm:"column:created_at;type:timestamp;not null"`
	UpdateAt  *time.Time `gorm:"column:update_at;type:timestamp"`
}
