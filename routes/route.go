package routes

import (
	"19api/features/barang"
	"19api/features/users"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoute(e *echo.Echo, uc users.Handler, bc barang.Handler) {
	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	routeUser(e, uc)
	routeBarang(e, bc)
}

func routeUser(e *echo.Echo, uc users.Handler) {
	e.POST("/users", uc.Register())
	e.POST("/login", uc.Login())
	// e.GET("/users", uc.GetListUser(), echojwt.JWT([]byte("$!1gnK3yyy!!!")))
}

func routeBarang(e *echo.Echo, bc barang.Handler) {
	e.POST("/barangs", bc.Add(), echojwt.JWT([]byte("$!1gnK3yyy!!!")))
}
