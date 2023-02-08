package db

import (
	"awesomeProject/models"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// InsertoRegistro es la parada final con la BD para insertar los datos de la cancion.
func InsertoRegistro(u models.Song) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("twittor")
	col := db.Collection("songs")
	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}
	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
