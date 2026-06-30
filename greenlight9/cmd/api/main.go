package main

import (
	//New Import
	//New import
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	// Import the pq driver so that it can register itself with the database/sql
	// package. Note that we alias this import to the blank identifier, to stop the Go
	// compiler complaining that the package isn't being used.
	_ "github.com/go-sql-driver/mysql" // New import
)

// Declare a string containing the application version number
// could be generated  at build time
const version = "1.0.0"

// config struct to hold config settings for the app
// operatinng envt and port the server will listen on
// We will read in these
// configuration settings from command-line flags when the application starts.
type config struct {
	port int
	env  string
}

//Application Struct to  hold dependencies for HTTP haNDLERS,HELPERS AND middleware

type application struct {
	config config
	logger *log.Logger
}

func main() {
	//Declare an istance ofthe config struct
	var cfg config

	//using port 4000  by default  and environment develoment setting sofar
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")

	// Define a new command-line flag for the MySQL DSN string.
	dsn := flag.String("dsn", "greenlight:Developer12@/greenlight?parseTime=true", "MySQL data source name")
	flag.Parse()

	//Initialise a new logger to write message to the stdout stream
	//prefixed with current date and time
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	//APPLICATION DEPENDENCIES
	//Passing new values to struct , via our new instance
	//Instance of the applocation struct , containing
	app := &application{
		config: cfg,    ///Set the struct's config field to the value of your cfg variable
		logger: logger, //Set the struct's logger field to the value of your logger variable
	}

	//Decalre a new servermux and add  /v1/healthcheck route which dispatches requests
	// to teh healthcheckHandler
	//mux := http.NewServeMux()
	//mux.HandleFunc("/v1/healthcheck", app.healthcheckHandler)

	//SETTING OUR SERVER WITH TIMEOUT SETTINGS
	//listening to port 4000 in the config struct and the servermux above as the handler
	srv := &http.Server{
		Addr: fmt.Sprintf(":%d", cfg.port),
		//Handler:      mux, change to  use httprouter instance app.routes()
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	//start the HTTP Server
	logger.Printf("starting %s server on %s", cfg.env, srv.Addr)
	err := srv.ListenAndServe()
	logger.Fatal(err)

	// The openDB() function returns a sql.DB connection pool.
	func openDB(cfg config) (*sql.DB, error) {
		//Use sql from .Open
	}

}
