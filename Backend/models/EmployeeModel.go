package models

import (
	"backend-nbs/helpers"

	_ "github.com/go-playground/validator/v10"
)

type EmployeeRegisterRequest struct {
	Fullname        string `json:"fullname" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,gte=6"`
	ConfirmPassword string `json:"confirmPassword" validate:"required"`
	Phone           string `json:"phone" validate:"required"`
}

type Employee struct {
	Id         string `json:"id"`
	Fullname   string `json:"fullname"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Phone      string `json:"phone"`
	IsAdmin    bool   `json:"isAdmin"`
	Verified   bool   `json:"verified"`
	Created_At string `json:"createdAt"`
	Updated_At string `json:"updatedAt"`
}

func GetEmployeeByEmail(email string) Employee {
	var employees Employee

	db, err := helpers.ConnectMySQL()
	helpers.LogFatal(err)
	defer db.Close()

	db.QueryRow("select id,fullname,email,password,phone,isAdmin,verified,created_at,updated_at from employees where email = ?", email).Scan(
		&employees.Id, &employees.Fullname, &employees.Email, &employees.Password, &employees.Phone, &employees.IsAdmin, &employees.Verified, &employees.Created_At, &employees.Updated_At)

	return employees
}

func CreateEmployee(employee Employee) error {

	db, err := helpers.ConnectMySQL()
	helpers.LogFatal(err)
	defer db.Close()

	_, err = db.Exec("insert into employees (id,fullname,email,password,phone) VALUES(?, ?, ?, ?, ?)", employee.Id, employee.Fullname, employee.Email, employee.Password, employee.Phone)
	return err
}
