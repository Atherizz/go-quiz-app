package web

type UserRequest struct {
	Name            string `json:"name" binding:"required,min=1,max=200"`
	Email           string `json:"email" binding:"required,email"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"eqfield=Password"`
}
