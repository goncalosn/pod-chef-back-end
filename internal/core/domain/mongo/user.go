package mongo

//User struct with the needed values from the user account
type User struct {
	id    string   `bson:",inline"`
	Email string   `bson:"email"`
	Hash  [32]byte `bson:"hash"`
	Name  string   `bson:"name"`
	Role  string   `bson:"role"`
}
