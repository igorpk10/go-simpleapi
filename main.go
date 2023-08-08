package main

import (
	"fmt"
	"igorpk/simpleapi/database"
	"igorpk/simpleapi/models"
	"igorpk/simpleapi/routes"
)

func main() {
	models.Personalities = []models.Personality{
		{Id: 1, Name: "Name 1", History: "History 1"},
		{Id: 2, Name: "Name 2", History: "History 2"},
	}

	database.ConnectDB()
	fmt.Println("Server up")
	routes.HandleRequest()
}
