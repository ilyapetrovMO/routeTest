package routes

import (
	"fmt"
	"net/http"
)

func DefaultRoute() *Route {
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `<!DOCTYPE html>
		<html lang="en">
		  <head>
			<meta charset="utf-8">
			<title>title</title>
			<link rel="stylesheet" href="style.css">
			<script src="script.js"></script>
		  </head>
		  <body>
			<p> Hello from Go! %s </p>
			<img src="/image/resitas.png">
		  </body>
		</html>`, r.URL.Path[1:])
	}
	route := &Route{Route: "/", Handler: handler}
	return route
}
