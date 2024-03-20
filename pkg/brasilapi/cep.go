package brasilapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/allanmaral/go-expert-multithreading-challenge/pkg/model"
)

type brasilAPICEPResponse struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
}

func SearchCEP(cep string) (model.CEP, error) {
	url := fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep)
	req, err := http.Get(url)
	if err != nil {
		return model.CEP{}, err
	}
	defer req.Body.Close()

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return model.CEP{}, err
	}

	var b brasilAPICEPResponse
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
		Service:      "Brasil API",
	}

	return c, nil
}
