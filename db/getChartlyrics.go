package db

import (
	"awesomeProject/models"
	"awesomeProject/utils"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func GetChartlyrics(artist string, song string) (models.Song, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("twittor")
	col := db.Collection("lyrics")
	var resultado models.Song
	err := col.FindOne(
		ctx,
		bson.M{"$or": []bson.M{{"artist": utils.LikeMongo(artist)},
			{"song": utils.LikeMongo(song)}}}).Decode(&resultado)
	if err != nil {
		fmt.Println("entra a buscar en itunes")
		GetItunes(artist)
		return resultado, err
	}
	return resultado, nil
}
