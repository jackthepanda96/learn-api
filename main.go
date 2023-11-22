package main

import (
	"19api/config"
	bh "19api/features/barang/handler"
	br "19api/features/barang/repository"
	bs "19api/features/barang/services"
	uh "19api/features/users/handler"
	ur "19api/features/users/repository"
	us "19api/features/users/services"
	"19api/helper/enkrip"
	"19api/routes"

	"19api/utils/database"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"github.com/go-playground/validator/v10"
)

type User struct {
	gorm.Model
	Name string `json:"nama" validate:"required"`
	Hp   string `json:"handphone" gorm:"unique"`
}

func RegisterHandler(conn *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// BACA INPUT
		var input = new(User)
		err := c.Bind(input)

		if err != nil {
			c.Echo().Logger.Error("input error :", err.Error())
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "input data kurang tepat",
				"data":    nil,
			})
		}
		// VALIDASI INPUT
		validate := validator.New(validator.WithRequiredStructEnabled())

		err = validate.Struct(input)

		if err != nil {
			c.Echo().Logger.Error("input error :", err.Error())
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": err.Error(),
				"data":    nil,
			})
		}

		// SIMPAN KE DB
		// var cekData = new(User)

		// err = conn.Where("hp = ?", input.Hp).First(cekData).Error
		// if err != nil || cekData.ID != 0 {
		// 	c.Echo().Logger.Error("database error : duplicate entry")
		// 	return c.JSON(http.StatusConflict, map[string]any{
		// 		"message": "duplicate nomor hp",
		// 		"data":    nil,
		// 	})
		// }

		err = conn.Create(input).Error

		if err != nil {
			c.Echo().Logger.Error("database error :", err.Error())
			if strings.Contains(err.Error(), "Duplicate") {
				return c.JSON(http.StatusConflict, map[string]any{
					"message": "duplicate nomor hp",
					"data":    nil,
				})
			}
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"message": "terdapat permasalahan pada pengolahan data",
				"data":    nil,
			})
		}

		// RETURN RESULT
		return c.JSON(http.StatusCreated, map[string]any{
			"message": "success",
			"data":    input,
		})
	}
}

func main() {
	e := echo.New()

	cfg := config.InitConfig()

	if cfg == nil {
		e.Logger.Fatal("tidak bisa start karena ENV error")
		return
	}

	db, err := database.InitMySQL(*cfg)

	if err != nil {
		e.Logger.Fatal("tidak bisa start karena DB error:", err.Error())
		return
	}

	db.AutoMigrate(&ur.UserModel{}, &br.BarangModel{})

	// m := model.UserQuery{DB: db}
	// userController := user.UserController{Model: m}

	// bm := model.BarangQuery{DB: db}
	// barangController := barang.BarangController{Model: bm}
	hash := enkrip.New()
	userRepo := ur.New(db)
	userService := us.New(userRepo, hash)
	userHandler := uh.New(userService)

	barangRepo := br.New(db)
	barabgService := bs.New(barangRepo)
	barangHandler := bh.New(barabgService)

	routes.InitRoute(e, userHandler, barangHandler)

	e.Logger.Fatal(e.Start(":8000"))
}
