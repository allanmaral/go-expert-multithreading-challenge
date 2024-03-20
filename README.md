
# Desafio Go - Multithreading

Esse repositório, contém a solução para o desafio proposto, que envolve utilizar multithreading para buscar o resultado mais rápido entre duas APIs distintas. Para isso, duas requisições são feitas simultaneamente para as APIs da [Brasil API](https://brasilapi.com.br) e da [ViaCEP](http://viacep.com.br), acatando a API que responder mais rápida e descartando a resposta da mais lenta e, caso nenhuma das APIs retornem a resposta em menos de 1 segundo um erro de timeout é exibido para o usuário.

## Pré-requisitos

Antes de executar os sistemas, certifique-se de ter o Go instalado em sua máquina. Você pode baixá-lo em [https://go.dev/dl/](https://go.dev/dl/).

## Como executar

1. Compile a aplicação:

   ```bash
   go build -o cep cmd/cli/main.go
   ```

2. Consulte o CEP:

   ```bash
   ./cep 01153000
   ```

3. Você deve ver algo como:
   
   ```
   $ ./cep 01153000
   ViaCEP
   Logradouro: Rua Vitorino Carmilo
   Bairro: Barra Funda
   Cidade: São Paulo
   Estado: SP
   ```

## Estrutura do Projeto

- `cmd/cli/main.go`: Entrypoint da aplicação, cliente que irá ler o cep dos argumentos e realizar as chamadas nas APIs.
- `pkg/brasilapi/cep.go`: Implementação do client da API de CEP do Brasil API.
- `pkg/viacep/cep.go`: Implementação do client da API de CEP do ViaCEP.
- `pkg/model/cep.go`: Definição do modelo de CEP usado como abstração entre a resposta das duas APIs.
