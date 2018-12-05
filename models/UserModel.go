package models

import "gopkg.in/mgo.v2/bson"

type UserModel struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	Username  string        `bson:"username" json:"username"`
	Firstname string        `bson:"firstname" json:"firstname"`
	Lastname  string        `bson:"lastname" json:"lastname"`
	Email     string        `bson:"email" json:"email"`
	Password  string        `bson:"password" json:"password"`
}
