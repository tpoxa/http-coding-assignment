package cli

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/qredo-external/go-maksym-trofimenko/api"
	"github.com/qredo-external/go-maksym-trofimenko/internal/auth"
	dataanalyser "github.com/qredo-external/go-maksym-trofimenko/internal/data-analyser"
	"github.com/qredo-external/go-maksym-trofimenko/internal/hasher"
	"github.com/qredo-external/go-maksym-trofimenko/internal/http"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

var (
	rootCmd = &cobra.Command{
		Use:   "app",
		Short: "app – coding assignment ",
		Long:  ``,
		Run:   runServer,
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func runServer(_ *cobra.Command, _ []string) {

	e := echo.New()
	//a := auth.NewAuth()
	a := auth.NewAuth([]byte(viper.GetString("jwt_private_key")))
	// Create middleware for validating tokens.
	authMiddleware, err := http.CreateMiddleware(a)
	if err != nil {
		log.Fatalln("error creating middleware:", err)
	}

	e.HideBanner = true
	e.Use(middleware.Logger())
	e.Use(authMiddleware)

	analyser := dataanalyser.NewAnalyser()
	h := hasher.NewSha256()

	svr := http.NewServer(analyser, a, h)

	api.RegisterHandlers(e, svr)
	e.Logger.Fatal(e.Start(viper.GetString("listen_address")))
}
