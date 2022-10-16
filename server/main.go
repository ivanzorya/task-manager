package main

import (
	"os"

	"server/routes"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

    router.Use(cors.Default())

	//C
	router.POST("/task/create", routes.AddTask)
	//R
	router.GET("/tasks", routes.GetTasks)
	router.GET("/task/:id/", routes.GetTaskById)
	//U
	router.PUT("/task/update/:id", routes.UpdateTask)
	//D
	router.DELETE("/task/delete/:id", routes.DeleteTask)

	router.Run(":" + port)
}
