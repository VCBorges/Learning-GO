package models

import "github.com/google/uuid"

type Question struct {
	BaseModel
	Text    string   `gorm:"column:text;type:varchar(100);not null"`
	Answers []Answer `gorm:"foreignKey:QuestionId;references:Id;constraint:OnDelete:CASCADE"`
}

type Answer struct {
	BaseModel
	Text       string    `gorm:"column:text;type:varchar(100);not null"`
	IsCorrect  bool      `gorm:"column:is_correct;type:bool;not null;default:false"`
	QuestionId uuid.UUID `gorm:"column:question_id;type:varchar(100);not null"`
	Question   Question  `gorm:"foreignKey:QuestionId;references:Id"`
}
