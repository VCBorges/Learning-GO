package content

import (
	"errors"
	"project_name/core"

	"github.com/google/uuid"
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
			BaseModel: core.NewBaseModel(),
			Name:      tagDTO.Name,
		}
		question.Tags = append(question.Tags, tag)
	}

	if err := db.Create(question).Error; err != nil {
		return nil, err
	}

	return &question, nil
}

func FindQuestionById(
	questionId uuid.UUID,
	db *gorm.DB,
) (*Question, error) {
	var question Question
	err := db.First(&question, questionId).Error
	return &question, err
}

func FindAnswerById(
	answerId uuid.UUID,
	db *gorm.DB,
) (*Answer, error) {
	var answer Answer
	err := db.First(&answer, answerId).Error
	return &answer, err
}

func ChooseQuestionAnswer(
	answerId uuid.UUID,
	questionId uuid.UUID,
	db *gorm.DB,
) error {
	var answer Answer

	err := db.Where("id = ?, question_id = ?", answerId, questionId).First(&answer).Error
	if err != nil {
		return errors.New("answer does not exists")
	}

	return err
}
