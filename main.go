package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type context struct {
	Request *http.Request
	Writer  http.ResponseWriter
	Albums  album
}

func (c *context) appendJson(status int, obj any) {
	c.Writer.Header().Set("Content-Type", "applications/json")
	response, err := json.Marshal(obj)
	if err != nil {
		fmt.Println("Error marshalling the program. Err: ", err)
	}

	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write(response)
}

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

// func appendJson(w http.ResponseWriter, alb album) {
// 	w.Header().Set("Content-Type", "applications/json")
// 	response, err := json.Marshal(alb)
// 	if err != nil {
// 		fmt.Println("Error marshalling the program. Err: ", err)
// 	}
//
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(response)
// }

func getAlbums(w http.ResponseWriter, r *http.Request) {
	var con *context
	con.appendJson(http.StatusCreated, albums)
}

func postAlbums(w http.ResponseWriter, r *http.Request, c *context) {
	if r.Method == http.MethodPost {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println("Error reading the body request. Error: ", err)
		}

		var newAlbums album
		json.Unmarshal(body, &newAlbums)

		// albums := append(albums, newAlbums)
		// getAlbums(http.StatusCreated, newAlbums)
		c.appendJson(http.StatusCreated, newAlbums)

	}
}

func main() {

	http.HandleFunc("/albums", getAlbums)

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println(err)
	}
}
