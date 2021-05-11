package server

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

type Service struct {
	connString string
	items      map[string]Item
	sync.RWMutex
}

func NewService(connString string, items map[string]Item) *Service {
	return &Service{
		connString: connString,
		items:      items,
	}
}

func (s *Service) Serve() error {
	r := mux.NewRouter()

	r.HandleFunc("/item", s.PostItem).Methods("POST")
	r.HandleFunc("/item", s.GetItems).Methods("GET")
	r.HandleFunc("/item/{name}", s.GetItem).Methods("GET")
	r.HandleFunc("/item/{name}", s.PutItem).Methods("PUT")
	r.HandleFunc("/item/{name}", s.DeleteItem).Methods("DELETE")

	log.Printf("Starting server on %s", s.connString)
	err := http.ListenAndServe(s.connString, r)
	if err != nil {
		return err
	}

	return nil
}
