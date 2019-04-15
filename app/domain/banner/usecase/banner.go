package usecase

import (
	"crucial/banner/app/domain/banner/entity"
	"crucial/banner/app/domain/banner/repository"
	"crucial/banner/app/http/models"
)

// BannerUsecase - сценарии баннеров
type BannerUsecase struct {
	repository repository.BannerRepository
}

// NewBannerUsecase - создание объекта BannerUsecase
func NewBannerUsecase(repository repository.BannerRepository) BannerUsecase {
	return BannerUsecase{
		repository: repository,
	}
}

// Create - создание баннера
func (this *BannerUsecase) Create(request models.SaveBannerRequest) error {
	var err error
	banner := entity.Banner{
		Title:  request.Title,
		Brand:  request.Brand,
		Size:   request.Size,
		Active: request.Active,
	}

	err = this.repository.Save(&banner)
	if err != nil {
		return err
	}

	return nil
}

// Search - поиск баннеров по определенным критерия
func (this *BannerUsecase) Search(search *entity.SearchBanner) ([]entity.Banner, error) {
	var (
		banners []entity.Banner
		err     error
	)

	banners, err = this.repository.Search(search)
	return banners, err
}

// Destroy - удаления баннера
func (this *BannerUsecase) Destroy(id int) error {
	var err error

	err = this.repository.DeleteByID(id)
	if err != nil {
		return err
	}

	return nil
}
