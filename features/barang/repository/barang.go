package repository

import "gorm.io/gorm"

type BarangModel struct {
	gorm.Model
	NamaBarang string
	Stok       int
	Harga      int
	UserID     uint
}

type BarangQuery struct {
	db *gorm.DB
}

func (bq *BarangQuery) AddBarang(newBarang BarangModel) (BarangModel, error) {
	if err := bq.db.Create(&newBarang).Error; err != nil {
		return BarangModel{}, err
	}

	return newBarang, nil
}
