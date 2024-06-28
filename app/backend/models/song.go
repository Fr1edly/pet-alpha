package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Song struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title  string             `bson:"title" json:"title"`
	Artist string             `bson:"artist" json:"artist"`
	URL    string             `bson:"url" json:"url"`
}
