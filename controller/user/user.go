package user

import (
	"19api/model"
	"19api/utils/jwt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	Model model.UserQuery
}

func (uc *UserController) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = new(UserRequest)
		if err := c.Bind(input); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "input yang diberikan tidak sesuai",
			})
		}

		var inputProcess = new(model.UserModel)
		inputProcess.Nama = input.Nama
		inputProcess.Password = input.Password
		inputProcess.Hp = input.Hp

		result, err := uc.Model.Register(*inputProcess)

		if err != nil {
			c.Logger().Error("ERROR Register, explain:", err.Error())
			if strings.Contains(err.Error(), "duplicate") {
				return c.JSON(http.StatusBadRequest, map[string]any{
					"message": "data yang diinputkan sudah terdaftar ada sistem",
				})
			}
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"message": "terjadi permasalahan ketika memproses data",
			})
		}

		var response = new(UserResponse)
		response.Hp = result.Hp
		response.Nama = result.Nama
		response.ID = result.ID

		return c.JSON(http.StatusCreated, map[string]any{
			"message": "success create data",
			"data":    response,
		})
	}
}

func (uc *UserController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = new(LoginRequest)
		if err := c.Bind(input); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "input yang diberikan tidak sesuai",
			})
		}

		result, err := uc.Model.Login(input.Hp, input.Password)

		if err != nil {
			c.Logger().Error("ERROR Login, explain:", err.Error())
			if strings.Contains(err.Error(), "not found") {
				return c.JSON(http.StatusBadRequest, map[string]any{
					"message": "data yang diinputkan tidak ditemukan",
				})
			}
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"message": "terjadi permasalahan ketika memproses data",
			})
		}

		strToken, err := jwt.GenerateJWT()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"message": "terjadi permasalahan ketika mengenkripsi data",
			})
		}

		var response = new(LoginResponse)
		response.Hp = result.Hp
		response.Nama = result.Nama
		response.ID = result.ID
		response.Token = strToken

		return c.JSON(http.StatusOK, map[string]any{
			"message": "success create data",
			"data":    response,
		})
	}
}

func (uc *UserController) GetListUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, "return data")
	}
}
