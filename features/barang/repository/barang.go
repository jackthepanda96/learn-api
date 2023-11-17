package repository

import (
	"19api/features/barang"

	"gorm.io/gorm"
)

type BarangModel struct {
	gorm.Model
	NamaBarang string
	Stok       int
	Harga      int
	UserID     uint
}

type barangQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) barang.Repo {
	return &barangQuery{
		db: db,
	}
}

func (bq *barangQuery) InsertBarang(userID uint, newBarang barang.Barang) (barang.Barang, error) {
	var inputData = new(BarangModel)
	inputData.UserID = userID
	inputData.NamaBarang = newBarang.Nama
	inputData.Harga = newBarang.Harga
	inputData.Stok = newBarang.Stok

	if err := bq.db.Create(&inputData).Error; err != nil {
		return barang.Barang{}, err
	}

	newBarang.ID = inputData.ID

	return newBarang, nil
}
