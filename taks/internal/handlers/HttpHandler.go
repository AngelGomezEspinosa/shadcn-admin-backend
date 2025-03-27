package handlers

import (
	"github.com/AngelGomezEspinosa/shadcn-admin-backend/internal/core/ports"
	"github.com/gin-gonic/gin"
)

type HttpHandler struct {
	TaskService ports.TaskService
}

func NewHttpHandler(taskService ports.TaskService) *HttpHandler {
	return &HttpHandler{TaskService: taskService}
}

func (h *HttpHandler) FindAll(c *gin.Context) {
	task, err := h.TaskService.FindAll()
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}
	// Handle task
	c.JSON(200, task)
}
