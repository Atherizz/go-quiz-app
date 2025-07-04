package web

type UserAnswerInput struct {
	QuestionId     int `json:"question_id" binding:"required"`
	SelectedOption int `json:"selected_option" binding:"required"`
}

type SubmitQuizRequest struct {
	UserId  int               `json:"user_id"`
	QuizId  int               `json:"quiz_id"`
	Answers []UserAnswerInput `json:"answers" binding:"required,dive"`
}
