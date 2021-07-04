package cloudflare

import (
	cf "github.com/cloudflare/cloudflare-go"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

//CloudflareRepository cloudflare client responsible for acessing the cloudflare api
type CloudflareRepository struct {
	Client *cf.API
	Viper  *viper.Viper
}

//NewCloudflareRepository new connection to the cloudflare api
func NewCloudflareRepository(viper *viper.Viper) *CloudflareRepository {
	return &CloudflareRepository{
		Client: Client(),
		Viper:  viper,
	}
}

//Client responsible for creating the connection to the email api
func Client() *cf.API {
	apiToken := viper.GetString("TOKEN_CLOUDFLARE")

	// Construct a new API object
	client, err := cf.NewWithAPIToken(apiToken)
	if err != nil {
		log.Fatal(err)
	}

	return client
}
