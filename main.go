package main

import (
	"net/http"

	"github.com/rs/cors"

	//"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/getquantity/", getquantity).Methods("GET")
	router.HandleFunc("/getbyid/", getbyid).Methods("GET")
	router.HandleFunc("/saveproduct/", addproduct).Methods("POST")
	router.HandleFunc("/saveproduct2/", addproduct2).Methods("POST")
	router.HandleFunc("/saveproduct3/", addproduct3).Methods("POST")
	router.HandleFunc("/saveproduct4/", addproduct4).Methods("POST")
	router.HandleFunc("/saveproduct5/", addproduct5).Methods("POST")
	router.HandleFunc("/uptproduct/{id}", updateqty).Methods("PUT")
	//cors := handlers.AllowedMethods([]string{"*", "PUT", "POST", "GET", "DELETE"})
	handler := cors.AllowAll().Handler(router)
	http.ListenAndServe(":8080", handler) // handlers.CORS(cors)(router)
}
