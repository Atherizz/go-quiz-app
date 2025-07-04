package handler

import (
	"google-oauth/service"
	"google-oauth/web"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SubjectHandler struct {
	Service *service.SubjectService
}

func NewSubjectHandler(service *service.SubjectService) *SubjectHandler {
	return &SubjectHandler{
		Service: service,
	}
}

func (handler *SubjectHandler) Insert(c *gin.Context) {
	newSubject := web.SubjectRequest{}

	if err := c.ShouldBindJSON(&newSubject); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := handler.Service.Insert(c.Request.Context(), newSubject)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (handler *SubjectHandler) Update(c *gin.Context) {
	updateSubject := web.SubjectRequest{}

	if err := c.ShouldBindJSON(&updateSubject); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := handler.Service.Update(c.Request.Context(), updateSubject)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (handler *SubjectHandler) GetAll(c *gin.Context) {

	response, err := handler.Service.GetAll(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (handler *SubjectHandler) Delete(c *gin.Context) {

	id := c.Param("subjectId")
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

func (handler *SubjectHandler) GetSubjectById(c *gin.Context) {

	id := c.Param("subjectId")
	intId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := handler.Service.GetSubjectById(c.Request.Context(), intId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}
