package delivery

import (
	"portal/config"
	"portal/domain"
	"portal/feature/common"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RouteUser(e *echo.Echo, bc domain.UserHandler) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))
	e.POST("/register", bc.InsertUser())
	e.POST("/login", bc.LoginUser())
	e.GET("/my", bc.GetProfile(), middleware.JWTWithConfig(common.UseJWT([]byte(config.SECRET))))
	e.GET("/datausers", bc.GetAllUser())
	e.DELETE("/users", bc.DeleteUser(), middleware.JWTWithConfig(common.UseJWT([]byte(config.SECRET))))
	e.PUT("/profile", bc.UpdateUser(), middleware.JWTWithConfig(common.UseJWT([]byte(config.SECRET))))
}
