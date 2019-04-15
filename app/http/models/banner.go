package models

// SaveBannerRequest - структура данных для сохранения
type SaveBannerRequest struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Brand  string `json:"brand"`
	Size   string `json:"size"`
	Active bool   `json:"active"`
}

// Validate - проверка данных для сохранения
func (this *SaveBannerRequest) Validate() (bool, string) {
	if this.Title == "" {
		return false, "поле Title не должно быть пустым"
	} else if this.Brand == "" {
		return false, "поле Brand не должно быть пустым"
	} else if this.Size == "" {
		return false, "поле Title не должно быть пустым"
	}

	return true, ""
}

// SearchBannerRequest -
type SearchBannerRequest struct {
	Banner SaveBannerRequest
}
