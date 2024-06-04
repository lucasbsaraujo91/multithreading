package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"

	"Multithreading/api/api"
	"Multithreading/api/service"
)

func AddressHandler(fetcher api.AddressFetcher) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cep := chi.URLParam(r, "cep")

		address, err := service.FetchAddress(cep, fetcher)
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
}

func main() {
	r := chi.NewRouter()

	// Criar uma instância do fetcher padrão
	fetcher := &api.DefaultAddressFetcher{}

	// Definir rotas
	r.Get("/address/{cep}", AddressHandler(fetcher))

	http.ListenAndServe(":8060", r)
}
