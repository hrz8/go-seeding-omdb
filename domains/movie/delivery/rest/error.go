package rest

import (
	"errors"
	"net/http"

	MovieError "github.com/hrz8/go-seeding-omdb/domains/movie/error"
	"github.com/hrz8/go-seeding-omdb/utils"
)

type (
	RestErrorInterface interface {
		Throw(ctx *utils.CustomContext, domainErr error, dataErr error) error
	}

	restErrorImpl struct {
		prefix string
	}
)

func (i *restErrorImpl) Throw(ctx *utils.CustomContext, domainErr error, dataErr error) error {
	if errors.Is(domainErr, MovieError.List.Err) {
		status := uint16(MovieError.List.Status)
		return ctx.ErrorResponse(
			map[string]interface{}{
				"reason": dataErr.Error(),
			},
			domainErr.Error(),
			status,
			i.prefix+"-001",
			nil,
		)
	}
	if errors.Is(domainErr, MovieError.Detail.Err) {
		status := uint16(MovieError.Detail.Status)
		return ctx.ErrorResponse(
			map[string]interface{}{
				"reason": dataErr.Error(),
			},
			domainErr.Error(),
			status,
			i.prefix+"-002",
			nil,
		)
	}
	return ctx.ErrorResponse(
		nil,
		"Internal Server Error",
		http.StatusInternalServerError,
		i.prefix+"-500",
		nil,
	)
}

func NewMovieError() RestErrorInterface {
	return &restErrorImpl{
		prefix: "MOVIE",
	}
}
