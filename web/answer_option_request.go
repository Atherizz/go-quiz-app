package web

type AnswerOptionRequest struct {
	ID           int    `json:"id" binding:"required"`
	QuestionId   int    `json:"question_id" binding:"required"`
	OptionText   string `json:"option_text" binding:"required,min=1,max=300"`
	OptionNumber int    `json:"option_number" binding:"required,oneof=1 2 3 4"`
	IsCorrect    bool   `json:"is_correct" binding:"required"`
}
