package repository

import (
	"context"
	"go-sample-echo/domain/object/user"
	"go-sample-echo/domain/object/value"
)

type UserRepository interface {
	Create(ctx context.Context, user user.User) (user.User, error)
	Find(ctx context.Context, id value.UserId) (user.User, error)
}
