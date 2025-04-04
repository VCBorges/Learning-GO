package content

import (
	"project_name/core"

	"github.com/google/uuid"
)

type Question struct {
	core.BaseModel
	Text    string   `gorm:"column:text;type:varchar(100);not null"`
	Answers []Answer `gorm:"foreignKey:QuestionId;references:Id;constraint:OnDelete:CASCADE"`
	Tags    []Tag    `gorm:"foreignKey:QuestionId;references:Id;constraint:OnDelete:CASCADE"`
}

type Answer struct {
	core.BaseModel
	Text       string    `gorm:"column:text;type:varchar(100);not null"`
	IsCorrect  bool      `gorm:"column:is_correct;type:bool;not null;default:false"`
	QuestionId uuid.UUID `gorm:"column:question_id;type:varchar(100);not null"`
	Question   Question  `gorm:"foreignKey:QuestionId;references:Id"`
}

type Tag struct {
	core.BaseModel
	Name       string    `gorm:"column:name;type:varchar(100);not null"`
	QuestionId uuid.UUID `gorm:"column:question_id;type:varchar(100);not null"`
	Question   Question  `gorm:"foreignKey:QuestionId;references:Id"`
}

