package mongo

import "time"

//User struct with the needed values from the user account
type User struct {
	ID    string    `bson:"_id" json:"id"`
	Email string    `bson:"email" json:"email"`
	Hash  string    `bson:"hash" json:"-"`
	Name  string    `bson:"name" json:"name"`
	Role  string    `bson:"role" json:"role"`
	Date  time.Time `bson:"date" json:"date"`
}
