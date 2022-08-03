package delivery

import (
	"altaproject2/config"
	"altaproject2/domain"
	"altaproject2/features/cart/delivery/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RouteCart(e *echo.Echo, ph domain.CartHandler) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.PATCH, echo.OPTIONS},
	}))
	e.Pre(middleware.RemoveTrailingSlash())

	e.PUT("/carts", ph.Update(), middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET))))
}
