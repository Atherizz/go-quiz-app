package handler

import (
	"encoding/json"
	"google-oauth/helper"
	"google-oauth/service"
	"google-oauth/web"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type QuestionHandler struct {
	Service *service.QuestionService
}

func NewQuestionHandler(service *service.QuestionService) *QuestionHandler {
	return &QuestionHandler{
		Service: service,
	}
}

func (handler *QuestionHandler) Insert(c *gin.Context) {
	newQuestion := web.QuestionRequest{}

	id := c.Param("quizId")
	intId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&newQuestion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := handler.Service.Insert(c.Request.Context(), newQuestion, intId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (handler *QuestionHandler) Update(c *gin.Context) {
	updateQuestion := web.QuestionRequest{}
	id := c.Param("quizId")
	intId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.ShouldBindJSON(&updateQuestion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := handler.Service.Update(c.Request.Context(), updateQuestion, intId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (handler *QuestionHandler) GetQuestionGroupByQuiz(c *gin.Context) {
	id := c.Param("quizId")
	intId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var responses []web.QuestionResponse

	
	client := helper.Client

	result, err := client.Get(c.Request.Context(), "quiz:"+id).Result()

	if err == redis.Nil {

		responses, err = handler.Service.GetQuestionGroupByQuiz(c.Request.Context(), intId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		jsonResponse, _ := json.Marshal(responses)
		client.SetEx(c.Request.Context(), "quiz:"+id, jsonResponse, 2 * time.Hour)

	} else if err != nil {  
    log.Println("❌ Redis error:", err)
    c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
    return
	} else {
		err = json.Unmarshal([]byte(result), &responses)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, responses)
}

func (handler *QuestionHandler) Delete(c *gin.Context) {

	id := c.Param("questionId")
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

func (handler *QuestionHandler) GetQuestionById(c *gin.Context) {

	id := c.Param("questionId")
	intId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := handler.Service.GetQuestionById(c.Request.Context(), intId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}
