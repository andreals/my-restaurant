# Capital Gain
## _Code Challenge: Ganho de Capital_

Capital Gain, ou Ganho de Capital é um "code challenge" oferecido pela Nubank como step de um processo seletivo de Engenharia de Software.

A aplicação CLI tem como responsabilidade calcular o imposto a ser pago sobre lucros ou prejuízos de operações no mercado financeiro de ações.

## Stack

Capital Gain usa como tech stack:
- Golang
- Docker (opcional)
- Make (opcional)

## Tech Decisions

As decisões técnicas utilizadas para o desenvolvimento da aplicação foram de criar uma aplicação que fosse facilmente compreendida, para isso utilizei uma linguagem altamente utilizada na atualidade (golang) e de fácil entendimento.
Além disso, utilizei algumas das boas práticas de clean code e as premissas do [Uber - Go Style Guide](https://github.com/uber-go/guide/blob/master/style.md).
Apliquei o desenvolvimento baseado em TDD para que os casos de testes já informados no code challenge pudessem me ajudar na resolução do problema, a fim de evitar bugs de regressão e facilitar no andamento da solução.
Utilizei um "scaffold" baseado em DDD para que o projeto possa evoluir futuramente, porém de forma bem "lite" pois trata-se de uma aplicação CLI, na qual não há interações com camadas de transport.

## Installation

A aplicação necessite de um ambiente com [Golang](https://go.dev/doc/install) 1.17+ para rodar.

Instale a dependência e rode o comando:
```sh
go run cmd/cli/main.go
```

## Tests

Para rodar os testes da aplicação, basta rodar o comando na raiz do projeto:
```sh
go test ./...
```

## Docker

Capital Gain é muito fácil de instalar e rodar em um container Docker.
Para isso, siga os passos abaixo para ter sucesso:

Para buildar o container:
```sh
docker build -t capital-gain . -f ./build/Dockerfile
```

Para rodar o container:

```sh
docker run -it capital-gain
```

## Makefile

Uma alternativa para rodar, é usando o Makefile que encontra-se na raiz do projeto.

Comandos existentes:
```sh
make run
```
Esse comando vai rodar a aplicação Capital Gain em Go na sua máquina local.

```sh
make build
```
Esse comando vai gerar um binário da aplicação Capital Gain para ser executado em qualquer local.

```sh
make test
```
Esse comando vai rodar todos os testes existentes na aplicação Capital Gain para validar que tudo esteja OK com sua execução (inclusive os casos de teste do code challenge).

```sh
make docker-build
```
Esse comando vai buildar uma nova imagem do docker com o ambiente pronto para rodar a aplicação Capital Gain.

```sh
make docker-run
```
Esse comando vai rodar a imagem do docker gerada para executar a aplicação Capital Gain.

## Notes

Por se tratar de uma linguagem em que não há uma "regra" de arquitetura, utilizei algumas premissas da comunidade e aderente à algumas boas práticas de mercado, nas quais venho aprimorando desde 2018 quando tive o primeiro contato com a linguagem em um monolito.

## License

MIT