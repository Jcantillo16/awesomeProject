package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Song struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Results  interface{}        `bson:"results" json:"results"`
	Name     string             `bson:"trackName" json:"name"`
	Artist   string             `bson:"artistName" json:"artist"`
	Price    float64            `bson:"trackPrice" json:"price"`
	Duration int64              `bson:"trackTimeMillis" json:"duration"`
	Album    string             `bson:"collectionName" json:"album"`
	Artwork  string             `bson:"artworkUrl100" json:"artwork"`
	Origin   string             `bson:"url" json:"origin"`
}
