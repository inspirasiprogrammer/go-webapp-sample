package main

import (
	"github.com/labstack/echo/v4"
	"github.com/ybkuroki/go-webapp-sample/config"
	"github.com/ybkuroki/go-webapp-sample/logger"
	"github.com/ybkuroki/go-webapp-sample/middleware"
	"github.com/ybkuroki/go-webapp-sample/migration"
	"github.com/ybkuroki/go-webapp-sample/repository"
	"github.com/ybkuroki/go-webapp-sample/router"
)

func main() {
	e := echo.New()

	config.Load()
	logger.InitLogger()
	middleware.InitLoggerMiddleware(e)
	logger.GetZapLogger().Infof("Loaded this configuration : application." + *config.GetEnv() + ".yml")

	repository.InitDB()
	db := repository.GetDB()

	migration.CreateDatabase(config.GetConfig())
	migration.InitMasterData(config.GetConfig())

	router.Init(e, config.GetConfig())
	middleware.InitSessionMiddleware(e, config.GetConfig())
	if err := e.Start(":8080"); err != nil {
		logger.GetZapLogger().Errorf(err.Error())
	}

	defer db.Close()
}
