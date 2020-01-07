package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dineshgowda24/go-authenticator/routes"
	"github.com/joho/godotenv"
)

func main() {

	e := godotenv.Load()
	if e != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println(e)
	port := os.Getenv("WEBAPP_PORT")

	// Handle routes
	http.Handle("/", routes.Handlers())
	// Set up a http fileServer to serve all our static css and javascript files
	// files will only be served via www.webapp.com/static/css/file.css or www.webapp.com/jss/file.js
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static"))))

	log.Printf("Server started and is listening on port '%s'", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
