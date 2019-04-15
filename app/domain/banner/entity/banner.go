package entity

import "errors"

var (
	errEmptyTitle = errors.New("Title не должно быть пустым")
	errEmptyBrand = errors.New("Brand не должно быть пустым")
	errEmptySize  = errors.New("Size не должно быть пустым")
)

// Banner - сущность баннера
type Banner struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Brand  string `json:"brand"`
	Size   string `json:"size"`
	Active bool   `json:"active"`
}

// SearchBanner - критерии поиска баннера
type SearchBanner struct {
	Title string `json:"title"`
	Brand string `json:"brand"`
	Size  string `json:"size"`
}

// Validate - валидация данных
func (this *Banner) Validate() error {
	if this.Title == "" {
		return errEmptyTitle
	}
	if this.Brand == "" {
		return errEmptyBrand
	}
	if this.Size == "" {
		return errEmptySize
	}

	return nil
}
