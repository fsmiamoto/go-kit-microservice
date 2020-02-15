package account

import "context"

type Service interface {
	CreateUser(ctx context.Context, email, password string) (id string, err error)
	GetUser(ctx context.Context, id string) (string, error)
}
