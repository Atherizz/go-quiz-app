package web

type SubjectRequest struct {
	ID            int `json:"id"`
	SubjectName          string `json:"subject_name" binding:"required"`
}
