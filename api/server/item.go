package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Item struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (s *Service) GetItems(w http.ResponseWriter, r *http.Request) {
	s.RLock()
	defer s.RUnlock()

	err := json.NewEncoder(w).Encode(s.items)
	if err != nil {
		log.Println(err)
	}
}

func (s *Service) PostItem(w http.ResponseWriter, r *http.Request) {
	var item Item

	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	s.Lock()
	defer s.Unlock()

	s.items[item.Name] = item

	err = json.NewEncoder(w).Encode(item)
	if err != nil {
		log.Printf("Error sending response: %s", err)
	}
}

func (s *Service) PutItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	itemName := params["name"]
	if itemName == "" {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	var updatedItem Item
	err := json.NewDecoder(r.Body).Decode(&updatedItem)
	if err != nil {
		http.Error(w, err.Error(), 400)
	}

	s.Lock()
	defer s.Unlock()

	if !s.doesItemExist(itemName) {
		log.Printf("item %s does not exist on server", itemName)
		http.Error(w, fmt.Sprintf("item %v does not exist on server", itemName), http.StatusBadRequest)
		return
	}

	s.items[itemName] = updatedItem
	log.Printf("updated item: %s", itemName)
	err = json.NewEncoder(w).Encode(updatedItem)
	if err != nil {
		log.Printf("error sending response - %s", err)
	}
}

func (s *Service) DeleteItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	itemName := params["name"]
	if itemName == "" {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	s.Lock()
	defer s.Unlock()

	if !s.doesItemExist(itemName) {
		log.Printf("item %s does not exist on server", itemName)
		http.Error(w, fmt.Sprintf("item %v does not exist on server", itemName), http.StatusNotFound)
		return
	}

	delete(s.items, itemName)

	_, err := fmt.Fprintf(w, "deleted item: %s", itemName)
	if err != nil {
		log.Println(err)
	}
}

func (s *Service) GetItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	itemName := params["name"]

	if itemName == "" {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	s.RLock()
	defer s.RUnlock()

	if !s.doesItemExist(itemName) {
		log.Printf("item %s does not exist on server", itemName)
		http.Error(w, fmt.Sprintf("item %v does not exist on server", itemName), http.StatusNotFound)
		return
	}

	err := json.NewEncoder(w).Encode(s.items[itemName])
	if err != nil {
		log.Println(err)
		return
	}
}

func (s *Service) doesItemExist(itemName string) bool {
	if _, ok := s.items[itemName]; ok {
		return true
	}
	return false
}
