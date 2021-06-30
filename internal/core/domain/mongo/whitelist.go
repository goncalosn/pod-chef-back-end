package mongo

import "time"

type WhitelistUser struct {
	ID    string    `bson:"_id" json:"id"`
	Email string    `bson:"email" json:"email"`
	Date  time.Time `bson:"date" json:"date"`
}
