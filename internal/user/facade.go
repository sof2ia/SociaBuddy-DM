package user

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

type facade struct {
	client *http.Client
}
type Facade interface {
	FindCep(cepUser string, number string, complement string) (*Address, error)
}

func (f *facade) FindCep(cepUser string, number string, complement string) (*Address, error) {
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json", cepUser)
	resUrl, err := f.client.Get(url)
	if err != nil {
		return nil, err
	}
	if resUrl.StatusCode != 200 {
		log.Printf("finding this cep is failed")
		return nil, errors.New("finding this cep is failed")
	}
	var result map[string]string
	err = json.NewDecoder(resUrl.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	if result["erro"] == "true" {
		return nil, errors.New("finding this cep is failed")
	}

	err = resUrl.Body.Close()
	if err != nil {
		return nil, err
	}
	return &Address{
		ZipCode:      cepUser,
		Country:      "Brasil",
		State:        result["uf"],
		City:         result["localidade"],
		Neighborhood: result["bairro"],
		Street:       result["logradouro"],
		Number:       number,
		Complement:   complement,
	}, nil
}

func NewFacade(client *http.Client) Facade {
	return &facade{client}
}
