package db

import (
	"context"
	"log"
	"time"
)

func InsertoRegistro(t interface{}, collection string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("twittor")
	col := db.Collection("songs")
	_, err := col.InsertOne(ctx, t)
	if err != nil {
		log.Fatal(err)
	}
}
