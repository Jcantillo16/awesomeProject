package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func InsertoRegistro(t interface{}, url string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("twittor")
	col := db.Collection("songs")
	//se debe insertar la url de donde se obtuvo la cancion en el campo origin del struct Song
	_, err := col.InsertOne(ctx, t, options.InsertOne().SetBypassDocumentValidation(false))

	if err != nil {
		log.Fatal(err)
	}
}
