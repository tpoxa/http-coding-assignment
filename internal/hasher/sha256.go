package hasher

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
)

type Sha256 struct {
}

func NewSha256() *Sha256 {
	return &Sha256{}
}

func (s Sha256) Hash(_ context.Context, in string) string {
	h := sha256.New()
	h.Write([]byte(in))
	return hex.EncodeToString(h.Sum(nil))
}
