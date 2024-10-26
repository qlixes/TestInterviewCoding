package main

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type Smartphone struct {
	Code   string `json:"code"`
	Name   string `json:"name"`
	Model  string `json:"model"`
	Tech   string `json:"tech"`
	Status string `json:"status"`
}

var file *os.File

func writeToCsv() {
	items := []Smartphone{
		{
			Code:   "rv1",
			Name:   "Rover #1",
			Model:  "car",
			Tech:   "AI, car",
			Status: "progress",
		},
		{
			Code:   "tfx",
			Name:   "Transformer X",
			Model:  "transformation",
			Tech:   "car, robot",
			Status: "active",
		},
		{
			Code:   "px1",
			Name:   "Pacifista 1",
			Model:  "humanoid",
			Tech:   "AI, robot",
			Status: "active",
		},
		{
			Code:   "ax8",
			Name:   "Android #8",
			Model:  "humanoid",
			Tech:   "cyborg",
			Status: "active",
		},
		{
			Code:   "ax7",
			Name:   "Android #7",
			Model:  "humanoid",
			Tech:   "cyborg",
			Status: "inactive",
		},
	}

	file, err := os.Create("smartphone.csv")

	if err != nil {
		log.Fatal("error creating file", err)
	}

	defer file.Close()

	writer := csv.NewWriter(file)

	defer writer.Flush()

	for _, item := range items {
		data := []string{item.Code, item.Name, item.Model, item.Tech, item.Status}

		writer.Write(data)
	}
}

func showItems(w *http.Response, r *http.Request) {

}

func main() {

	writeToCsv()

	file, err := os.Open("smartphone.csv")

	if err != nil {
		log.Fatal("error open csv because ", err)
	}

	defer file.Close()

	reader := csv.NewReader(file)

	if err != nil {
		log.Fatal("error load csv because ", err)
	}

	http.HandleFunc("/items", showItems)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("error start server because ", err)
	}
}
