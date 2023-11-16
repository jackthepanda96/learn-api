package services

import (
	"19api/features/barang"
	"19api/utils/jwt"
	"errors"
	"strings"

	golangjwt "github.com/golang-jwt/jwt/v5"
)

type BarangServices struct {
	m barang.Repo
}

func New(model barang.Repo) barang.Service {
	return &BarangServices{
		m: model,
	}
}

func (bs *BarangServices) TambahBarang(token *golangjwt.Token, newBarang barang.Barang) (barang.Barang, error) {
	userID, err := jwt.ExtractToken(token)
	if err != nil {
		return barang.Barang{}, err
	}

	result, err := bs.m.InsertBarang(userID, newBarang)

	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return barang.Barang{}, errors.New("barang sudah pernah diinputkan")
		}
		return barang.Barang{}, errors.New("terjadi kesalahan pada server")
	}

	return result, nil
}
