package main

import (
	"awesomeProject1/internal/db"
	"awesomeProject1/internal/handlers"
	"awesomeProject1/internal/taskservice"
	"awesomeProject1/internal/web/tasks"
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
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	strictHandler := tasks.NewStrictHandler(taskHandlers, nil) // тут будет ошибка
	tasks.RegisterHandlers(e, strictHandler)

	if err := e.Start("localhost:8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}

//	e.GET("/", taskHandlers.GetTasks)
//	e.POST("/", taskHandlers.PostTask)
//	e.PATCH("/:id", taskHandlers.PatchTask)
//	e.DELETE("/:id", taskHandlers.DeleteTask)
//	e.Start("localhost:8080")
