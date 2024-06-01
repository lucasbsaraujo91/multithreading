package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"Multithreading/api/entity"
)

func FetchAddressFromBrasilAPI(cep string) (entity.BrasilAPIAddress, error) {
	//time.Sleep(2 * time.Second)
	var address entity.BrasilAPIAddress

	client := http.Client{}

	resp, err := client.Get(fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep))
	if err != nil {
		return address, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return address, fmt.Errorf("failed to fetch address from BrasilAPI: %s", resp.Status)
	}

	if err := json.NewDecoder(resp.Body).Decode(&address); err != nil {
		return address, err
	}

	return address, nil
}
