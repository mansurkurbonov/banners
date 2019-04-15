package repository

import "crucial/banner/app/domain/banner/entity"

// BannerRepository - хранилище для сущности баннер
type BannerRepository interface {
	Save(*entity.Banner) error
	DeleteByID(int) error
	GetByID(string) (*entity.Banner, error)
	Search(*entity.SearchBanner) ([]entity.Banner, error)
}
