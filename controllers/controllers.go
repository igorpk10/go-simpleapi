package controllers

import (
	"encoding/json"
	"igorpk/simpleapi/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ReturnPersonalitie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, personalitie := range models.Personalities {
		if strconv.Itoa(personalitie.Id) == id {
			json.NewEncoder(w).Encode(personalitie)
		}
	}
}

func GetAllPersonalities(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personalities)
}
