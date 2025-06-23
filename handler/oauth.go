package handler

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"google-oauth/helper"
	"google-oauth/middleware"
	"google-oauth/model"
	"google-oauth/service"
	"google-oauth/web"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

type OauthController struct {
	Service service.UserService
}

func NewOauthController(service *service.UserService) *OauthController {
	return &OauthController{
		Service: *service,
	}
}

func (controller *OauthController) BasicOauth(c *gin.Context) {
	fmt.Fprint(c.Writer, "selamat datang di endpoint basic auth! anda berhasil terautentikasi \n")
}

func (controller *OauthController) LoginOauth(c *gin.Context) {
	url := middleware.OauthConfig.AuthCodeURL("", oauth2.AccessTypeOffline)
	c.Redirect(http.StatusSeeOther, url)
	// http.Redirect(c.Writer, c.Request, url, http.StatusSeeOther)

}

func (controller *OauthController) RegisterDefault(c *gin.Context) {
	registeredUser := web.UserRequest{}

	if err := c.ShouldBindJSON(&registeredUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := controller.Service.RegisterDefault(c.Request.Context(), registeredUser)

	helper.WriteEncodeResponse(c.Writer, web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})

}

func (controller *OauthController) Callback(c *gin.Context) {
	code := c.Request.URL.Query().Get("code")
	token, err := middleware.OauthConfig.Exchange(c.Request.Context(), code)
	if err != nil {
		http.Error(c.Writer, "failed get token", http.StatusInternalServerError)
		return
	}

	idToken, ok := token.Extra("id_token").(string)

	if !ok {
		http.Error(c.Writer, "no id_token in field token", http.StatusInternalServerError)
	}

	tokenPayload, err := helper.DecodeIdToken(idToken)
	if err != nil {
		http.Error(c.Writer, "failed decode token", http.StatusInternalServerError)
	}

	tokenJson, err := json.Marshal(token)
	if err != nil {
		http.Error(c.Writer, "failed to marshal token", http.StatusInternalServerError)
		return
	}

	encoded := base64.StdEncoding.EncodeToString(tokenJson)

	cookie, err := c.Cookie("oauth_token")

	if err != nil {
		cookie = "NotSet"
		c.SetCookie("oauth_token", encoded, 3600, "/", "localhost", false, true)
	}

	fmt.Println("OAuth token:", cookie)

	userResponse := controller.Service.GetUserByEmail(c.Request.Context(), tokenPayload.Email)

	if userResponse.Email == "" {
		userRequest := model.User{
			GoogleId: tokenPayload.Sub,
			Name:     tokenPayload.Name,
			Email:    tokenPayload.Email,
			Picture:  tokenPayload.Picture,
		}

		controller.Service.RegisterFromGoogle(c.Request.Context(), userRequest)
	}

	session, _ := helper.Store.Get(c.Request, "user_info")
	session.Values["user"] = model.User{
		Name:     tokenPayload.Name,
		Email:    tokenPayload.Email,
		Picture:  tokenPayload.Picture,
		GoogleId: tokenPayload.Sub,
	}

	err = session.Save(c.Request, c.Writer)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	c.Redirect(http.StatusFound, "/home")
	// http.Redirect(c.Writer, c.Request, "/home", http.StatusFound)
}

func (controller *OauthController) Logout(c *gin.Context) {

	c.SetCookie("oauth_token", "", -1, "/", "localhost", false, true)
	c.Redirect(http.StatusFound, "/login")
	// http.Redirect(c.Writer, c.Request, "/login", http.StatusFound)
}
