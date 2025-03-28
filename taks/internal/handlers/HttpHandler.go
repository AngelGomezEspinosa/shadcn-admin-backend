package handlers

import (
	"net/http"

	"github.com/AngelGomezEspinosa/shadcn-admin-backend/internal/core/ports"
	"github.com/gin-gonic/gin"
)

type HttpHandler struct {
	TaskService ports.TaskService
}

func NewHttpHandler(taskService ports.TaskService) *HttpHandler {
	return &HttpHandler{TaskService: taskService}
}

func enableCors(w *http.ResponseWriter) {
	// Permitir cualquier origen
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	// Permitir ciertos métodos
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	// Permitir ciertos encabezados
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Llamar a la función de habilitación de CORS
	enableCors(&w)

	if r.Method == "OPTIONS" {
		// Responder a la solicitud OPTIONS (pre-flight request)
		w.WriteHeader(http.StatusOK)
		return
	}
}

func (h *HttpHandler) FindAll(c *gin.Context) {
	task, err := h.TaskService.FindAll()
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}
	// Handle task
	handler(c.Writer, c.Request)
	c.JSON(200, task)
}
