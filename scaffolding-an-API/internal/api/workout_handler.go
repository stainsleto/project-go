package api

import (
	"net/http"

	"fmt"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Workouthandler struct{}

func NewWorkouthandler() *Workouthandler {
	return &Workouthandler{}
}

func (wh *Workouthandler) HandleGetWorkoutByID(w http.ResponseWriter, r *http.Request) {
	paramsWorkOutID := chi.URLParam(r, "id")

	if paramsWorkOutID == "" {
		http.NotFound(w, r)
		return
	}

	workoutID, err := strconv.ParseInt(paramsWorkOutID, 10, 64)

	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "This is the workout ID %d\n", workoutID)

}

func (wh *Workouthandler) HandleCreateWorkout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "created a workout\n")
}
