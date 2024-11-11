package core

import (
	"time"

	"github.com/google/uuid"
)

type BaseModel struct {
	Id        uuid.UUID  `gorm:"column:id;type:varchar(50);primaryKey"`
	CreatedAt time.Time  `gorm:"column:created_at;type:timestamp;not null"`
	UpdateAt  *time.Time `gorm:"column:update_at;type:timestamp"`
}

func NewBaseModel() BaseModel {
	return BaseModel{
		Id:        uuid.New(),
		CreatedAt: time.Now(),
	}
}
