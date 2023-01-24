package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})
	e.GET("/users/:userId", func(c echo.Context) error {
		return c.JSON(http.StatusOK, User{Id: "dummy", Name: "Name"})
	})
	e.POST("/users/", func(c echo.Context) error {
		return c.JSON(http.StatusCreated, User{Id: "dummy", Name: "Name"})
	})
	e.GET("/tasks/:taskId", func(c echo.Context) error {
		return c.JSON(http.StatusOK, Task{
			Id:          "dummy",
			UserId:      "dummyUserID",
			Status:      "TODO",
			Description: "DummyTask",
		})
	})
	e.PATCH("/tasks/:taskId", func(c echo.Context) error {
		return c.JSON(http.StatusCreated, Task{
			Id:          "dummy",
			UserId:      "dummyUserID",
			Status:      "TODO",
			Description: "DummyTask",
		})
	})
	e.DELETE("/tasks/:taskId", func(c echo.Context) error {
		return c.NoContent(http.StatusNoContent)
	})
	e.POST("/tasks", func(c echo.Context) error {
		return c.JSON(http.StatusCreated, Task{
			Id:          "dummy",
			UserId:      "dummyUserID",
			Status:      "TODO",
			Description: "DummyTask",
		})
	})
	e.Logger.Fatal(e.Start(":1234"))
}

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Task struct {
	Id          string `json:"id"`
	UserId      string `json:"userId"`
	Status      string `json:"status"`
	Description string `json:"description"`
}
