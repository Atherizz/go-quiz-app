package handler

import (
	"encoding/json"
	"fmt"
	"google-oauth/helper"
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

	client := helper.Client

	
	// client.XGroupCreateConsumer(c.Request.Context(), "quiz_answer_stream", "group-1", "consumer-1")

	key2 := fmt.Sprintf("rate:user:%d:quiz:%d", newUserAnswer.UserId, newUserAnswer.QuizId)
	count, err := client.Incr(c.Request.Context(), key2).Result()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if count > 1 {
		c.JSON(http.StatusTooManyRequests, gin.H{"error": "you already submit this quiz!"})
		return
	}

	// user, exists := c.Get("user")
	// if !exists {
	// 	c.JSON(404, gin.H{"error": "value not found"})
	// 	return
	// }
	// authUser := user.(model.User)

	// newUserAnswer.UserId = 1
	newUserAnswer.QuizId = intId

	jsonBody, err := json.Marshal(newUserAnswer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = client.XAdd(c.Request.Context(), &redis.XAddArgs{
		Stream: "quiz_answer_stream",
		Values: map[string]interface{}{
			"payload": string(jsonBody),
		},
	}).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to add to stream"})
		return
	}
	// response, err := handler.Service.SaveAllAnswers(c.Request.Context(), newUserAnswer, intId)

	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// key := "score:" + strconv.Itoa(response.Quiz.ID)
	// client.ZAdd(c.Request.Context(), key, redis.Z{
	// 	Member: response.User.Name,
	// 	Score:  response.Score,
	// })

	c.JSON(http.StatusOK, gin.H{"message": "Jawaban berhasil dikirim dan sedang diproses!"})
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
