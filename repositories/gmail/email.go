package services

import (
	"encoding/base64"
	"net/http"

	"pod-chef-back-end/pkg"

	"github.com/labstack/gommon/log"
	gmail "google.golang.org/api/gmail/v1"
)

//SendEmailOAUTH2 method responsible for sending an email
func (repo *GmailRepository) SendEmailOAUTH2(to string, data interface{}, template string) (interface{}, error) {

	emailBody, err := pkg.ParseTemplate(template, data)
	if err != nil { //unable to parse email template
		//print the error stack
		log.Error(err)

		//return a custom error
		return false, &pkg.Error{Err: err, Code: http.StatusInternalServerError, Message: "Internal error"}
	}

	var message gmail.Message

	emailTo := "To: " + to + "\r\n"
	subject := "Subject: " + "Test Email form Gmail API using OAuth" + "\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"
	msg := []byte(emailTo + subject + mime + "\n" + emailBody)

	message.Raw = base64.URLEncoding.EncodeToString(msg)

	//call driven adapter responsible for sending an email
	_, err = repo.Client.Users.Messages.Send("me", &message).Do()
	if err != nil {
		//print the error stack
		log.Error(err)

		//return a custom error
		return false, &pkg.Error{Err: err, Code: http.StatusInternalServerError, Message: "Internal error"}
	}
	return true, nil
}
