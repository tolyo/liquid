package models

import (
	"errors"
	"liquid/pkg/db"
	"log"
)

// AppEntityId `app_entity.pub_id` db reference
type AppEntityId string

// AppEntityExternalId `app_entity.pub_id` db reference
type AppEntityExternalId string

// AppEntityType Type of application entity
type AppEntityType string

const (
	Client   AppEntityType = "CLIENT"
	Master   AppEntityType = "MASTER"
	Provider AppEntityType = "PROVIDER"
)

const MASTER = "MASTER"
const PROVIDER1 = "PROVIDER_1"

// AppEntity Application entity is any generic entity capable of being an actor in financial transaction
type AppEntity struct {
	Id         AppEntityId
	Type       AppEntityType
	ExternalId AppEntityExternalId
}

func FindAppEntityExternalId(id AppEntityExternalId) (AppEntityId, error) {
	res := db.QueryVal[string]("SELECT pub_id FROM app_entity WHERE external_id = $1", id)
	if res == "" {
		return "", errors.New("entity not found")
	}
	return AppEntityId(res), nil
}

func GetMaster() AppEntityId {
	res, err := FindAppEntityExternalId(MASTER)
	if err != nil {
		log.Fatal("Master entity not initialized")
	}
	return res
}

func GetProvider() AppEntityId {
	res, err := FindAppEntityExternalId(PROVIDER1)
	if err != nil {
		log.Fatal("Provider not initialized")
	}
	return res
}

func CreateClientEntity(id AppEntityExternalId) AppEntityId {
	var newId string
	err := db.Instance().QueryRow("SELECT create_client_entity($1)", id).Scan(&newId)
	if err != nil {
		log.Fatal("Unable to create client entity")
	}
	return AppEntityId(newId)
}
