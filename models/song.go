package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Song struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Results  interface{}        `bson:"results" json:"results"`
	Name     string             `bson:"name" json:"trackCensoredName"`
	Artist   string             `bson:"artist" json:"artistName"`
	Price    float64            `bson:"price" json:"trackPrice"`
	Duration int64              `bson:"duration" json:"trackTimeMillis"`
	Album    string             `bson:"album" json:"collectionName"`
	Artwork  string             `bson:"artwork" json:"artworkUrl130"`
}
