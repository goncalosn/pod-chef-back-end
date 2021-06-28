package mongo

import "time"

type WhitelistUser struct {
	Email string    `bson:"email"`
	Date  time.Time `bson:"date"`
}
