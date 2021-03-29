package routes

import (
	controller "backend-nbs/controllers"

	"github.com/labstack/echo"
)

func AuthRoutes(router *echo.Echo) {
	router.POST("/register", controller.Register)
	router.POST("/login", controller.Login)
	router.POST("/logout", controller.Logout)
	router.GET("/account/activated/:token", controller.ActivateAccount)
	router.GET("/account/forget-password/:token", controller.ForgetPassword)
	router.GET("/send/email-verification/:email", controller.ResendEmailVerification)
	router.GET("/send/forget-password/:email", controller.ResendForgetPassword)
}
