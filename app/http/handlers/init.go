package handlers

import (
	"crucial/banner/app/domain/banner/repository"
	"crucial/banner/app/domain/banner/usecase"
)

// BannerHandler - структура хендлнров баннера
type BannerHandler struct {
	repository repository.BannerRepository
	usecase    usecase.BannerUsecase
}

// NewBannerHandler - создание обеъкта хендлера баннеров
func NewBannerHandler(repository repository.BannerRepository, usecase usecase.BannerUsecase) BannerHandler {
	return BannerHandler{
		repository: repository,
		usecase:    usecase,
	}
}
