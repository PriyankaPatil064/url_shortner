package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type URL struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	ShortCode string             `bson:"short_code"`
	LongURL   string             `bson:"long_url"`
}
