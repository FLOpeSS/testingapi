package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func createJson(w http.ResponseWriter, status int, obj any) {
	w.Header().Set("Content-type", "applications/json")
	response, err := json.MarshalIndent(obj, "", " ")
	if err != nil {
		fmt.Println("Error marshalling the program. Err: ", err)
	}
	w.WriteHeader(status)
	w.Write(response)
}

func requestAlbums(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var newAlbums album
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println("Error reading the body request. Error: ", err)
		}

		json.Unmarshal(body, &newAlbums)

		albums = append(albums, newAlbums)
		createJson(w, http.StatusCreated, newAlbums)
	}

	if r.Method == http.MethodGet {
		createJson(w, http.StatusCreated, albums)
	}
}

func main() {
	PORT := ":8000"
	http.HandleFunc("/albums", requestAlbums)

	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		fmt.Println(err)
	}
}
