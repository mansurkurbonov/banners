package db

import (
	"crucial/banner/app/domain/banner/entity"
)

// PsqlBannerRepository - структура имплементирующая интерфейс BannerRepository
type PsqlBannerRepository struct{}

// NewPsqlBannerRepository - новый объект
func NewPsqlBannerRepository() *PsqlBannerRepository {
	return &PsqlBannerRepository{}
}

// Save - сохранения баннера в таблицу banners
func (this *PsqlBannerRepository) Save(banner *entity.Banner) error {
	var (
		sql = `INSERT INTO banners(
				title,
				brand,
				size,
				active
			) VALUES($1, $2, $3, $4)`

		err error
	)

	result, err := conn.Exec(sql, banner.Title, banner.Brand, banner.Size, banner.Active)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil || affected == 0 {
		return ErrNoRowsAffected
	}

	return nil
}

// DeleteByID - удаление баннера по id
func (this *PsqlBannerRepository) DeleteByID(id int) error {
	var (
		sql = `DELETE FROM banners where id = $1`
		err error
	)

	result, err := conn.Exec(sql, id)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil || affected == 0 {
		return ErrNoRowsAffected
	}

	return nil
}

// GetByID - возврашает баннер по id
func (this *PsqlBannerRepository) GetByID(id string) (*entity.Banner, error) {
	var (
		sql    = `SELECT id, title, brand, size, active FROM banners`
		err    error
		banner entity.Banner
	)

	err = conn.QueryRow(sql, id).Scan(
		&banner.ID,
		&banner.Title,
		&banner.Brand,
		&banner.Size,
		&banner.Active,
	)

	if err != nil {
		return &banner, err
	}

	return &banner, nil
}

// Search -  поиск баннера по критериям
func (this *PsqlBannerRepository) Search(search *entity.SearchBanner) ([]entity.Banner, error) {
	var (
		banner  entity.Banner
		banners []entity.Banner
		err     error
		sql     = "SELECT id, title, brand, size, active FROM banners WHERE 1 = 1"
	)

	if search.Title != "" {
		sql += " AND title = '" + search.Title + "'"
	}
	if search.Brand != "" {
		sql += " AND brand = '" + search.Brand + "'"
	}
	if search.Size != "" {
		sql += " AND brand = '" + search.Size + "'"
	}

	rows, err := conn.Query(sql)
	if err != nil {
		return banners, err
	}

	for rows.Next() {
		err = rows.Scan(
			&banner.ID,
			&banner.Title,
			&banner.Brand,
			&banner.Size,
			&banner.Active,
		)

		if err != nil {
			return banners, err
		}

		banners = append(banners, banner)
	}

	if len(banners) == 0 {

		return banners, ErrNotFound
	}

	return banners, err

}
