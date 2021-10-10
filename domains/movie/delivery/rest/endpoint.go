package rest

import (
	"reflect"

	"github.com/hrz8/go-seeding-omdb/models"
	"github.com/hrz8/go-seeding-omdb/utils"
	"github.com/labstack/echo/v4"
)

func RegisterEndpoint(e *echo.Echo, rest RestInterface) {
	e.GET("/api/v1/movie", rest.List, utils.ValidatorMiddleware(reflect.TypeOf(models.MoviePayloadList{})))
	e.GET("/api/v1/movie/:id", rest.Detail, utils.ValidatorMiddleware(reflect.TypeOf(models.MoviePayloadDetail{})))
}
