package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type requestBody struct {
	Task string `json:"task"`
}

var task string

func getTask(c echo.Context) error {
	return c.JSON(http.StatusOK, task)
}

func postTask(c echo.Context) error {
	var body requestBody
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}
	task = body.Task
	return c.JSON(http.StatusAccepted, task)
}

func main() {
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.GET("/", getTask)
	e.POST("/", postTask)
	e.Start("localhost:8080")

}
