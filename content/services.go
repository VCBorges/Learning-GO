package content

import (
	"project_name/core"

	"gorm.io/gorm"
)

func CreateQuestion(
	dto *CreateQuestionDTO,
	db *gorm.DB,
) (*Question, error) {
	question := Question{
		BaseModel: core.NewBaseModel(),
		Text:      dto.Text,
		Answers:   make([]Answer, 0, len(dto.Answers)),
		Tags:      make([]Tag, 0, len(dto.Tags)),
	}

	for _, answerDTO := range dto.Answers {
		answer := Answer{
			BaseModel:  core.NewBaseModel(),
			Text:       answerDTO.Text,
			IsCorrect:  answerDTO.IsCorrect,
			QuestionId: question.Id,
		}
		question.Answers = append(question.Answers, answer)
	}

	for _, tagDTO := range dto.Tags {
		tag := Tag{
			BaseModel:  core.NewBaseModel(),
			Name: tagDTO.Name,
		}
		question.Tags = append(question.Tags, tag)
	}

	if err := db.Create(question).Error; err != nil {
		return nil, err
	}

	return &question, nil
}