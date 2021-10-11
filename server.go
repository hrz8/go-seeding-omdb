package main

import (
	"fmt"
	"log"
	"net"

	Config "github.com/hrz8/go-seeding-omdb/config"
	Database "github.com/hrz8/go-seeding-omdb/database"
	LogRequestRepository "github.com/hrz8/go-seeding-omdb/domains/log_request/repository"
	LogRequestUsecase "github.com/hrz8/go-seeding-omdb/domains/log_request/usecase"
	MovieGRPC "github.com/hrz8/go-seeding-omdb/domains/movie/delivery/grpc"
	MovieRest "github.com/hrz8/go-seeding-omdb/domains/movie/delivery/rest"
	MovieRepository "github.com/hrz8/go-seeding-omdb/domains/movie/repository"
	MovieUsecase "github.com/hrz8/go-seeding-omdb/domains/movie/usecase"
	"github.com/hrz8/go-seeding-omdb/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"google.golang.org/grpc"
)

func main() {
	appConfig := Config.NewConfig()
	mysql := Database.NewMysql(appConfig)
	mysqlSess := mysql.Connect()

	// usecase and repos
	movieRepo := MovieRepository.NewRepository()
	movieUsecase := MovieUsecase.NewUsecase(movieRepo)
	logRequestRepo := LogRequestRepository.NewRepository(mysqlSess)
	logRequestUsecase := LogRequestUsecase.NewUsecase(logRequestRepo)

	// REST
	restServer := echo.New()
	restServer.Validator = utils.NewValidator()
	restServer.Pre(middleware.RemoveTrailingSlash())
	restServer.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			cc := &utils.CustomContext{
				Context:   ctx,
				MysqlSess: mysqlSess,
				AppConfig: appConfig,
			}
			return next(cc)
		}
	})
	movieRest := MovieRest.NewRest(movieUsecase, logRequestUsecase)
	MovieRest.RegisterEndpoint(restServer, movieRest)
	// REST serve
	restServer.Logger.Fatal(restServer.Start(fmt.Sprintf(":%d", appConfig.SERVICE.RESTPORT)))

	// GRPC
	grpcServer := grpc.NewServer()
	var movieGrpc MovieGRPC.MoviesServer
	MovieGRPC.RegisterServer(grpcServer, movieGrpc)
	// GRPC serve
	l, err := net.Listen("tcp", fmt.Sprintf("%d", appConfig.SERVICE.GRPCPORT))
	if err != nil {
		log.Fatalf("could not listen to %d: %v", appConfig.SERVICE.GRPCPORT, err)
	}
	log.Fatal(grpcServer.Serve(l))
}
