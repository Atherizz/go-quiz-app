package web

type QuestionResponse struct {
	ID            int                    `json:"id"`
	QuizId        int                    `json:"quiz_id"`
	Quiz          QuizResponseMini       `json:"quiz"`
	QuestionText  string                 `json:"question_text"`
	AnswerOptions []AnswerOptionResponse `json:"answer_options"`
}
