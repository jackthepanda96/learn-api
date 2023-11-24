package barang

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type Barang struct {
	ID    uint
	Nama  string
	Stok  int
	Harga int
}

type Handler interface {
	Add() echo.HandlerFunc
}

type Service interface {
	TambahBarang(token *jwt.Token, newBarang Barang) (Barang, error)
}

type Repo interface {
	InsertBarang(userID uint, newBarang Barang) (Barang, error)
}
