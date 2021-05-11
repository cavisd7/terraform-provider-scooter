package server

import (
	"sync"
	"net/http"

	"github.com/gorilla/mux"
)

type Service struct {
	connString string
	items      map[string]Item
	sync.RWMutex
}

func handler(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWritter, r *http.Request) {
		handlerFunc(w, r)
		return
	}
}

func NewService(connString string, items map[string]Item) *Service {
	return &Service{
		connString: connString,
		items:      items,
	}
}

func (s *Service) Serve() error {
	r := mux.NewRouter()

	r.HandleFunc("/item", handler(s.PostItem)).Methods("POST")
	r.HandleFunc("/item", handler(s.GetItems)).Methods("GET")
	r.HandleFunc("/item/{name}", handler(s.GetItem)).Methods("GET")
	r.HandleFunc("/item/{name}", handler(s.PutItem)).Methods("PUT")
	r.HandleFunc("/item/{name}", handler(s.DeleteItem)).Methods("DELETE")

	err := http.ListenAndServe(s.connString, r)
	if err != nil {
		return err
	}

	return nil
}
