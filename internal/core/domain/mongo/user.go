package mongo

//User struct with the needed values from the user account
type User struct {
	id       string `bson:",inline" json:"id"`
	Email    string `bson:"email" json:"email"`
	Password string `bson:"password" json:"password"`
	TokenIv  string `bson:"tokenIv" json:"tokenIv"`
	Name     string `bson:"name" json:"name"`
	Role     string `bson:"role" role:"role"`
}
