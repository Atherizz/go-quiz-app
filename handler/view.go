package handler

import (
	"google-oauth/helper"
	"google-oauth/model"
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

	session, _ := helper.Store.Get(c.Request, "user_info")

	user, ok := session.Values["user"].(model.User)
	if !ok || user.Email == "" || user.Name == "" {
		http.Error(c.Writer, "unauthorized", http.StatusUnauthorized)
		return
	}
	// fmt.Fprint(writer, "welcome ", name)
	tmpl := template.Must(template.ParseFiles("./resources/auth/welcome.gohtml"))
	tmpl.ExecuteTemplate(c.Writer, "welcome.gohtml", user.Name)

}

func ProfileView(c *gin.Context) {
	session, _ := helper.Store.Get(c.Request, "user_info")

	user, ok := session.Values["user"].(model.User)
	if !ok || user.Email == "" || user.Name == "" {
		http.Error(c.Writer, "unauthorized", http.StatusUnauthorized)
		return
	}
	// fmt.Fprint(writer, "welcome ", name)
	http.ServeFile(c.Writer, c.Request, "./resources/auth/profile.html")

}
