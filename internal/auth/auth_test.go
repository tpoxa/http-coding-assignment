package auth

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestInvalidAuth_SignIn(t *testing.T) {
	key := []byte("pk")
	auth := NewAuth(key)
	ctx := context.Background()

	type args = struct {
		username string
		password string
	}

	for _, a := range []args{
		{

			username: "",
			password: "",
		},
		{
			username: "u",
			password: "",
		},
		{
			username: "",
			password: "p",
		},
	} {
		_, err := auth.SignIn(ctx, a.username, a.password)
		assert.Errorf(t, err, "invalid credentials should cause error")
		assert.Equal(t, err, ErrInvalidCredentials)
	}
}

func TestSubject_SignIn(t *testing.T) {
	key := []byte("pk")
	auth := NewAuth(key)

	ctx := context.Background()
	type args = struct {
		username string
		password string
	}

	for _, a := range []args{
		{
			username: "subj",
			password: "p",
		},
		{
			username: "a",
			password: "b",
		},
	} {

		token, err := auth.SignIn(ctx, a.username, a.password)
		parsed, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			return key, nil
		})
		assert.NoError(t, err, "token parse should not cause an error")
		s, _ := parsed.Claims.GetSubject()
		// JWT subject equals username
		assert.Equal(t, s, a.username)
	}
}

func TestPrivKey_SignIn(t *testing.T) {

	ctx := context.Background()
	auth := NewAuth([]byte("key1"))

	token, err := auth.SignIn(ctx, "u", "p")
	assert.NoError(t, err, "no error")
	err = auth.Validate(ctx, token)
	assert.NoError(t, err, "same key no error")

	auth2 := NewAuth([]byte("key2"))
	err = auth2.Validate(ctx, token)
	assert.Errorf(t, err, "wrong key should error")
	assert.Equal(t, err, ErrInvalidToken)
}

func TestExpiry_SignIn(t *testing.T) {
	key := []byte("pk")
	auth := NewAuth(key)

	ctx := context.Background()

	token, err := auth.SignIn(ctx, "u", "p")
	parsed, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	assert.NoError(t, err, "token parse should not cause an error")
	s, err := parsed.Claims.GetExpirationTime()
	assert.NoError(t, err, "should get no error")

	if s.After(time.Now().Add(time.Hour)) {
		t.Errorf("expiration time should be no longer than 1 hour")
	}
}
