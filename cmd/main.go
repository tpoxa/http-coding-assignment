package main

import (
	"github.com/qredo-external/go-maksym-trofimenko/cmd/cli"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"log"
)

func main() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		log.Fatal("fatal error config file", zap.Error(err))
	}
	err = cli.Execute()
	if err != nil && err.Error() != "" {
		log.Fatalln(err)
	}
}
