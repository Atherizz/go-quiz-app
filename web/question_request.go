package web

type QuestionRequest struct {
	ID           int    `json:"id"`
	QuizId       int    `json:"quiz_id" binding:"required"`
	QuestionText string `json:"question_text" binding:"required,min=1,max=300"`
}