package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"Multithreading/api/service"
)

func AddressHandler(w http.ResponseWriter, r *http.Request) {
	cep := chi.URLParam(r, "cep")

	address, err := service.FetchAddress(cep)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching address: %v", err), http.StatusInternalServerError)
		return
	}

	// Marshal the address to JSON
	response, err := json.Marshal(address)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error marshaling address: %v", err), http.StatusInternalServerError)
		return
	}

	// Write the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func main() {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Routes
	r.Get("/address/{cep}", AddressHandler)

	// Start server
	http.ListenAndServe(":8060", r)
}
