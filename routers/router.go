package routers

import (
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	router = SetUserRoutes(router)
	router = SetTaskRoutes(router)

	return router
}
