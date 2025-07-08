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
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"golang.org/x/oauth2"
)

type AuthHandler struct {
	Service *service.UserService
}

func NewAuthHandler(service *service.UserService) *AuthHandler {
	return &AuthHandler{
		Service: service,
	}
}

func (handler *AuthHandler) BasicOauth(c *gin.Context) {
	fmt.Fprint(c.Writer, "selamat datang di endpoint basic auth! anda berhasil terautentikasi \n")
}

func (handler *AuthHandler) LoginOauth(c *gin.Context) {
	url := middleware.OauthConfig.AuthCodeURL("", oauth2.AccessTypeOffline)
	c.Redirect(http.StatusSeeOther, url)
	// http.Redirect(c.Writer, c.Request, url, http.StatusSeeOther)

}

func (handler *AuthHandler) RegisterDefault(c *gin.Context) {
	registeredUser := web.UserRequest{}

	if err := c.ShouldBindJSON(&registeredUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := handler.Service.RegisterDefault(c.Request.Context(), registeredUser)

	c.JSON(http.StatusOK, response)

}

func (handler *AuthHandler) Callback(c *gin.Context) {
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

	// tokenPayload, err := helper.DecodeIdToken(idToken)

	tokenPayload, err := helper.VerifyGoogleIdToken(c.Request.Context(), idToken, middleware.OauthConfig.ClientID)
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

	userResponse := handler.Service.GetUserByEmail(c.Request.Context(), tokenPayload.Claims["email"].(string))

	if userResponse.Email == "" {
		userRequest := model.User{
			GoogleId: tokenPayload.Claims["sub"].(string),
			Name:     tokenPayload.Claims["name"].(string),
			Email:    tokenPayload.Claims["email"].(string),
			Picture:  tokenPayload.Claims["picture"].(string),
		}

		handler.Service.RegisterFromGoogle(c.Request.Context(), userRequest)
	}

	sessionId := uuid.NewString()

	userData := map[string]interface{}{
		"name":    tokenPayload.Claims["name"],
		"email":   tokenPayload.Claims["email"],
		"picture": tokenPayload.Claims["picture"],
		"sub":     tokenPayload.Claims["sub"],
	}

	jsonData, _ := json.Marshal(userData)

	client := helper.Client

	client.SetEx(c.Request.Context(), "session:"+sessionId, jsonData, 60*time.Minute)
	c.SetCookie("session_id", sessionId, 3600, "/", "localhost", false, true)

	c.Redirect(http.StatusSeeOther, "/home")

}

func (handler *AuthHandler) Logout(c *gin.Context) {

	c.SetCookie("oauth_token", "", -1, "/", "localhost", false, true)
	c.Redirect(http.StatusFound, "/login")
}
