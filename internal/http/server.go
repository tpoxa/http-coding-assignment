package http

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/qredo-external/go-maksym-trofimenko/api"
	"github.com/qredo-external/go-maksym-trofimenko/internal/auth"
	dataanalyser "github.com/qredo-external/go-maksym-trofimenko/internal/data-analyser"
	"github.com/qredo-external/go-maksym-trofimenko/internal/hasher"
	"net/http"
)

type Server struct {
	dataAnalyser dataanalyser.IDataAnalyser
	auth         auth.IAuth
	hasher       hasher.Hasher
}

func NewServer(dataAnalyser dataanalyser.IDataAnalyser, auth auth.IAuth, hasher hasher.Hasher) *Server {
	return &Server{dataAnalyser: dataAnalyser, auth: auth, hasher: hasher}
}

func (s Server) Authenticate(c echo.Context) error {
	var req api.AuthRequest
	err := c.Bind(&req)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	tokenStr, err := s.auth.SignIn(c.Request().Context(), req.Username, req.Password)
	if err != nil {
		if err == auth.ErrInvalidCredentials {
			return c.String(http.StatusUnauthorized, "invalid credentials")
		}
		// 500
		return err
	}
	return c.String(http.StatusOK, tokenStr)
}

func (s Server) CheckHealth(c echo.Context) error {
	return c.String(http.StatusOK, "working")
}

func (s Server) Sum(c echo.Context) error {
	sum, err := s.dataAnalyser.CalcSum(c.Request().Context(), c.Request().Body)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, s.hasher.Hash(c.Request().Context(), fmt.Sprintf("%v", sum)))
}

// static validation helps also helps to easily find interface we implement
var _ api.ServerInterface = (*Server)(nil)
