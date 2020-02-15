//Alex Edwards course.
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	//define command line flags, default = 4000
	addr := flag.String("addr", ":4000", "HTTP network address")

	//parse command line
	flag.Parse()

	//add logging
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	infoLog.Printf("Starting Server on %s", *addr)

	fmt.Println("http://localhost" + *addr)
	fmt.Println("http://localhost" + *addr + "/snippet?id=123")
	fmt.Println("http://localhost" + *addr + "/snippet/create")
	fmt.Println("http://localhost" + *addr + "/static/")

	//new http server struct to use error logging
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	err := srv.ListenAndServe()
	errorLog.Fatal(err)

}
