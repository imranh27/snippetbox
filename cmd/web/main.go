//Alex Edwards course.
package main

import (
	"fmt"
	"log"
	"net/http"
)



func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Print("Starting server on :4000")

	fmt.Println("http://localhost:4000")
	fmt.Println("http://localhost:4000/snippet?id=123")
	fmt.Println("http://localhost:4000/snippet/create")
	fmt.Println("http://localhost:4000/static/")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
