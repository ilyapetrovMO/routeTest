package routes

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func ImgRoute() *Route {
	handler := func(w http.ResponseWriter, r *http.Request) {
		path := filepath.Join("images", r.URL.Path[len("/image/"):])
		img, err := os.Open(path)

		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusNotFound)
			return
		}
		defer img.Close()

		log.Println(path)

		w.Header().Set("Content-Type", "image/png")
		io.Copy(w, img)
	}
	route := &Route{Route: "/image/", Handler: handler}
	return route
}
