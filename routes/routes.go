package routes

import (
	"igorpk/simpleapi/controllers"
	"igorpk/simpleapi/middlewares"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middlewares.ContentType)
	r.HandleFunc("/", controllers.ReturnHome).Methods("Get")
	r.HandleFunc("/api/personalities", controllers.GetAllPersonalities).Methods("Get")
	r.HandleFunc("/api/personalitie/{id}", controllers.ReturnPersonalitie).Methods("Get")
	r.HandleFunc("/api/personalitie", controllers.CreatePersonality).Methods("Post")
	r.HandleFunc("/api/personalitie/{id}", controllers.DeletePersonality).Methods("Delete")
	r.HandleFunc("/api/personalitie", controllers.CreatePersonality).Methods("Put")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
