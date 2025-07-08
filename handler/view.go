package handler

import (
	"google-oauth/helper"
	"html/template"

	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginView(c *gin.Context) {
	// fmt.Fprint(writer, "welcome ", name)
	http.ServeFile(c.Writer, c.Request, "./resources/auth/login.html")
}

func RegisterView(c *gin.Context) {
	// fmt.Fprint(writer, "welcome ", name)
	http.ServeFile(c.Writer, c.Request, "./resources/auth/register.html")
}

func HomeView(c *gin.Context) {

	user, exists := c.Get("user")
	authUser := user.(helper.UserSession)

	if !exists || authUser.Email == "" || authUser.Name == "" {
		http.Error(c.Writer, "unauthorized", http.StatusUnauthorized)
		return
	}

	tmpl := template.Must(template.ParseFiles("./resources/auth/welcome.gohtml"))
	tmpl.ExecuteTemplate(c.Writer, "welcome.gohtml", authUser.Name)
	

}

func ProfileView(c *gin.Context) {
	user, exists := c.Get("user")
	authUser := user.(helper.UserSession)
	
	if !exists || authUser.Email == "" || authUser.Name == "" {
		http.Error(c.Writer, "unauthorized", http.StatusUnauthorized)
		return
	}

	http.ServeFile(c.Writer, c.Request, "./resources/auth/profile.html")

}
