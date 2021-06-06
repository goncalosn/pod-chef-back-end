package services

import (
	"context"
	"time"

	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	gmail "google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

//GmailRepository gmail client responsible for acessing the gmail api
type GmailRepository struct {
	Client *gmail.Service
}

//NewGmailRepository new connection to the gmail api
func NewGmailRepository(viper *viper.Viper) *GmailRepository {
	return &GmailRepository{
		Client: Client(viper),
	}
}

//Client responsible for creating the connection to the gmail api
func Client(viper *viper.Viper) *gmail.Service {
	id := viper.Get("GMAIL_ID").(string)
	secret := viper.Get("GMAIL_SECRET").(string)
	refresh := viper.Get("GMAIL_REFRESH_TOKEN").(string)
	access := viper.Get("GMAIL_ACCESS_TOKEN").(string)

	log.Info("starting connection to the gmail api")
	config := oauth2.Config{
		ClientID:     id,
		ClientSecret: secret,
		Endpoint:     google.Endpoint,
		RedirectURL:  "http://localhost",
	}

	token := oauth2.Token{
		AccessToken:  access,
		RefreshToken: refresh,
		TokenType:    "Bearer",
		Expiry:       time.Now(),
	}

	var tokenSource = config.TokenSource(context.Background(), &token)

	srv, err := gmail.NewService(context.Background(), option.WithTokenSource(tokenSource))
	if err != nil {
		log.Fatal("Unable to retrieve Gmail client: %v", err)
	}

	if srv != nil {
		log.Info("Email service is initialized \n")
	}

	log.Info("connection to the database sucessful")
	return srv
}
