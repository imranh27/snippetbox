//Alex Edwards course.
package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golangcollege/sessions"
	"github.com/imranh27/snippetbox/pkg/models/mysql"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

//to make this globally available. Functions are then created as methods against this.
type application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	session       *sessions.Session
	snippets      *mysql.SnippetModel
	templateCache map[string]*template.Template
}

func main() {

	//define command line flags, default = 4000
	addr := flag.String("addr", ":4000", "HTTP network address")
	//MySQL dsn connection string
	dsn := flag.String("dsn", "web:password@/snippetbox?parseTime=true", "MySQL data source name")

	//Define session secret
	secret := flag.String("secret", "s6Ndh+pPbnzHbS*+9Pk8qGWhTzb-pa@ge", "Secret key")

	//parse command line
	flag.Parse()

	//add logging
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	//database connection pool
	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	//Initialise new session manager, session expires after 12 hours
	session := sessions.New([]byte(*secret))
	session.Lifetime = 12 * time.Hour

	//initialise new template cache 216
	templateCache, err := newTemplateCache("./ui/html/")
	if err != nil {
		errorLog.Fatal(err)
	}

	app := &application{
		errorLog:      errorLog,
		infoLog:       infoLog,
		session:       session,
		snippets:      &mysql.SnippetModel{DB: db},
		templateCache: templateCache,
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting Server on %s", *addr)

	fmt.Println("https://localhost" + *addr)
	fmt.Println("https://localhost" + *addr + "/snippet/1")
	fmt.Println("https://localhost" + *addr + "/snippet/create")
	fmt.Println("https://localhost" + *addr + "/static/")

	//For HTTPS we use ListenAndServeTLS passing in Certificate and Private key
	err = srv.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem")
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
