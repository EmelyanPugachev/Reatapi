package main

import (
	"awesomeProject1/internal/db"
	"awesomeProject1/internal/handlers"
	"awesomeProject1/internal/taskservice"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	e := echo.New()
	taskRepo := taskservice.NewTaskRepository(database)
	taskService := taskservice.NewTaskService(taskRepo)
	taskHandlers := handlers.NewTaskHandler(taskService)
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.GET("/", taskHandlers.GetTasks)
	e.POST("/", taskHandlers.PostTask)
	e.PATCH("/:id", taskHandlers.PatchTask)
	e.DELETE("/:id", taskHandlers.DeleteTask)
	e.Start("localhost:8080")
}
