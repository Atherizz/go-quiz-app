package web

type UserQuizResultResponse struct {
	ID             int              `json:"id"`
	UserId         int              `json:"user_id"`
	User           UserResponseMini `json:"user"`
	QuizId         int              `json:"quiz_id"`
	Quiz           QuizResponseMini `json:"quiz"`
	Score          float64              `json:"score"`
	TotalQuestions int              `json:"total_questions"`
	CorrectAnswers int              `json:"correct_answers"`
}
