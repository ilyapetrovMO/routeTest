package routes

import (
	"log"
	"net/http"
)

type Route struct {
	Route   string
	Handler func(http.ResponseWriter, *http.Request)
}

var (
	routes []*Route
)

func init() {
	routes = []*Route{}
	routes = append(routes, DefaultRoute(), ImgRoute())
}

func RegisterRoutes() []*Route {
	for _, v := range routes {
		log.Println("added route: " + v.Route)
	}
	return routes
}
