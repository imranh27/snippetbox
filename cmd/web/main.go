//Alex Edwards course.
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

//to make this globally available. Functions are then created as methods against this.
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {

	//define command line flags, default = 4000
	addr := flag.String("addr", ":4000", "HTTP network address")

	//parse command line
	flag.Parse()

	//add logging
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting Server on %s", *addr)

	fmt.Println("http://localhost" + *addr)
	fmt.Println("http://localhost" + *addr + "/snippet?id=123")
	fmt.Println("http://localhost" + *addr + "/snippet/create")
	fmt.Println("http://localhost" + *addr + "/static/")

	err := srv.ListenAndServe()
	errorLog.Fatal(err)

}
