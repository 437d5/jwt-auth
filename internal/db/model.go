package db

type User struct {
	Username string `bson:"name"`
	Email    string `bson:"email,omitempty"`
	Password string `bson:"password"`
}
