package main

import (
	"LogIt/backend/movies"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/view/", movies.ViewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
