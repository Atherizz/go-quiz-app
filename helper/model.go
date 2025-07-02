package helper

import (
	"google-oauth/model"
	"google-oauth/web"
)

func ToUserResponse(user model.User) web.UserResponse {
	return web.UserResponse{
		Id:       user.ID,
		GoogleId: user.GoogleId,
		Name:     user.Name,
		Email:    user.Email,
		Picture:  user.Picture,
		Provider: user.Provider,
		Role:     user.Role,
	}
}

func ToSubjectResponse(subject model.Subject) web.SubjectResponse {
	return web.SubjectResponse{
		Id:          int(subject.Model.ID),
		SubjectName: subject.SubjectName,
	}
}

func ToQuizResponse(quiz model.Quiz) web.QuizResponse {
	return web.QuizResponse{
		ID:          int(quiz.ID),
		SubjectId:   quiz.SubjectId,
		Subject:     ToSubjectResponse(quiz.Subject),
		Description: quiz.Description,
		Title:       quiz.Title,
		Questions:   ToQuestionResponseSlice(quiz.Questions),
	}

}

func ToQuestionResponseSlice(questions []model.Question) []web.QuestionResponse {
	var result []web.QuestionResponse
	for _, q := range questions {
		result = append(result, ToQuestionResponse(q))
	}
	return result
}

func ToAnswerOptionResponseSlice(answers []model.AnswerOption) []web.AnswerOptionResponse {
	var result []web.AnswerOptionResponse
	for _, a := range answers {
		result = append(result, ToAnswerOptionResponse(a))
	}
	return result
}

func ToQuestionResponse(question model.Question) web.QuestionResponse {
	return web.QuestionResponse{
		ID:            int(question.ID),
		QuizId:        question.QuizId,
		Quiz:          web.QuizResponseMini{ID: int(question.Quiz.ID), Title: question.Quiz.Title},
		QuestionText:  question.QuestionText,
		AnswerOptions: ToAnswerOptionResponseSlice(question.AnswerOptions),
	}
}

func ToAnswerOptionResponse(answerOption model.AnswerOption) web.AnswerOptionResponse {
	return web.AnswerOptionResponse{
		ID:         int(answerOption.ID),
		QuestionId: answerOption.QuestionId,
		OptionText: answerOption.OptionText,
		IsCorrect:  answerOption.IsCorrect,
	}
}



func ToUserQuizResultResponse(userQuizResult model.UserQuizResult) web.UserQuizResultResponse {
	return web.UserQuizResultResponse{
		ID: int(userQuizResult.ID),
		UserId: userQuizResult.UserId,
		User: web.UserResponseMini{ID: userQuizResult.User.ID, Name: userQuizResult.User.Name},
		QuizId: userQuizResult.QuizId,
		Quiz: web.QuizResponseMini{ID: int(userQuizResult.Quiz.ID), Title: userQuizResult.Quiz.Title},
		Score: userQuizResult.Score,
		TotalQuestions: userQuizResult.TotalQuestions,
		CorrectAnswers: userQuizResult.CorrectAnswers,
	}
}
