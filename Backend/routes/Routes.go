package routes

import (
	"backend-nbs/helpers"
	"github.com/labstack/echo"
)

func Controller() *echo.Echo {
	router := echo.New()

	//validation
	helpers.Validation(router)

	// routes
	AuthRoutes(router)
	AdminRoutes(router)

	return router
}
