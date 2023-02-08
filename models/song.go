package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Song struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Results  interface{}        `bson:"results" json:"results"`
	Name     string             `bson:"collectionCensoredName" json:"trackCensoredName"`
	Artist   string             `bson:"artistName" json:"artistName"`
	Price    float64            `bson:"trackPrice" json:"trackPrice"`
	Duration int64              `bson:"trackTimeMillis" json:"trackTimeMillis"`
	Album    string             `bson:"collectionName" json:"collectionName"`
	Artwork  string             `bson:"artworkUrl100" json:"artworkUrl100"`
}
