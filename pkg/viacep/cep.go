package viacep

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/allanmaral/go-expert-multithreading-challenge/pkg/model"
)

type viaCEPResponse struct {
	Cep          string `json:"cep"`
	Street       string `json:"logradouro"`
	Complement   string `json:"complemento"`
	Neighborhood string `json:"bairro"`
	City         string `json:"localidade"`
	State        string `json:"uf"`
}

func SearchCEP(cep string) (model.CEP, error) {
	url := fmt.Sprintf("http://viacep.com.br/ws/%s/json/", cep)
	req, err := http.Get(url)
	if err != nil {
		return model.CEP{}, err
	}
	defer req.Body.Close()

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return model.CEP{}, err
	}

	var b viaCEPResponse
	err = json.Unmarshal(body, &b)
	if err != nil {
		return model.CEP{}, err
	}

	c := model.CEP{
		Cep:          b.Cep,
		Street:       b.Street,
		Neighborhood: b.Neighborhood,
		City:         b.City,
		State:        b.State,
		Service:      "ViaCEP",
	}

	return c, nil
}
