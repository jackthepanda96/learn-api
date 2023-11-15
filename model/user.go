package model

import "gorm.io/gorm"

type UserModel struct {
	gorm.Model
	Nama     string
	Hp       string
	Password string
	Barangs  []BarangModel `gorm:"foreignKey:UserID"`
}

type UserQuery struct {
	DB *gorm.DB
}

// func (uq *UserQuery) Register(newUser UserModel) bool {
// 	if err := uq.DB.Create(&newUser).Error; err != nil {
// 		return false
// 	}
// }

func (uq *UserQuery) Register(newUser UserModel) (UserModel, error) {
	if err := uq.DB.Create(&newUser).Error; err != nil {
		return UserModel{}, err
	}

	return newUser, nil
}

func (uq *UserQuery) Login(hp string, password string) (UserModel, error) {
	var result = new(UserModel)

	if err := uq.DB.Where("hp = ? AND password = ?", hp, password).First(result).Error; err != nil {
		return UserModel{}, err
	}

	return *result, nil
}
