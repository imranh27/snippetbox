//Alex Edwards course.
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)



func main() {

	//define command line flags, default = 4000
	addr := flag.String("addr", ":4000", "HTTP network address")

	//parse command line
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))


	log.Printf("Starting server on %s", *addr)

	fmt.Println("http://localhost" + *addr )
	fmt.Println("http://localhost:4000/snippet?id=123")
	fmt.Println("http://localhost:4000/snippet/create")
	fmt.Println("http://localhost:4000/static/")


	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)

}
