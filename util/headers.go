package util

import (
	"github.com/Leonardo-Antonio/api.process-vtex/entity"
	"github.com/labstack/echo/v4"
)

func GetHeaders(c echo.Context) (headers entity.Header) {
	headers.Token = c.Request().Header.Get("Token")
	headers.Key = c.Request().Header.Get("Key")
	headers.Store = c.Request().Header.Get("Store")
	return
}
