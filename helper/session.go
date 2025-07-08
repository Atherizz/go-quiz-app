package helper

import (
	"github.com/gorilla/sessions"
)

var Store = sessions.NewCookieStore([]byte(LoadEnv("SESSION_SECRET")))

type UserSession struct {
	Name string
	Email string
	Picture string
	Sub string
}
