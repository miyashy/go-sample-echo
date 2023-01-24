package api

import (
	"github.com/labstack/echo/v4"
	"go-sample-echo/domain/object/value"
	"go-sample-echo/domain/usecase"
	"net/http"
)

type UserDto struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type GetUserDto struct {
	Id string `param:"id"`
}

type CreateUserDto struct {
	Name string `json:"name"`
}

type UserHandler interface {
	Get(ctx echo.Context) error
	Create(ctx echo.Context) error
}

type userHandler struct {
	useCase usecase.UserUseCase
}

func NewUserHandler(useCase usecase.UserUseCase) UserHandler {
	return &userHandler{useCase: useCase}
}

func (u userHandler) Get(ctx echo.Context) error {
	var req GetUserDto
	err := ctx.Bind(&req)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}

	id, err := value.NewUserIdWithHint(req.Id)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}

	result, err := u.useCase.Find(ctx.Request().Context(), id)
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}

	return ctx.JSON(http.StatusOK, UserDto{
		Id:   result.Id.Value(),
		Name: result.Name,
	})
}

func (u userHandler) Create(ctx echo.Context) error {
	var req CreateUserDto
	err := ctx.Bind(&req)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}

	result, err := u.useCase.Create(ctx.Request().Context(), req.Name)
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}

	return ctx.JSON(http.StatusOK, UserDto{
		Id:   result.Id.Value(),
		Name: result.Name,
	})
}
