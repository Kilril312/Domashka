package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

var db *gorm.DB

func initDB() {
	dsn := "host=localhost user= postgres password=qwerty dbname=postgres port=5438 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	if err := db.AutoMigrate(&requestBodyTask{}); err != nil {
		log.Fatalf("Failed to migrate tasks: %v", err)
	}
}

var id = 1

type requestBodyTask struct {
	ID   int    `gorm:"primaryKey" json:"id"`
	Task string `json:"task"`
}

type ResponseBodyTask struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func getTaskHandler(c echo.Context) error {
	var tasks []requestBodyTask
	if err := db.Find(&tasks).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseBodyTask{
			Status:  "error",
			Message: "Could not get tasks",
		})
	}
	return c.JSON(http.StatusOK, tasks)
}

func postTaskHandler(c echo.Context) error {
	var task requestBodyTask
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, ResponseBodyTask{
			Status:  "error",
			Message: "Error binding task",
		})
	}
	if err := db.Create(&task).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseBodyTask{
			Status:  "error",
			Message: "Could not create task",
		})
	}
	return c.JSON(http.StatusCreated, task)
}

func patchTaskHandler(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ResponseBodyTask{
			Status:  "error",
			Message: "Error converting id to int",
		})
	}
	var updatedTask requestBodyTask
	if err := c.Bind(&updatedTask); err != nil {
		return c.JSON(http.StatusBadRequest, ResponseBodyTask{
			Status:  "error",
			Message: "Error binding task",
		})
	}

	var task requestBodyTask

	if err := db.First(&task, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseBodyTask{
			Status:  "error",
			Message: "Could not id",
		})
	}
	task.Task = updatedTask.Task
	if err := db.Save(&task).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseBodyTask{
			Status:  "error",
			Message: "Could not update task",
		})
	}
	return c.JSON(http.StatusOK, task)
}

func deleteTaskHandler(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ResponseBodyTask{
			Status:  "error",
			Message: "Error converting id to int",
		})
	}
	if err := db.Delete(&requestBodyTask{}, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseBodyTask{
			Status:  "error",
			Message: "Could not delete task",
		})
	}
	return c.NoContent(http.StatusNoContent)
}

func main() {
	initDB()
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.GET("/tasks", getTaskHandler)
	e.POST("/tasks", postTaskHandler)
	e.PATCH("/tasks/:id", patchTaskHandler)
	e.DELETE("/tasks/:id", deleteTaskHandler)
	e.Start("localhost:8080")

}
