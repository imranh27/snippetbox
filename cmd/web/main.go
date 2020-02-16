//Alex Edwards course.
package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
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
	//MySQL dsn connection string
	dsn := flag.String("dsn", "web:password@/snippetbox?parseTime=true", "MySQL data source name")

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

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	infoLog.Printf("Starting Server on %s", *addr)

	fmt.Println("http://localhost" + *addr)
	fmt.Println("http://localhost" + *addr + "/snippet?id=123")
	fmt.Println("http://localhost" + *addr + "/snippet/create")
	fmt.Println("http://localhost" + *addr + "/static/")

	err = srv.ListenAndServe()
	errorLog.Fatal(err)

}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
