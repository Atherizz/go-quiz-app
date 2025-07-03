package handler

import (
	"google-oauth/service"
	"google-oauth/web"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AnswerOptionHandler struct {
	Service *service.AnswerOptionService
}

func NewAnswerOptionHandler(service *service.AnswerOptionService) *AnswerOptionHandler {
	return &AnswerOptionHandler{
		Service: service,
	}
}

func (handler *AnswerOptionHandler) Insert(c *gin.Context) {
	newAnswerOption := web.AnswerOptionRequest{}

	id := c.Param("questionId")
	intId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}



	if err := c.ShouldBindJSON(&newAnswerOption); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	response, err := handler.Service.Insert(c.Request.Context(), newAnswerOption, intId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (handler *AnswerOptionHandler) Update(c *gin.Context) {
	updateAnswerOption := web.AnswerOptionRequest{}

	if err := c.ShouldBindJSON(&updateAnswerOption); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("questionId")
	intId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}


	response, err := handler.Service.Update(c.Request.Context(), updateAnswerOption, intId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (handler *AnswerOptionHandler) GetAnswerOptionGroupByQuestion(c *gin.Context) {
	id := c.Param("questionId")
	intId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := handler.Service.GetAnswerOptionGroupByQuestion(c.Request.Context(), intId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (handler *AnswerOptionHandler) Delete(c *gin.Context) {

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

func (handler *AnswerOptionHandler) GetAnswerOptionById(c *gin.Context) {

	id := c.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := handler.Service.GetAnswerOptionById(c.Request.Context(), intId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}
