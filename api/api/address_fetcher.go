package api

import "Multithreading/api/entity"

type AddressFetcher interface {
	FetchAddressFromBrasilAPI(cep string) (entity.BrasilAPIAddress, error)
	FetchAddressFromViaCEP(cep string) (entity.ViaCEPAddress, error)
}
