package main

import (
	"log"
	"net/http"
)

type api struct {
	addr string
}

func (s *api) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:

		switch r.URL.Path {

		case "/":
			w.Write([]byte("GET Method"))
			return

		case "/index":
			w.Write([]byte("GET Index"))
			return

		default:
			w.Write([]byte("404 page"))
		}

	case http.MethodPost:
		w.Write([]byte("POST Method"))
	}
}

func (a *api) createUsersHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("User Created!"))
}

func (a *api) getUsersHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Got Users!"))
}

func main() {

	api := &api{addr: ":8080"}

	mux := http.NewServeMux()

	srv := &http.Server{
		Addr:    api.addr,
		Handler: mux,
	}

	mux.HandleFunc("GET /users", api.getUsersHandler)
	mux.HandleFunc("POST /users", api.createUsersHandler)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
