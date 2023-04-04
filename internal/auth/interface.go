package auth

import (
	"context"
)

//go:generate mockery --name=IAuth --filename=auth.go
type IAuth interface {
	SignIn(ctx context.Context, username, password string) (string, error)
	Validate(ctx context.Context, token string) error
}
