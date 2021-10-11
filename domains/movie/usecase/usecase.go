package usecase

import (
	"github.com/hrz8/go-seeding-omdb/domains/movie/repository"
	"github.com/hrz8/go-seeding-omdb/models"
	Context "github.com/hrz8/go-seeding-omdb/utils"
)

type (
	UsecaseInterface interface {
		List(ctx *Context.CustomContext, payload *models.MoviePayloadList) (*[]models.Movie, *int, error)
		Detail(ctx *Context.CustomContext, id *string) (*models.Movie, error)
	}

	impl struct {
		repository repository.RepositoryInterface
	}
)

func (i *impl) List(ctx *Context.CustomContext, payload *models.MoviePayloadList) (*[]models.Movie, *int, error) {
	apiKey := &ctx.AppConfig.SERVICE.APIKEY
	result, total, err := i.repository.List(apiKey, payload)
	return result, total, err
}

func (i *impl) Detail(ctx *Context.CustomContext, id *string) (*models.Movie, error) {
	apiKey := &ctx.AppConfig.SERVICE.APIKEY
	result, err := i.repository.Detail(apiKey, id)
	return result, err
}

func NewUsecase(r repository.RepositoryInterface) UsecaseInterface {
	return &impl{
		repository: r,
	}
}
