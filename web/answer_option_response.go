package web

type AnswerOptionResponse struct {
	ID         int    `json:"id"`
	QuestionId int    `json:"question_id"`
	OptionText string `json:"option_text"`
	IsCorrect  bool   `json:"is_correct"`
}
