package main

import (
	"github.com/labstack/echo/v4"
	"go-sample-echo/domain/usecase"
	"go-sample-echo/infrastructure/dao"
	"go-sample-echo/ui/api"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})

	userHandler := api.NewUserHandler(usecase.NewUserUseCase(dao.NewUserRepositoryImpl()))
	e.GET("/users/:id", userHandler.Get)
	e.POST("/users", userHandler.Create)
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
