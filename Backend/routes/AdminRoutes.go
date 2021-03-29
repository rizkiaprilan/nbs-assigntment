package routes

import (
	controller "backend-nbs/controllers"
	"backend-nbs/helpers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func AdminRoutes(router *echo.Echo) {
	router.GET("/list-employee-performance", controller.GetListEmployeePerformance, middleware.JWT([]byte(helpers.GetEnv("JWT_SIGNATURE_KEY"))))
}
