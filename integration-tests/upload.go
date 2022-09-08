package main

import "github.com/Eun/go-hit"

const (
	host     = "localhost:8080"         // Attempts connection
	basePath = "http://" + host + "/v1" // HTTP REST
)

//TODO remake into real tests

func main() {
	hit.MustDo(
		hit.Post(basePath + "/tinkoff/transactions/update"),
	)
}
