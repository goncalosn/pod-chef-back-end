module pod-chef-back-end

go 1.16

require (
	github.com/cloudflare/cloudflare-go v0.18.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/labstack/echo/v4 v4.3.0
	github.com/labstack/gommon v0.3.0
	github.com/spf13/viper v1.8.1
	go.mongodb.org/mongo-driver v1.5.3
	golang.org/x/crypto v0.0.0-20210616213533-5ff15b29337e
	golang.org/x/net v0.0.0-20210614182718-04defd469f4e // indirect
	k8s.io/api v0.21.2
	k8s.io/apimachinery v0.21.2
	k8s.io/client-go v0.21.2
)
