package cloudflare

import (
	"context"
	"net/http"

	"pod-chef-back-end/pkg"

	cf "github.com/cloudflare/cloudflare-go"
	"github.com/labstack/gommon/log"
)

//AddSubDomainName method responsible for adding a new subdomain to the cloudlfare account
func (repo *CloudflareRepository) AddSubDomainName(name string) (bool, error) {
	// Most API calls require a Context
	ctx := context.Background()

	zoneID := repo.Viper.GetString("TOKEN_CLOUDFLARE")
	zoneIP := repo.Viper.GetString("ZONE_IP_CLOUDFLARE")

	_, err := repo.Client.CreateDNSRecord(ctx, zoneID, cf.DNSRecord{Type: "CNAME", Name: name, Content: zoneIP, TTL: 1})

	if err != nil {
		//print the error stack
		log.Error(err)

		//return a custom error
		return false, &pkg.Error{Err: err, Code: http.StatusInternalServerError, Message: "Internal error"}
	}
	return true, nil
}

//DeleteSubDomainName method responsible for deleting a new subdomain to the cloudlfare account
func (repo *CloudflareRepository) DeleteSubDomainName(name string) (bool, error) {
	// Most API calls require a Context
	ctx := context.Background()

	zoneID := repo.Viper.GetString("TOKEN_CLOUDFLARE")

	err := repo.Client.DeleteDNSRecord(ctx, zoneID, name)

	if err != nil {
		//print the error stack
		log.Error(err)

		//return a custom error
		return false, &pkg.Error{Err: err, Code: http.StatusInternalServerError, Message: "Internal error"}
	}
	return true, nil
}
