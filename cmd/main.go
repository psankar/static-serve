package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

// go:embed public
var public embed.FS

func main() {
	// cd to the public directory
	publicFS, err := fs.Sub(public, "public")
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/embedded/", http.StripPrefix("/embedded/",
		http.FileServer(http.FS(publicFS))))

	http.Handle("/public/",
		http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
