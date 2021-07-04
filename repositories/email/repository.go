package email

import (
	"net/smtp"

	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

//EmailRepository email client responsible for acessing the email api
type EmailRepository struct {
	Client smtp.Auth
	Viper  *viper.Viper
}

//NewEmailRepository new connection to the email api
func NewEmailRepository(viper *viper.Viper) *EmailRepository {
	return &EmailRepository{
		Client: Client(),
		Viper:  viper,
	}
}

//Client responsible for creating the connection to the email api
func Client() smtp.Auth {
	log.Info("creating smtp client")
	emailHost := viper.Get("EMAIL_HOST").(string)
	emailFrom := viper.Get("EMAIL_FROM").(string)
	emailPassword := viper.Get("EMAIL_PASSWORD").(string)

	emailAuth := smtp.PlainAuth("", emailFrom, emailPassword, emailHost)
	return emailAuth
}
