package user

import "go-sample-echo/domain/object/value"

type User struct {
	Id   value.UserId
	Name string
}

func NewUser(id value.UserId, name string) User {
	return User{id, name}
}
