package value

import "github.com/google/uuid"

type UserId string

func NewUserId() UserId {
	return UserId(uuid.New().String())
}

func NewUserIdWithHint(value string) (UserId, error) {
	_, err := uuid.Parse(value)
	if err != nil {
		return UserId(""), err
	}
	return UserId(value), nil
}

func (userId UserId) Value() string {
	return string(userId)
}
