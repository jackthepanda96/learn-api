package model

import "gorm.io/gorm"

type BarangModel struct {
	gorm.Model
	NamaBarang string
	Stok       int
	Harga      int
	UserID     uint
}

type BarangQuery struct {
	DB *gorm.DB
}

func (bq *BarangQuery) AddBarang(newBarang BarangModel) (BarangModel, error) {
	if err := bq.DB.Create(&newBarang).Error; err != nil {
		return BarangModel{}, err
	}

	return newBarang, nil
}
