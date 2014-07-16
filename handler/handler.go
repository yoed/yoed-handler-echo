package handler

import (
	clientInterface "github.com/yoed/yoed-client-interface"
	"log"
)

type Handler struct {
	Config clientInterface.Config
}

func (c *Handler) Handle(username string) {
	log.Printf("Yo'ed by %s", username)
}

func New() *Handler {

	c := &Handler{}

	if err := clientInterface.LoadConfig("./config.json", &c.Config); err != nil {
		log.Fatalf("failed loading config: %s", err)
	}

	return c
}