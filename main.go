package main

import (
	"log"
	"net/http"

	"github.com/ilyapetrovMO/routeTest/routes"
)

var (
	mux       *http.ServeMux
	routeList []*routes.Route
)

func init() {
	mux = http.NewServeMux()

	routeList = routes.RegisterRoutes()

	for _, v := range routeList {
		mux.HandleFunc(v.Route, v.Handler)
	}
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", mux))
}
