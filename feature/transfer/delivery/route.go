package delivery

import (
	"portal/config"
	"portal/domain"
	"portal/feature/common"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RouteTransfer(e *echo.Echo, bc domain.TransferHandler) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))
	e.POST("/transfer", bc.InsertTransfer(), middleware.JWTWithConfig(common.UseJWT([]byte(config.SECRET))))

}
