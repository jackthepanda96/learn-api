package model

import "gorm.io/gorm"

type BarangModel struct {
	gorm.Model
	NamaBarang string
	Stok       int
	Harga      int
	UserID     uint
}
