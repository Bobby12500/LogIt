package main

import (

	"log"
	"net/http"
	
	
)


func main() {
	http.HandleFunc("/view/", movies.viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
