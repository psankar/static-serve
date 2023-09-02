package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/public/",
		http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
