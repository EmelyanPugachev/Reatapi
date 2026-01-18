package main

import (
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func initDB() {
	dsn := "host=localhost user=postgres password=yourpassword dbname=postgres port=5432 sslmode=disable"
	var err error

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := db.AutoMigrate(&Taska{}); err != nil {
		log.Fatalf("Failed to migrate tasks: %v", err)
	}
}

type RequestBody struct {
	Task string `json:"task"`
}

type Taska struct {
	Task string `json:"task"`
	ID   string `gorm:"primarykey" json:"id"`
}

func getTasks(c echo.Context) error {
	var tasks = []Taska{}

	if err := db.Find(&tasks).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to get tasks"})
	}
	return c.JSON(http.StatusOK, tasks)
}

func postTasks(c echo.Context) error {
	var body RequestBody
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}
	newtask := Taska{
		Task: "Hello " + body.Task,
		ID:   uuid.NewString(),
	}
	if err := db.Create(&newtask).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create task"})
	}
	return c.JSON(http.StatusAccepted, newtask)
}

func patchTask(c echo.Context) error {
	id := c.Param("id")
	var body RequestBody
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}
	var newtask Taska
	if err := db.First(&newtask, "id=?", id).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Failed to get task"})
	}
	newtask.Task = body.Task
	if err := db.Save(&newtask).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update task"})
	}
	return c.JSON(http.StatusOK, newtask)
}

func deleteTask(c echo.Context) error {
	id := c.Param("id")
	if err := db.Delete(&Taska{}, "id=?", id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete task"})
	}
	return c.NoContent(http.StatusAccepted)
}

func main() {
	initDB()
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.GET("/", getTasks)
	e.POST("/", postTasks)
	e.PATCH("/:id", patchTask)
	e.DELETE("/:id", deleteTask)
	e.Start("localhost:8080")
}
