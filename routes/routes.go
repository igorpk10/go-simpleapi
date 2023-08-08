package routes

import (
	"igorpk/simpleapi/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/api/personalities", controllers.GetAllPersonalities)
	r.HandleFunc("/api/personalities/{id}", controllers.ReturnPersonalitie)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
