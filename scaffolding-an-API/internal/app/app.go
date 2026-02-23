package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/stainsleto/project-go/scaffolding/internal/api"
)

type Application struct {
	Logger         *log.Logger
	WorkoutHandler *api.Workouthandler
}

func NewApplication() (*Application, error) {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// out stores will go here

	// our handlers will go here

	workouthandler := api.NewWorkouthandler()

	app := &Application{Logger: logger, WorkoutHandler: workouthandler}

	return app, nil
}

// Health Check function

func (a *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status is available\n")
}
