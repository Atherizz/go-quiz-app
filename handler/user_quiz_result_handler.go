package handler

import (
	"google-oauth/helper"
	"google-oauth/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserQuizResultHandler struct {
	Service *service.UserQuizResultService
}

type UserScore struct {
	User string `json:"user"`
	Score float64 `json:"score"`
}

type LeaderboardResponse struct {
	UserScore []UserScore `json:"user_score"`

}
func NewUserQuizResultHandler(service *service.UserQuizResultService) *UserQuizResultHandler {
	return &UserQuizResultHandler{
		Service: service,
	}
}

func (handler *UserQuizResultHandler) Leaderboard(c *gin.Context) {
	id := c.Param("quizId")

	var userScores []UserScore

	client := helper.Client
	slice, err := client.ZRangeWithScores(c.Request.Context(), "score:"+id, 0, -1).Result()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, z := range slice {
		userScore := UserScore{
			User: z.Member.(string),
			Score: z.Score,
		}
		userScores = append(userScores, userScore)
}

	response := LeaderboardResponse{
		UserScore: userScores,
	}

	c.JSON(http.StatusOK, response)
}

func (handler *UserQuizResultHandler) GetQuizResultGroupByQuizAndUser(c *gin.Context) {
	// user, exists := c.Get("user")
	// if !exists {
	// 	c.JSON(404, gin.H{"error": "value not found"})
	// 	return
	// }
	// authUser := user.(model.User)

	quizId := c.Param("quizId")
	quizIdInt, err := strconv.Atoi(quizId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := handler.Service.GetQuizResultGroupByQuizAndUser(c.Request.Context(), quizIdInt, 2)

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
