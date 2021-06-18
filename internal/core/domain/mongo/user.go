package mongo

//User struct with the needed values from the user account
type User struct {
	Email string `bson:"email" json:"email"`
	Hash  string `bson:"hash" json:"password"`
	Name  string `bson:"name" json:"name"`
	Role  string `bson:"role" role:"role"`
}
