package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string `json:"nama"`
	Hp   string `json:"handphone"`
}

func RegisterHandler(conn *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = new(User)

		err := c.Bind(input)

		if err != nil {
			c.Echo().Logger.Error("input error :", err.Error())
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "input data kurang tepat",
				"data":    nil,
			})
		}

		err = conn.Create(input).Error

		if err != nil {
			c.Echo().Logger.Error("database error :", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"message": "terdapat permasalahan pada pengolahan data",
				"data":    nil,
			})
		}

		return c.JSON(http.StatusCreated, map[string]any{
			"message": "success",
			"data":    input,
		})
	}
}

func main() {
	e := echo.New()

	dsn := "root:@tcp(localhost:3306)/orm_19?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&User{})

	if err != nil {
		e.Logger.Fatal("cannot connect to DB.", err.Error())
		return
	}

	e.POST("/users", RegisterHandler(db))

	e.Logger.Fatal(e.Start(":8000"))
}
