package web

type QuizRequest struct {
	ID int `json:"id" binding:"required"`
	SubjectId int `json:"subject_id" binding:"required"`
	Description string `json:"description" binding:"required, min=1, max=200"`
	Title       string `json:"title" binding:"required, min=1, max=100"`
}