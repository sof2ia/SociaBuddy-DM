package user

import (
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type argCEP struct {
	name            string
	userCEP         string
	number          string
	complement      string
	expectedAddress *Address
}

func TestFindCepSuccessfully(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(w)
		result := `{
			"cep": "12246-260",
			"logradouro": "Avenida Salmão",
			"complemento": "",
			"bairro": "Parque Residencial Aquarius",
			"localidade": "São José dos Campos",
			"uf": "SP",
			"ibge": "3549904",
			"gia": "6452",
			"ddd": "12",
			"siafi": "7099"
		}`
		w.WriteHeader(200)
		_, err := w.Write([]byte(result))
		if err != nil {
			return
		}
	}))
	defer mockServer.Close()
	api := NewFacade(mockServer.Client())
	test := []argCEP{
		{
			name:       "FindCep() is succeed",
			userCEP:    "12246260",
			number:     "10",
			complement: "Torre C",
			expectedAddress: &Address{
				ZipCode:      "12246260",
				Country:      "Brasil",
				State:        "SP",
				City:         "São José dos Campos",
				Neighborhood: "Parque Residencial Aquarius",
				Street:       "Avenida Salmão",
				Number:       "10",
				Complement:   "Torre C",
			},
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			address, err := api.FindCep(tt.userCEP, tt.number, tt.complement)
			log.Printf("user: %v, err: %v", address, err)
			if !reflect.DeepEqual(address, tt.expectedAddress) {
				t.Fatalf("expected %+v, got %+v", tt.expectedAddress, address)
			}
		})
	}
}

func TestFindCepUnsuccessfully(t *testing.T) {
	mockServer2 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		result := `{
		 "status" : "Bad request"
		}`
		w.WriteHeader(400)

		_, err := w.Write([]byte(result))
		if err != nil {
			return
		}
	}))
	defer mockServer2.Close()
	api2 := NewFacade(mockServer2.Client())
	test := []argCEP{
		{
			name:            "CEP invalid - syntax",
			userCEP:         "122462601",
			number:          "11",
			complement:      "Torre A",
			expectedAddress: nil,
		},
		{
			name:            "CEP invalid - doesn't registered",
			userCEP:         "99999999",
			expectedAddress: nil,
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			address, err := api2.FindCep(tt.userCEP, tt.number, tt.complement)
			log.Printf("user: %v, err: %v", address, err)
			if !reflect.DeepEqual(address, tt.expectedAddress) {
				t.Fatalf("expected %+v, got %+v", tt.expectedAddress, address)
			}
		})
	}
}
