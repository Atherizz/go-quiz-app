package web


type QuizResponse struct {
	ID          int `json:"id"`
	SubjectId   int `json:"subject_id"`
	Subject     SubjectResponse `json:"subject"`
	Questions   []QuestionResponse `json:"questions"`
	Description string `json:"description"`
	Title       string `json:"title"`
}


type QuizResponseMini struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}