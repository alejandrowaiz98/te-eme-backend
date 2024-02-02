package controller

import (
	"te-eme-backend/app/shared/archetype/container"
	einar "te-eme-backend/app/shared/archetype/echo_server"

	"net/http"

	"github.com/labstack/echo/v4"
)

func init() {
	container.InjectInboundAdapter(func() error {
		einar.Echo().POST("/api/insert_your_pattern_here", postUser)
		return nil
	})
}

func postUser(c echo.Context) error {

	//LOGIC GOES HERE!!!

	return c.JSON(http.StatusOK, "insert_your_custom_response")
}
