package main

import (
	"log"
	"time"

	"github.com/damiannolan/eventing-init/config"
	"github.com/damiannolan/eventing-init/eventing"
	"github.com/jpillora/backoff"
)

func main() {
	log.Println("Starting InitContainer...")
	// Backoff is a time.Duration counter. It starts at Min. After every call to Duration() it is multiplied by Factor.
	// It is capped at Max. It returns to Min on every call to Reset(). Jitter can be used to add randomness.
	backoff := &backoff.Backoff{
		Min:    1 * time.Second,
		Max:    30 * time.Second,
		Factor: 2,
		Jitter: false,
	}

	clusterAdmin := eventing.WaitForKafka(backoff, []string{config.KafkaSettings.Broker()}, config.SaramaConfig)
	defer clusterAdmin.Close()

	backoff.Reset()

}
