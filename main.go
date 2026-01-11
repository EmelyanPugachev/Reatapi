package main

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type RequestBody struct {
	Task string `json:"task"`
}

type Taska struct {
	Task string `json:"task"`
	ID   string `json:"id"`
}

var tasks = []Taska{}

func getTasks(c echo.Context) error {
	return c.JSON(http.StatusOK, tasks)
}

func postTasks(c echo.Context) error {
	var body RequestBody
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}
	newtask := Taska{
		Task: body.Task,
		ID:   uuid.NewString(),
	}
	tasks = append(tasks, newtask)
	return c.JSON(http.StatusAccepted, tasks)
}

func patchTask(c echo.Context) error {
	var body RequestBody
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}
	return c.JSON(http.StatusAccepted, tasks)
}

func main() {
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.GET("/", getTasks)
	e.POST("/", postTasks)
	e.Start("localhost:8080")

}
