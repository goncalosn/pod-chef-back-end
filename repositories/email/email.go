package services

import (
	"fmt"
	"net/http"
	"net/smtp"

	"pod-chef-back-end/pkg"

	"github.com/labstack/gommon/log"
)

//SendEmailSMTP method responsible for sending an email
func (repo *EmailRepository) SendEmailSMTP(to string, subject string, emailBody string) error {
	emailFrom := repo.Viper.Get("EMAIL_FROM").(string)
	emailHost := repo.Viper.Get("EMAIL_HOST").(string)
	emailPort := repo.Viper.Get("EMAIL_PORT").(string)

	mime := "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"
	emailSubject := "Subject: " + subject + "\n"
	msg := []byte(emailSubject + mime + "\n" + emailBody)
	addr := fmt.Sprintf("%s:%s", emailHost, emailPort)

	if err := smtp.SendMail(addr, repo.Client, emailFrom, []string{to}, msg); err != nil {
		//print the error stack
		log.Error(err)

		//return a custom error
		return &pkg.Error{Err: err, Code: http.StatusInternalServerError, Message: "Internal error"}
	}
	return nil
}
