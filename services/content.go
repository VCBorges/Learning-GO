package services

import (
	"project_name/models"

	"gorm.io/gorm"
)

type CreateQuestionDTO struct {
	Text    string
	Answers []CreateAnswerDTO
}

type CreateAnswerDTO struct {
	Text      string
	IsCorrect bool
}

func CreateQuestion(
	dto CreateQuestionDTO, 
	db *gorm.DB,
) (*models.Question, error) {
	question := models.Question{
		BaseModel: models.NewBaseModel(),
		Text:      dto.Text,
		Answers:   make([]models.Answer, 0, len(dto.Answers)),
	}

	for _, answerDTO := range dto.Answers {
		answer := models.Answer{
			BaseModel: models.NewBaseModel(),
			Text:      answerDTO.Text,
			IsCorrect: answerDTO.IsCorrect,
			QuestionId: question.Id,
		}
		question.Answers = append(question.Answers, answer)
	}

	if err := db.Create(question).Error; err != nil {
		return nil, err
	}

	return &question, nil
}
