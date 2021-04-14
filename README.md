# Fluent Resp Api
Este é um projeto destinado a prática da linguagem GO.
Tem como proposta o desenvolvimento de um pacote que visa facilitar a implementação de uma Web API Rest através do padrão [FluentInterface](https://martinfowler.com/bliki/FluentInterface.html)

## Exemplo de como utilizar
```
package main

import (
	"net/http"

	"github.com/welberty/fluentrestapi"
)

func main() {
	type MyDto struct {
		Body string
	}
	h := func(req fluentrestapi.Request) fluentrestapi.HandlerResponse {
		return fluentrestapi.HandlerResponse{StatusCode: http.StatusOK, JsonResponse: MyDto{Body: "test"}}
	}
	hn := func(req fluentrestapi.Request) fluentrestapi.HandlerResponse {
		return fluentrestapi.
			HandlerResponse{
			StatusCode: http.StatusOK,
			JsonResponse: MyDto{
				Body: "Olá " + req.Params["Name"]}}
	}
	err := fluentrestapi.
		Init().
		AddHealthCheck().
		Get("/Test", h).
		Get("/Test2", h).
		Get("/Ola/{Name}", hn).
		Start(9000, http.Server{})
	if err != nil {
		panic("Ops algo deu errado" + err.Error())
	}
}

```
