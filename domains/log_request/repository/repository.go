package repository

import (
	"github.com/hrz8/go-seeding-omdb/models"
	"gorm.io/gorm"
)

type (
	RepositoryInterface interface {
		Create(trx *gorm.DB, logRequest *models.LogRequest) error
	}

	impl struct {
		db *gorm.DB
	}
)

func (i *impl) Create(trx *gorm.DB, logRequest *models.LogRequest) error {
	if trx == nil {
		trx = i.db
	}
	if err := trx.Create(&logRequest).Error; err != nil {
		return err
	}
	return nil
}

func NewRepository(db *gorm.DB) RepositoryInterface {
	db.AutoMigrate(&models.LogRequest{})
	return &impl{
		db: db,
	}
}
