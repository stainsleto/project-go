package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/stainsleto/project-go/scaffolding/internal/routes"

	"github.com/stainsleto/project-go/scaffolding/internal/app"
)

func main() {

	var port int

	// we initialize this in the start of the application as go run main.go -port 8081 -->> if not mentioned the default is 8080

	flag.IntVar(&port, "port", 8080, "go backend server port")
	flag.Parse()

	app, err := app.NewApplication()

	if err != nil {
		panic(err) // using panic will retuen error and stops the application (crashes the application)
	}

	app.Logger.Printf("We are running on port %d\n", port)

	r := routes.SetupRoutes(app)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal(err)
	}

}
