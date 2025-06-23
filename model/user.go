package model

import "time"

type User struct {
	ID          int `gorm:"primaryKey"`
	GoogleId    string
	Name        string
	Email       string
	Picture     string
	Provider    string
	Role        string
	Password 	string
	LastLoginAt time.Time
}
