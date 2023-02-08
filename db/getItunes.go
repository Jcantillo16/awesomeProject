package db

import (
	"awesomeProject/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

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
