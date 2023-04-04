package hasher

import "context"

//go:generate mockery --name=Hasher --filename=hasher.go
type Hasher interface {
	Hash(ctx context.Context, in string) string
}
