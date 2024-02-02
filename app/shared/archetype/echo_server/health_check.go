package echo_server

import (
	"te-eme-backend/app/shared/archetype/container"
	"net/http"

	"github.com/labstack/echo/v4"
)

func init() {
	container.InjectInboundAdapter(func() error {
		Echo().GET("/health", func(c echo.Context) error {
			return c.String(http.StatusOK, "UP")
		})
		return nil
	})
}
