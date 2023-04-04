package http

import (
	"bytes"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/qredo-external/go-maksym-trofimenko/api"
	"github.com/qredo-external/go-maksym-trofimenko/internal/auth"
	authmock "github.com/qredo-external/go-maksym-trofimenko/internal/auth/mocks"
	dataanalysermock "github.com/qredo-external/go-maksym-trofimenko/internal/data-analyser/mocks"
	hashermock "github.com/qredo-external/go-maksym-trofimenko/internal/hasher/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strings"

	"testing"
)

func initEcho() (*echo.Echo, *authmock.IAuth, *hashermock.Hasher, *dataanalysermock.IDataAnalyser) {
	// mocks
	authComponent := new(authmock.IAuth)
	hashComponent := new(hashermock.Hasher)
	dataAnalyser := new(dataanalysermock.IDataAnalyser)

	authMiddleware, _ := CreateMiddleware(authComponent)

	e := echo.New()
	e.Debug = true
	e.Use(authMiddleware)
	api.RegisterHandlers(e, NewServer(dataAnalyser, authComponent, hashComponent))
	return e, authComponent, hashComponent, dataAnalyser
}

func TestSumUnauthorisedAccessForbidden(t *testing.T) {

	e, _, _, _ := initEcho()
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("POST", "/sum", nil)
	e.ServeHTTP(w, r)
	assert.Equal(t, http.StatusForbidden, w.Code)
}

func TestSignInBadRequest(t *testing.T) {
	e, _, _, _ := initEcho()

	//e, authCmp, hashCmp, analyserCmp := initEcho()
	for _, c := range []string{
		"dsds",
		"",
		"{}",
		`{"username": ""}`,
		`{"password": ""}`,
	} {

		r, err := http.NewRequest("POST", "/auth", strings.NewReader(c))
		r.Header.Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
		assert.NoError(t, err)

		w := httptest.NewRecorder()
		e.ServeHTTP(w, r)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	}
}

func TestSignInInvalidCredentialsRequest(t *testing.T) {
	e, authMock, _, _ := initEcho()

	authMock.On("SignIn", mock.Anything, mock.Anything, mock.Anything).Return("", auth.ErrInvalidCredentials).Times(3)
	//e, authCmp, hashCmp, analyserCmp := initEcho()
	for _, c := range []string{
		`{"username": "test", "password": ""}`,
		`{"username": "", "password": "ddd"}`,
		`{"username": "wwww", "password": "ddd"}`,
	} {

		r, err := http.NewRequest("POST", "/auth", strings.NewReader(c))
		r.Header.Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
		assert.NoError(t, err)

		w := httptest.NewRecorder()
		e.ServeHTTP(w, r)
		assert.Equal(t, http.StatusUnauthorized, w.Code)
	}
}

// TestSignInSuccessRequest test success signIn returns auth token in a response
func TestSignInSuccessRequest(t *testing.T) {
	e, authMock, _, _ := initEcho()

	var (
		username = "u1"
		password = "qwerty"
		token    = "tok"
	)

	authMock.On("SignIn", mock.Anything, username, password).Return(token, nil).Times(1)
	//e, authCmp, hashCmp, analyserCmp := initEcho()

	r, err := http.NewRequest("POST", "/auth", strings.NewReader(fmt.Sprintf(`{"username": "%s", "password": "%s"}`, username, password)))
	r.Header.Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	e.ServeHTTP(w, r)
	assert.Equal(t, http.StatusOK, w.Code)

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(w.Body)
	assert.NoError(t, err)
	assert.Equal(t, buf.String(), token)

}
