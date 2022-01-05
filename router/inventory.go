package router

import (
	"github.com/Leonardo-Antonio/api.process-vtex/handler"
	"github.com/labstack/echo/v4"
)

func Inventory(app *echo.Echo) {
	controller := new(handler.Inventory)
	group := app.Group("/api/v1/inventory")
	group.PUT("", controller.SetBySkuAndWarehouse)
}
