package web

type SubjectRequest struct {
	ID            int `json:"id" binding:"required"`
	SubjectName          string `json:"subject_name" binding:"required,email"`
}
