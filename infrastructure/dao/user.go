package dao

import (
	"context"
	"go-sample-echo/domain/object/user"
	"go-sample-echo/domain/object/value"
	"go-sample-echo/domain/repository"
)

type userRepositoryImpl struct {
}

func NewUserRepositoryImpl() repository.UserRepository {
	return &userRepositoryImpl{}
}

func (u userRepositoryImpl) Create(ctx context.Context, user user.User) (user.User, error) {
	return user, nil
}

func (u userRepositoryImpl) Find(ctx context.Context, id value.UserId) (user.User, error) {
	return user.User{
		Id:   id,
		Name: "Dummy",
	}, nil
}
