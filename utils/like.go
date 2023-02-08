package utils

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func LikeMongo(param string) bson.M {
	return bson.M{"$regex": primitive.Regex{Pattern: param, Options: "i"}}
}
