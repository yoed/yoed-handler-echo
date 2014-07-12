package main

import (
	clientInterface "github.com/yoed/yoed-client-interface"
	"log"
)

type EchoYoedClient struct {
	clientInterface.BaseYoedClient
}

func (c *EchoYoedClient) Handle(username string) {
	log.Printf("Yo'ed by %s", username)
}

func NewEchoYoedClient() (*EchoYoedClient, error) {
	c := &EchoYoedClient{}
	baseClient, err := clientInterface.NewBaseYoedClient()

	if err != nil {
		return nil, err
	}
	c.BaseYoedClient = *baseClient

	return c, nil
}

func main() {
	c, _ := NewEchoYoedClient()

	clientInterface.Run(c)
}