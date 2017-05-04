package main

import (
	"net/http"
	"log"
    "fmt"
	"github.com/gorilla/mux"
)

func YourHandler(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("Gorilla!\n"))
    vars := mux.Vars(r)
    nombre := vars["nombre"]
    id := vars["id"]
    fmt.Fprintf(w,"Los parametros son "+nombre + " " + id)
}

func main() {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
    r.HandleFunc("/{nombre}/{id:[0-9]}", YourHandler).Methods("PUT")

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}