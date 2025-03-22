package models

import "time"

type User struct {
	Username  string    `bson:"username" json: "username"`
	Email     string    `bson:"email" json: "email"`
	Password  string    `bson:"password" json: "password"`
	CreatedAt time.Time `bson:"create_at" json: "create_at"`
}
