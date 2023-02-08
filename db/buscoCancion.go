package db

import (
	"awesomeProject/models"
	"awesomeProject/utils"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func BuscoCancion(param string) (models.Song, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("twittor")
	col := db.Collection("songs")
	var resultado models.Song
	err := col.FindOne(
		ctx,
		bson.M{"$or": []bson.M{{"trackName": utils.LikeMongo(param)},
			{"artistName": utils.LikeMongo(param)},
			{"collectionName": utils.LikeMongo(param)}}}).Decode(&resultado)
	if err != nil {
		fmt.Println("entra a buscar en itunes")
		GetItunes(param)
		return resultado, err
	}
	return resultado, nil

}
