package usecase

import (
	"context"
	"go-sample-echo/domain/object/user"
	"go-sample-echo/domain/object/value"
	"go-sample-echo/domain/repository"
)

type UserUseCase interface {
	Create(ctx context.Context, name string) (*user.User, error)
	Find(ctx context.Context, id value.UserId) (*user.User, error)
}

type userUseCase struct {
	repo repository.UserRepository
}

func NewUserUseCase(userRepository repository.UserRepository) UserUseCase {
	return &userUseCase{repo: userRepository}
}

func (u userUseCase) Create(ctx context.Context, name string) (*user.User, error) {
	created, err := u.repo.Create(ctx, user.NewUser(value.NewUserId(), name))
	if err != nil {
		return nil, err
	}
	return &created, nil
}

func (u userUseCase) Find(ctx context.Context, id value.UserId) (*user.User, error) {
	found, err := u.repo.Find(ctx, id)
	if err != nil {
		return nil, err
	}
	return &found, nil
}
