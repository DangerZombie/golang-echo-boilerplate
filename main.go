//  Echo Boilerplate API:
//   version: 1.0
//   title: Echo Framework Boilerplate
//  Schemes: http, https
//  BasePath: /development/api/v1
//    Consumes:
//    - application/json
//  Produces:
//  - application/json
//  - text/html; charset=utf-8
//  securityDefinitions:
//   Bearer:
//    type: apiKey
//    name: Authorization
//    in: header
// swagger:meta
package main

import (
	"go-echo/initialization"
	"net/http"
	"strings"

	"github.com/labstack/echo"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var ActionIndex = func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("from action index"))
}

var ActionHome = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("from action home"))
	},
)

var ActionAbout = echo.WrapHandler(
	http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("from action about"))
		},
	),
)

func main() {
	viper.SetConfigFile("yaml")
	profile := "dev"

	var configFileName []string
	configFileName = append(configFileName, "config-")
	configFileName = append(configFileName, profile)
	viper.SetConfigName(strings.Join(configFileName, ""))
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	viper.AutomaticEnv()

	var logger *zap.Logger
	logOutput := viper.GetString("server.log-output")
	if logOutput == "file" {
		filename := viper.GetString("server.output-file-path")
		logger, _ = initialization.NewZapLogger(filename)
	} else {
		logger, _ = initialization.NewZapLogger("")
	}

	// init connection to DB
	db, err := initialization.DbInit()
	if err != nil {
		panic(err.Error())
	}

	logger.Info("connection db success")

	initialization.ServerInit(logger, db)
}
