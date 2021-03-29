package helpers

import (
	"crypto/rand"
	"crypto/tls"
	"fmt"
	"gopkg.in/gomail.v2"
	"strconv"
)

func LinkGenerator() string {
	b := make([]byte, 15)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func GenerateVerficationLink(token string) string {
	var domain = GetEnv("DOMAIN")
	if domain == "http://localhost" {
		domain = domain + ":" + GetEnv("PORT")
	}
	return fmt.Sprintf("%s/account/activated/%s", domain, token)
}

func GenerateChangePasswordLink(token string) string {
	var domain = GetEnv("DOMAIN")
	if domain == "http://localhost" {
		domain = domain + ":" + GetEnv("PORT")
	}
	return fmt.Sprintf("%s/account/forget-password/%s", domain, token)
}

func SendEmail(receiver string, link string) error {
	linkActive := fmt.Sprintf("<div> <a href='%s'>%s</a> </div>", link, link)
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", GetEnv("CONFIG_SENDER_NAME"))
	mailer.SetHeader("To", receiver)
	mailer.SetAddressHeader("Cc", receiver, "me")
	mailer.SetHeader("Subject", "Verification your account employee")
	mailer.SetBody("text/html", "Verification your account employee with click link: "+linkActive)

	port, _ := strconv.Atoi(GetEnv("CONFIG_SMTP_PORT"))
	dialer := gomail.NewDialer(
		GetEnv("CONFIG_SMTP_HOST"),
		port,
		GetEnv("CONFIG_AUTH_EMAIL"),
		GetEnv("CONFIG_AUTH_PASSWORD"),
	)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	err := dialer.DialAndSend(mailer)
	return err
}

func SendMailForgotPassword(receiver string, link string) error {
	linkActive := fmt.Sprintf("<div> <a href='%s'>%s</a> </div>", link, link)
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", GetEnv("CONFIG_SENDER_NAME"))
	mailer.SetHeader("To", receiver)
	mailer.SetAddressHeader("Cc", receiver, "me")
	mailer.SetHeader("Subject", "Forget Password")
	mailer.SetBody("text/html", "Click link for update your password, "+linkActive)

	port, _ := strconv.Atoi(GetEnv("CONFIG_SMTP_PORT"))
	dialer := gomail.NewDialer(
		GetEnv("CONFIG_SMTP_HOST"),
		port,
		GetEnv("CONFIG_AUTH_EMAIL"),
		GetEnv("CONFIG_AUTH_PASSWORD"),
	)
	err := dialer.DialAndSend(mailer)
	return err
}
