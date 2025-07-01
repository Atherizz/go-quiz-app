package web

import "google-oauth/model"

type SubjectResponse struct {
	Id          int `json:"id"`
	SubjectName string `json:"subject_name"`
	Quiz        []model.Quiz `json:"quiz"`
}