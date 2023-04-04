package auth

import "fmt"

var (
	ErrInvalidCredentials = fmt.Errorf("invalid credentials")
	ErrInvalidToken       = fmt.Errorf("invalid token")
)
