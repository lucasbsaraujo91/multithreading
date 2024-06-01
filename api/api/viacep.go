package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"Multithreading/api/entity"
)

func FetchAddressFromViaCEP(cep string) (entity.ViaCEPAddress, error) {
	//time.Sleep(2 * time.Second)
	var address entity.ViaCEPAddress

	client := http.Client{}

	resp, err := client.Get(fmt.Sprintf("http://viacep.com.br/ws/%s/json/", cep))
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
