package handler

import (
	"google-oauth/helper"
	"google-oauth/model"
	"google-oauth/service"
	"google-oauth/web"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
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

	id := c.Param("quizId")
	intId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&newUserAnswer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, exists := c.Get("user")
	if !exists {
		c.JSON(404, gin.H{"error": "value not found"})
		return
	}
	authUser := user.(model.User)

	newUserAnswer.UserId = authUser.ID
	newUserAnswer.QuizId = intId

	response, err := handler.Service.SaveAllAnswers(c.Request.Context(), newUserAnswer, intId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client := helper.Client

	key := "score:" + strconv.Itoa(response.Quiz.ID)
	client.ZAdd(c.Request.Context(), key, redis.Z{
		Member: response.User.Name,
		Score: response.Score,
	})

	c.JSON(http.StatusOK, response)
}

func (handler *UserAnswerHandler) Delete(c *gin.Context) {

	id := c.Param("userAnswerId")
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
