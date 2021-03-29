package controller

import (
	"backend-nbs/helpers"
	"backend-nbs/models"
	"encoding/base64"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/labstack/echo"
)

func Register(ctx echo.Context) error {
	var newEmployee models.EmployeeRegisterRequest

	err := ctx.Bind(&newEmployee)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.ConstructWebResponse(http.StatusBadRequest, err.Error(), nil))
	}

	//VALIDATE REQUESTF
	err = ctx.Validate(newEmployee)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.ConstructWebResponse(http.StatusBadRequest, err.Error(), nil))
	}

	getEmployee := models.GetEmployeeByEmail(newEmployee.Email)
	if getEmployee.Email != "" {
		return ctx.JSON(http.StatusBadRequest, models.ConstructWebResponse(http.StatusBadRequest, "Employee Already Exists", nil))
	}

	if newEmployee.ConfirmPassword != newEmployee.Password {
		return ctx.JSON(http.StatusBadRequest, models.ConstructWebResponse(http.StatusBadRequest, "Password and Confirm Password Not Match", nil))
	}

	//CONSTRUCT MEMBER AND STORE TO DB
	var employee models.Employee
	employee.Email = newEmployee.Email
	employee.Password = helpers.HashAndSalt([]byte(newEmployee.Password))
	employee.Fullname = newEmployee.Fullname
	employee.Phone = newEmployee.Phone
	employee.Id = uuid.NewString()

	err = models.CreateEmployee(employee)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.ConstructWebResponse(http.StatusBadRequest, err.Error(), nil))
	}

	var activeLink models.ActivateLink
	activeLink.Email = employee.Email
	activeLink.CreatedAt = time.Now().Unix()
	activeLink.Token = base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s&%s", activeLink.Email, strconv.FormatInt(activeLink.CreatedAt, 10))))

	linkActivate := helpers.GenerateVerficationLink(activeLink.Token)
	err = helpers.SendEmail(employee.Email, linkActivate)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.ConstructWebResponse(http.StatusBadRequest, err.Error(), nil))
	}

	return ctx.JSON(http.StatusOK, models.ConstructWebResponse(http.StatusOK, "SUCCESS", linkActivate))
}

func ResendEmailVerification(ctx echo.Context) error {
	email := ctx.Param("email")
	var activeLink models.ActivateLink
	activeLink.Email = email
	activeLink.CreatedAt = time.Now().Unix()
	activeLink.Token = base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s&%s", activeLink.Email, strconv.FormatInt(activeLink.CreatedAt, 10))))

	linkActivate := helpers.GenerateVerficationLink(activeLink.Token)
	err := helpers.SendEmail(email, linkActivate)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.ConstructWebResponse(http.StatusBadRequest, err.Error(), nil))
	}

	return ctx.JSON(http.StatusOK, models.ConstructWebResponse(http.StatusOK, "SUCCESS", linkActivate))
}

func ResendForgetPassword(ctx echo.Context) error {
	email := ctx.Param("email")
	var activeLink models.ActivateLink
	activeLink.Email = email
	activeLink.CreatedAt = time.Now().Unix()
	activeLink.Token = base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s&%s", activeLink.Email, strconv.FormatInt(activeLink.CreatedAt, 10))))

	linkChangePassword := helpers.GenerateChangePasswordLink(activeLink.Token)
	err := helpers.SendEmail(email, linkChangePassword)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.ConstructWebResponse(http.StatusBadRequest, err.Error(), nil))
	}

	return ctx.JSON(http.StatusOK, models.ConstructWebResponse(http.StatusOK, "SUCCESS", linkChangePassword))
}

func ActivateAccount(ctx echo.Context) error {
	token := ctx.Param("token")
	var decodedByte, _ = base64.StdEncoding.DecodeString(token)
	var decodedString = strings.Split(string(decodedByte), "&")
	var email = decodedString[0]
	tokenTime, _ := strconv.ParseInt(decodedString[1], 10, 64)
	employee := models.GetEmployeeByEmail(email)

	if employee.Email == "" {
		return ctx.JSON(http.StatusForbidden, models.ConstructWebResponse(http.StatusForbidden, "User not found", nil))
	}
	if employee.Verified == true {
		return ctx.JSON(http.StatusForbidden, models.ConstructWebResponse(http.StatusForbidden, "User already verified", nil))
	}
	if tokenTime+300 < time.Now().Unix() {
		return ctx.JSON(http.StatusForbidden, models.ConstructWebResponse(http.StatusForbidden, "Link expired", nil))
	}

	err := models.UpdateVerificationEmployee(email)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.ConstructWebResponse(http.StatusInternalServerError, err.Error(), nil))
	}

	err = models.CreateEmployeePerformance(employee.Id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.ConstructWebResponse(http.StatusInternalServerError, err.Error(), nil))
	}

	return ctx.JSON(http.StatusOK, models.ConstructWebResponse(http.StatusOK, "SUCCESS", nil))
}

func ForgetPassword(ctx echo.Context) error {
	var newPassword models.ForgetPasswordRequest
	err := ctx.Bind(&newPassword)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.ConstructWebResponse(http.StatusBadRequest, err.Error(), nil))
	}

	token := ctx.Param("token")
	var decodedByte, _ = base64.StdEncoding.DecodeString(token)
	var decodedString = strings.Split(string(decodedByte), "&")
	var email = decodedString[0]
	tokenTime, _ := strconv.ParseInt(decodedString[1], 10, 64)

	employee := models.GetEmployeeByEmail(email)
	if employee.Email == "" {
		return ctx.JSON(http.StatusForbidden, models.ConstructWebResponse(http.StatusForbidden, "User not found", nil))
	}
	if newPassword.Password != newPassword.ConfirmPassword {
		return ctx.JSON(http.StatusForbidden, models.ConstructWebResponse(http.StatusForbidden, "Password and Confirm Password Not Match", nil))
	}
	if tokenTime+300 < time.Now().Unix() {
		return ctx.JSON(http.StatusForbidden, models.ConstructWebResponse(http.StatusForbidden, "Link expired", nil))
	}
	newPassword.Password = helpers.HashAndSalt([]byte(newPassword.Password))

	err = models.UpdatePasswordEmployee(email, newPassword)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.ConstructWebResponse(http.StatusInternalServerError, err.Error(), nil))
	}

	return ctx.JSON(http.StatusOK, models.ConstructWebResponse(http.StatusOK, "SUCCESS", nil))
}

func Login(ctx echo.Context) error {
	var loginRequest models.LoginRequest
	var employee models.Employee
	err := ctx.Bind(&loginRequest)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.ConstructWebResponse(http.StatusInternalServerError, err.Error(), nil))
	}

	//VALIDATE REQUEST
	err = ctx.Validate(loginRequest)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.ConstructWebResponse(http.StatusInternalServerError, err.Error(), nil))
	}

	employee = models.GetEmployeeByEmail(loginRequest.Email)

	if employee.Email == "" {
		return ctx.JSON(http.StatusBadRequest, models.ConstructWebResponse(http.StatusBadRequest, "User Not Found", nil))
	}
	if helpers.ComparePasswords(employee.Password, []byte(loginRequest.Password)) == false {
		return ctx.JSON(http.StatusBadRequest, models.ConstructWebResponse(http.StatusBadRequest, "Invalid Password", nil))
	}
	if !employee.Verified {
		return ctx.JSON(http.StatusBadRequest, models.ConstructWebResponse(http.StatusBadRequest, "Member Not Active", nil))
	}

	//CONSTRUCT JWT
	sign := jwt.New(jwt.SigningMethodHS256)
	claims := sign.Claims.(jwt.MapClaims)

	claims["id"] = employee.Id
	claims["email"] = employee.Email
	claims["isAdmin"] = employee.IsAdmin


	token, err := sign.SignedString([]byte(helpers.GetEnv("JWT_SIGNATURE_KEY")))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.ConstructWebResponse(http.StatusInternalServerError, err.Error(), nil))
	}

	var getEmployeePerformance = models.GetEmployeePerformanceById(employee.Id)
	loc, _ := time.LoadLocation("Asia/Jakarta")
	hour := time.Now().In(loc).Hour()
	currentDate := strings.Split(time.Now().In(loc).String(), " ")[0]
	if currentDate != getEmployeePerformance.Updated_At {
		if 6 <= hour && hour <= 7 {
			err = models.UpdateScoreEmployeePerformance(getEmployeePerformance.Score+3, employee.Id, currentDate)
		} else if 7 < hour && hour <= 8 {
			err = models.UpdateScoreEmployeePerformance(getEmployeePerformance.Score+2, employee.Id, currentDate)
		} else if 8 < hour {
			err = models.UpdateScoreEmployeePerformance(getEmployeePerformance.Score+1, employee.Id, currentDate)
		}
	} 

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.ConstructWebResponse(http.StatusInternalServerError, err.Error(), nil))
	}

	c := &http.Cookie{}
	if storedCookie, _ := ctx.Cookie("token"); storedCookie != nil {
		c = storedCookie
	}

	if c.Value == "" {
		c.Name = "token"
		c.Value = token
		c.Path = "/"
		c.Domain = helpers.GetEnv("DOMAIN")
		c.Expires = time.Now().Add(time.Hour)

		ctx.SetCookie(c)
	}
	return ctx.JSON(http.StatusOK, models.ConstructWebResponse(http.StatusOK, "SUCCESS", token))
}

func Logout(ctx echo.Context) error {
	tokenCookie := &http.Cookie{}
	tokenCookie.Expires = time.Now().Add(-100 * time.Hour)
	tokenCookie.MaxAge = -1
	tokenCookie.Name = "token"
	tokenCookie.Path = "/"
	tokenCookie.Domain = helpers.GetEnv("DOMAIN")
	ctx.SetCookie(tokenCookie)

	return ctx.JSON(http.StatusOK, models.ConstructWebResponse(http.StatusOK, "SUCCESS", nil))
}
