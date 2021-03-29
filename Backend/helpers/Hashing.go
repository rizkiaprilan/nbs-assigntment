package helpers

import (
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func HashAndSalt(pwd []byte) string {

	i, err := strconv.Atoi(GetEnv("BCRYPT_COST"))
	LogFatal(err)

	hash, err := bcrypt.GenerateFromPassword(pwd, i)
	LogFatal(err)

	return string(hash)
}

func ComparePasswords(hashedPwd string, plainPwd []byte) bool {

	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err == nil {
		return true
	}
	return false
}
