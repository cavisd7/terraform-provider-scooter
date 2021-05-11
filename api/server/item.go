package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Item struct {
	Name        string `json:"name"`
	Description string `json:"description"`
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
		fmt.Println("Error sending response: %s", err)
	}
}
