package main

import (
	"net/http"

	"github.com/AngelGomezEspinosa/shadcn-admin-backend/internal/core/services"
	"github.com/AngelGomezEspinosa/shadcn-admin-backend/internal/handlers"
	"github.com/AngelGomezEspinosa/shadcn-admin-backend/internal/repositories"
	"github.com/gin-gonic/gin"
)

func main() {

	//gamesRepository := gamesrepo.NewMemKVS()
	//gamesService := gamesrv.New(gamesRepository, uidgen.New())
	//gamesHandler := gamehdl.NewHTTPHandler(gamesService)

	taskRepository := repositories.NewHardcodeRepository()
	taskService := services.NewTaskService(taskRepository)
	taskHandler := handlers.NewHttpHandler(taskService)

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.GET("/taskall", taskHandler.FindAll)

	router.Run() // listen and serve on 0.0.0.0:8080
}
