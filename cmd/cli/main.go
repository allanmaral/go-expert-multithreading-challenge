package main

import (
	"fmt"
	"os"
	"time"

	"github.com/allanmaral/go-expert-multithreading-challenge/pkg/brasilapi"
	"github.com/allanmaral/go-expert-multithreading-challenge/pkg/model"
	"github.com/allanmaral/go-expert-multithreading-challenge/pkg/viacep"
)

func searchViaCEP(cep string, ch chan<- model.CEP) {
	res, err := viacep.SearchCEP(cep)
	if err != nil {
		// If I log the error it can make it more dificult to read the result
		// if the other provider successfully responds. Just returning for now
		return
	}
	ch <- res
}

func searchBrasilAPI(cep string, ch chan<- model.CEP) {
	res, err := brasilapi.SearchCEP(cep)
	if err != nil {
		return
	}
	ch <- res
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("\"%s\" requer pelo menos 1 argumento.\n\nUso: %s CEP\n\nBusca um CEP e exibe Logradouro, Bairro, Cidade e Estado\n", os.Args[0], os.Args[0])
		return
	}
	cep := os.Args[1]

	// I see two ways to solve this: make two channels, one for each request,
	// each with its own response model, or abstract the response into a single
	// model and make both use the same channel. In this example, I chose the
	// latter.
	ch := make(chan model.CEP)

	go searchViaCEP(cep, ch)
	go searchBrasilAPI(cep, ch)

	select {
	case res := <-ch:
		fmt.Printf("%s\nLogradouro: %s\nBairro: %s\nCidade: %s\nEstado: %s\n", res.Service, res.Street, res.Neighborhood, res.City, res.State)

	case <-time.After(time.Second):
		fmt.Println("Não foi possível localizar o CEP, o servidor não respondeu dentro do tempo esperado.")
	}

}
