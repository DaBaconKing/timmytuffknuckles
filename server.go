package main

import (
	"log"
	"net/http"
)

func main() {
	// 1. Create a file server that serves files from the "public" directory.
	fs := http.FileServer(http.Dir("./public"))

	// 2. Handle requests for a specific URL path (e.g., "/static/").
	// http.StripPrefix removes the "/static" part from the request URL 
    // before the file server looks for the file in the "public" directory.
	http.Handle("/static/", http.StripPrefix("/static", fs))

    // Optional: Serve the index.html from the root path
    // http.Handle("/", http.FileServer(http.Dir("./public"))) // Use this if you want all routes to be static files

	// 3. Start the server
	log.Print("Listening on :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
,,jggyh
