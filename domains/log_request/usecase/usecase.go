package usecase

import (
	"github.com/hrz8/go-seeding-omdb/domains/log_request/repository"
	"github.com/hrz8/go-seeding-omdb/models"
	"github.com/hrz8/go-seeding-omdb/utils"
)

type (
	UsecaseInterface interface {
		Create(ctx *utils.CustomContext, logRequest *models.LogRequest) error
	}

	impl struct {
		repository repository.RepositoryInterface
	}
)

func (i *impl) Create(ctx *utils.CustomContext, logRequest *models.LogRequest) error {
	return i.repository.Create(nil, logRequest)
}

func NewUsecase(r repository.RepositoryInterface) UsecaseInterface {
	return &impl{
		repository: r,
	}
}
