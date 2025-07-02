package handler

import (
	"google-oauth/service"
	"google-oauth/web"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserAnswerHandler struct {
	Service *service.UserAnswerService
}

func NewUserAnswerHandler(service *service.UserAnswerService) *UserAnswerHandler {
	return &UserAnswerHandler{
		Service: service,
	}
}

func (handler *UserAnswerHandler) SaveAllAnswers(c *gin.Context) {
	newUserAnswer := web.SubmitQuizRequest{}

	if err := c.ShouldBindJSON(&newUserAnswer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := handler.Service.SaveAllAnswers(c.Request.Context(), newUserAnswer)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}


func (handler *UserAnswerHandler) Delete(c *gin.Context) {

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

