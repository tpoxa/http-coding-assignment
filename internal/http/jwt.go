package http

import (
	"context"
	"errors"
	"fmt"
	"github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/labstack/echo/v4"
	"github.com/qredo-external/go-maksym-trofimenko/api"
	"github.com/qredo-external/go-maksym-trofimenko/internal/auth"
	"net/http"
	"strings"
)

var (
	ErrNoAuthHeader      = errors.New("authorization header is missing")
	ErrInvalidAuthHeader = errors.New("authorization header is malformed")
)

// CreateMiddleware Creates echo middleware functions using Auth adaptor
// Utilises openapi information to validate if security measures are relevant
func CreateMiddleware(v auth.IAuth) (echo.MiddlewareFunc, error) {
	spec, err := api.GetSwagger()
	if err != nil {
		return nil, fmt.Errorf("loading spec: %w", err)
	}
	return middleware.OapiRequestValidatorWithOptions(spec,
		&middleware.Options{
			Options: openapi3filter.Options{
				AuthenticationFunc: newAuthenticator(v),
			},
		}), nil
}

func newAuthenticator(v auth.IAuth) openapi3filter.AuthenticationFunc {
	return func(ctx context.Context, input *openapi3filter.AuthenticationInput) error {
		return authenticate(v, ctx, input)
	}
}

func authenticate(v auth.IAuth, ctx context.Context, input *openapi3filter.AuthenticationInput) error {
	if input.SecuritySchemeName != "bearerAuth" {
		return fmt.Errorf("security scheme %s is not supported", input.SecuritySchemeName)
	}
	// Now, we need to get the token string from the request, to match the request expectations
	// against request contents.
	jws, err := getJWSFromRequest(input.RequestValidationInput.Request)
	if err != nil {
		return fmt.Errorf("getting jws: %w", err)
	}
	// if the JWS is valid, we have a JWT, which will contain a bunch of claims.
	err = v.Validate(ctx, jws)
	if err != nil {
		return fmt.Errorf("validating JWS: %w", err)
	}
	return nil
}

func getJWSFromRequest(req *http.Request) (string, error) {
	authHdr := req.Header.Get("Authorization")
	// Check for the Authorization header.
	if authHdr == "" {
		return "", ErrNoAuthHeader
	}
	// We expect a header value of the form "Bearer <token>", with 1 space after
	// Bearer, per spec.
	prefix := "Bearer "
	if !strings.HasPrefix(authHdr, prefix) {
		return "", ErrInvalidAuthHeader
	}
	return strings.TrimPrefix(authHdr, prefix), nil
}
