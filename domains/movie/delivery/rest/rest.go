package rest

import (
	"errors"
	"strconv"

	MovieError "github.com/hrz8/go-seeding-omdb/domains/movie/error"
	"github.com/hrz8/go-seeding-omdb/domains/movie/usecase"
	"github.com/hrz8/go-seeding-omdb/models"
	"github.com/hrz8/go-seeding-omdb/utils"
	"github.com/labstack/echo/v4"
)

type (
	RestInterface interface {
		List(c echo.Context) error
		Detail(c echo.Context) error
	}

	impl struct {
		usecase  usecase.UsecaseInterface
		errorLib RestErrorInterface
	}
)

func (i *impl) List(c echo.Context) error {
	ctx := c.(*utils.CustomContext)
	pagination, err := strconv.Atoi(ctx.QueryParam("pagination"))
	if err != nil {
		return i.errorLib.Throw(ctx, MovieError.List.Err, errors.New("invalid pagination data type"))
	}
	payload := &models.MoviePayloadList{
		Pagination: pagination,
		Searchword: ctx.QueryParam("searchword"),
	}
	result, total, err := i.usecase.List(ctx, payload)
	if err != nil {
		return i.errorLib.Throw(ctx, MovieError.List.Err, err)
	}
	return ctx.SuccessResponse(
		result,
		"success fetch movies list",
		utils.ListMetaResponse{
			Count: len(*result),
			Total: *total,
		},
	)
}

func (i *impl) Detail(c echo.Context) error {
	ctx := c.(*utils.CustomContext)
	id := ctx.Param("id")
	result, err := i.usecase.Detail(ctx, &id)
	if err != nil {
		return i.errorLib.Throw(ctx, MovieError.Detail.Err, err)
	}
	return ctx.SuccessResponse(
		result,
		"success fetch movie detail",
		nil,
	)
}

func NewRest(u usecase.UsecaseInterface) RestInterface {
	return &impl{
		usecase:  u,
		errorLib: NewMovieError(),
	}
}
