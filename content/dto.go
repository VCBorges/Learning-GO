package content

type CreateQuestionDTO struct {
	Text    string
	Answers []CreateAnswerDTO
	Tags    []CreateTagDTO
}

type CreateAnswerDTO struct {
	Text      string
	IsCorrect bool
}

type CreateTagDTO struct {
	Name string
}
