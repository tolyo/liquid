package tiger

import (
	"liquid/pkg/conf"
	"log"

	tiger "github.com/tigerbeetle/tigerbeetle-go"
	tigerType "github.com/tigerbeetle/tigerbeetle-go/pkg/types"
)

var instance tiger.Client

func Instance() tiger.Client {
	return instance
}

func SetupInstance() tiger.Client {
	port := conf.Get().TigerPort
	cl, err := tiger.NewClient(tigerType.ToUint128(0), []string{port})
	if err != nil {
		log.Fatalf("Error creating client: %s", err)
	}
	instance = cl
	return cl
}
