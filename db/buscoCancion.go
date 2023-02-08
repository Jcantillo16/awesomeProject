package db

import (
	"awesomeProject/models"
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"io/ioutil"
	"log"
	"net/http"
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
		bson.M{"$or": []bson.M{{"trackCensoredName": param},
			{"artistName": param},
			{"collectionName": param}}}).Decode(&resultado)

	if err != nil {
		fmt.Println("entra a buscar en itunes")
		GetItunes(param)
		return resultado, err

	}

	return resultado, nil

}

func GetItunes(query string) []byte {
	url := "https://itunes.apple.com/search?term=" + query
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var resultado models.Song
	fmt.Println("resultado", resultado)
	json.Unmarshal(body, &resultado)
	for _, v := range resultado.Results.([]interface{}) {
		InsertoRegistro(v)
		return body
		BuscoCancion(query)
	}
	if err != nil {
		log.Fatal(err)

	}
	return body
}
