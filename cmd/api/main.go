package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/Chindi-M/goapi/internal/handlers"
	log "github.com/sirupsen/logrus"
)

func main(){

	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting GO API service...")

	fmt.Println(`
	GO API`)

	err := http.ListenAndServe("localhost:8000", r)
	if err != nil {
		log.Error(err)
	}

}