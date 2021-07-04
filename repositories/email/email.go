package email

import (
	"fmt"
	"net/http"
	"net/smtp"

	emailDomain "pod-chef-back-end/internal/core/domain/email"
	"pod-chef-back-end/pkg"
	email "pod-chef-back-end/pkg/email"

	"github.com/labstack/gommon/log"
)

//SendEmailSMTP method responsible for sending an email
func (repo *EmailRepository) SendEmailSMTP(to string, data *emailDomain.Email, template string) (bool, error) {
	emailFrom := repo.Viper.Get("EMAIL_FROM").(string)
	emailHost := repo.Viper.Get("EMAIL_HOST").(string)
	emailPort := repo.Viper.Get("EMAIL_PORT").(string)

	emailBody, err := email.ParseTemplate(template, data)
	if err != nil {
		//print the error stack
		log.Error(err)

		//return a custom error
		return false, &pkg.Error{Err: err, Code: http.StatusInternalServerError, Message: "Internal error"}
	}

	mime := "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"
	emailSubject := "Subject: " + data.Subject + "\n"
	msg := []byte(emailSubject + mime + "\n" + emailBody)
	addr := fmt.Sprintf("%s:%s", emailHost, emailPort)

	if err := smtp.SendMail(addr, repo.Client, emailFrom, []string{to}, msg); err != nil {
		//print the error stack
		log.Error(err)

		//return a custom error
		return false, &pkg.Error{Err: err, Code: http.StatusInternalServerError, Message: "Internal error"}
	}
	return true, nil
}
