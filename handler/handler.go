package handler

import (
	httpInterface "github.com/yoed/yoed-http-interface"
	"log"
)

type Handler struct {
	Config httpInterface.Config
}

func (c *Handler) Handle(username string) {
	log.Printf("Yo'ed by %s", username)
}

func New() *Handler {

	c := &Handler{}

	if err := httpInterface.LoadConfig("./config.json", &c.Config); err != nil {
		log.Fatalf("failed loading config: %s", err)
	}

	return c
}