package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Song struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name     string             `bson:"name" json:"trackName"`
	Artist   string             `bson:"artist" json:"artistName"`
	Duration string             `bson:"duration" json:"trackTimeMillis"`
	Album    string             `bson:"album" json:"collectionName"`
	Artwork  string             `bson:"artwork" json:"artworkUrl130"`
	Price    string             `bson:"price" json:"collectionPrice"`
	Origin   string             `bson:"origin" json:"country"`
	Results  interface{}        `bson:"results" json:"results"`
}
