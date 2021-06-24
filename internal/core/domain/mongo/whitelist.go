package mongo

import "time"

type WhitelistUser struct {
	Email string    `bson:"email"`
	Data  time.Time `bson:"data"`
}
