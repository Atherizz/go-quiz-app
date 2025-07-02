package handler

type AppHandler struct {
	Auth   *AuthHandler
	Subject *SubjectHandler
	Quiz *QuizHandler
	Question *QuestionHandler
	AnswerOption *AnswerOptionHandler
	UserAnswer *UserQuizResultHandler
	UserQuizResult *UserQuizResultHandler
}