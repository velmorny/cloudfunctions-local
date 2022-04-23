package functionspubsub

import (
	"context"
	"log"
)

type PubSubMessage struct {
	Data []byte `json:"data"`
}

func PubSubFunc(ctx context.Context, m PubSubMessage) error {
	name := string(m.Data)
	if name == "" {
		name = "World"
	}
	log.Printf("Hello %s!", name)
	return nil
}
