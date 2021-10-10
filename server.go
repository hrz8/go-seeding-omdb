package main

import (
	"fmt"

	Config "github.com/hrz8/go-seeding-omdb/config"
	Database "github.com/hrz8/go-seeding-omdb/database"
	MovieRest "github.com/hrz8/go-seeding-omdb/domains/movie/delivery/rest"
	MovieRepository "github.com/hrz8/go-seeding-omdb/domains/movie/repository"
	MovieUsecase "github.com/hrz8/go-seeding-omdb/domains/movie/usecase"
	"github.com/hrz8/go-seeding-omdb/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	appConfig := Config.NewConfig()
	mysql := Database.NewMysql(appConfig)
	mysqlSess := mysql.Connect()

	e := echo.New()
	e.Validator = utils.NewValidator()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			cc := &utils.CustomContext{
				Context:   ctx,
				MysqlSess: mysqlSess,
				AppConfig: appConfig,
			}
			return next(cc)
		}
	})

	movieRepo := MovieRepository.NewRepository()
	movieUsecase := MovieUsecase.NewUsecase(movieRepo)
	movieRest := MovieRest.NewRest(movieUsecase)
	MovieRest.RegisterEndpoint(e, movieRest)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", appConfig.SERVICE.PORT)))
}
