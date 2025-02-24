package account

import "context"

type User struct {
	ID       string `json: "id,omitempty"`
	Email    string `json: "email"`
	Password string `json: "password"`
}

type Repository interface {
	CreateUser(ctx context.Context, u User) error
	GetUser(ctx context.Context, id string) (email string, err error)
}
