package main

import (
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/dufeng/usermanager/common"
	"github.com/dufeng/usermanager/routers"
)

//Entry point of the program
func main() {
	// Calls startup logic
	common.StartUp()

	router := routers.InitRoutes()
	n := negroni.Classic()
	n.UseHandler(router)

	server := &http.Server{
		Addr:    common.AppConfig.Server,
		Handler: n,
	}
	log.Println("Listening...")
	server.ListenAndServe()
}
