package models

import (
	"backend-nbs/helpers"
)

type ActivateLink struct {
	Email     string `json:"email"`
	Token     string `json:"token"`
	CreatedAt int64  `json:"created_at"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=6"`
}

type ForgetPasswordRequest struct {
	Password        string `json:"newPassword" validate:"required,gte=6"`
	ConfirmPassword string `json:"newConfirmPassword" validate:"required"`
}

func UpdateVerificationEmployee(email string) error {
	db, err := helpers.ConnectMySQL()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("update employees set verified = true where email = ?", email)
	return err
}

func UpdatePasswordEmployee(email string, newPassword ForgetPasswordRequest) error {
	db, err := helpers.ConnectMySQL()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("update employees set password = ? where email = ?", newPassword.Password, email)
	return err

}
