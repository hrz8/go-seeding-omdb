package main

import (
	"fmt"

	Config "github.com/hrz8/go-seeding-omdb/config"
	Context "github.com/hrz8/go-seeding-omdb/context"
	Database "github.com/hrz8/go-seeding-omdb/database"
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
			cc := Context.CustomContext{
				Context:   ctx,
				MysqlSess: mysqlSess,
				AppConfig: appConfig,
			}
			return next(cc)
		}
	})

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", appConfig.SERVICE.PORT)))
}
