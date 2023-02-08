package db

import (
	"awesomeProject/models"
	"awesomeProject/utils"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func BuscoCanciones(param string) ([]models.Song, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("twittor")
	col := db.Collection("songs")
	var resultado []models.Song
	cursor, err := col.Find(
		ctx,
		bson.M{"$or": []bson.M{{"trackName": utils.LikeMongo(param)},
			{"artistName": utils.LikeMongo(param)},
			{"collectionName": utils.LikeMongo(param)}}})
	if err != nil {
		fmt.Println("entra a buscar en itunes")
		GetItunes(param)
		GetChartlyrics(param)
		return resultado, err
	}
	for cursor.Next(context.TODO()) {
		var registro models.Song
		err := cursor.Decode(&registro)
		if err != nil {
			return resultado, err
		}
		resultado = append(resultado, registro)

	}
	count := len(resultado)
	fmt.Println("count", count)
	return resultado, nil
}
