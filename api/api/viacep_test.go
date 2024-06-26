package api

import (
	"Multithreading/api/entity"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchAddressFromViaCEP(t *testing.T) {
	expectedAddress := entity.ViaCEPAddress{
		CEP:         "12345678",
		Logradouro:  "Rua Domingos de Morais",
		Complemento: "Complemento",
		Bairro:      "Vila Mariana",
		Localidade:  "São Paulo",
		UF:          "SP",
		IBGE:        "123456",
		GIA:         "1234",
		DDD:         "11",
		SIAFI:       "1234",
	}

	// Cria um servidor HTTP de teste
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/ws/01001-000/json/" {
			t.Fatalf("expected request to /ws/01001-000/json/, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expectedAddress)
	}))
	defer ts.Close()

	// Substitui a URL base da função para apontar para o servidor de teste
	//originalURL := "http://viacep.com.br/ws/"
	newURL := ts.URL + "/ws/"

	client := &http.Client{}
	cep := "01001-000"

	// Cria uma função auxiliar para redirecionar a URL
	redirectedFetchAddress := func(cep string) (entity.ViaCEPAddress, error) {
		var address entity.ViaCEPAddress
		resp, err := client.Get(newURL + cep + "/json/")
		if err != nil {
			return address, err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return address, fmt.Errorf("failed to fetch address from ViaCEP: %s", resp.Status)
		}

		if err := json.NewDecoder(resp.Body).Decode(&address); err != nil {
			return address, err
		}

		return address, nil
	}

	// Testa a função auxiliar redirecionada
	address, err := redirectedFetchAddress(cep)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if address != expectedAddress {
		t.Errorf("expected address to be %+v, got %+v", expectedAddress, address)
	}
}
