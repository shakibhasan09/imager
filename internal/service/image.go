package service

import (
	"context"

	"github.com/shakibhasan09/imager/internal/db/models"
)

type ImageRepository interface {
	Get(ctx context.Context) (*[]models.Image, error)
	GetByUuid(ctx context.Context, uuid string) (*models.Image, error)
	Update(ctx context.Context, uuid string) error
	Delete(ctx context.Context, uuid string) error
	Create(ctx context.Context, image *models.Image) error
}

type ImageService struct {
	db ImageRepository
}

func NewImageService(db ImageRepository) *ImageService {
	return &ImageService{
		db: db,
	}
}

func (s *ImageService) Upload() {

}
