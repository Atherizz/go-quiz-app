package handler

import (
	"google-oauth/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserQuizResultHandler struct {
	Service *service.UserQuizResultService
}

func NewUserQuizResultHandler(service *service.UserQuizResultService) *UserQuizResultHandler {
	return &UserQuizResultHandler{
		Service: service,
	}
}


func (handler *UserQuizResultHandler) GetUserQuizResultGroupByQuiz(c *gin.Context) {
	id := c.Param("quizId")
	intId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := handler.Service.GetUserQuizResultGroupByQuiz(c.Request.Context(), intId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (handler *UserQuizResultHandler) GetQuizResultGroupByQuizAndUser(c *gin.Context) {
	userId := c.Param("userId")
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	quizId := c.Param("quizId")
	quizIdInt, err := strconv.Atoi(quizId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := handler.Service.GetQuizResultGroupByQuizAndUser(c.Request.Context(), quizIdInt,userIdInt)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (handler *UserQuizResultHandler) GetUserQuizResultGroupByUser(c *gin.Context) {
	id := c.Param("userId")
	intId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := handler.Service.GetUserQuizResultGroupByUser(c.Request.Context(), intId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (handler *UserQuizResultHandler) Delete(c *gin.Context) {

	id := c.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = handler.Service.Delete(c.Request.Context(), intId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "Success delete data")
}

