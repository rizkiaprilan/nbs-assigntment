package controller

import (
	"backend-nbs/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"net/http"
)

func GetListEmployeePerformance(ctx echo.Context) error {
	//GET DATA FROM JWT
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	isAdmin := claims["isAdmin"]

	listEmployeePerformance := &models.EmployeePerformance{}
	if isAdmin == true {
		*listEmployeePerformance = models.GetEmployeePerformance()
		return ctx.JSON(http.StatusOK, models.ConstructWebResponse(http.StatusOK, "success", listEmployeePerformance))
	} else {
		return ctx.JSON(http.StatusForbidden, models.ConstructWebResponse(http.StatusForbidden, "Sorry, You not have access", listEmployeePerformance))
	}
}
