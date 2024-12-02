package conf

import (
	"bytes"
	"fmt"

	log "github.com/sirupsen/logrus"

	_ "embed"

	"github.com/spf13/viper"
)

//go:embed dev.env
var dev string

// EnvName type constraint for Environment types
type EnvName string

// Environment types
const (
	DEV  EnvName = "DEV"
	PROD EnvName = "PROD"
	TEST EnvName = "TEST"
)

// Environment types to yaml mapper
var mapper = map[EnvName]string{
	DEV: dev,
}

// Configuration - structure that contains Configuration information from config variables
type Configuration struct {
	DBDsn     string
	HttpPort  string
	TigerPort string
}

var config *Configuration
var envName EnvName

// LoadConfig - loads configurations from config variables into Environment struct
func LoadConfig(conf string) *Configuration {
	envName = EnvName(conf)
	switch envName {
	case DEV:
		viper.SetConfigType("env")
		err := viper.ReadConfig(bytes.NewBuffer([]byte(mapper[envName])))
		if err != nil {
			log.Fatal("Invalid config")
		}
	case PROD:
		viper.AutomaticEnv()

	default:
		log.Fatal("Environment variable not supplied")
	}

	config = &Configuration{
		DBDsn: fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			viper.GetString("POSTGRES_HOST"),
			viper.GetString("POSTGRES_USER"),
			viper.GetString("POSTGRES_PASSWORD"),
			viper.GetString("POSTGRES_DB"),
			viper.GetString("POSTGRES_PORT"),
		),
		HttpPort:  viper.GetString("HTTP_PORT"),
		TigerPort: viper.GetString("TIGER_PORT"),
	}

	return config
}

// Get config
func Get() *Configuration {
	if config == nil {
		log.Fatal("Env not initialized")
	}
	return config
}

// LoadTestConfig - Helper for calling in tests
func LoadTestConfig() *Configuration {
	return LoadConfig(string(DEV))
}

// IsDevEnvironment to check if the current env is dev environment
func IsDevEnvironment() bool {
	return envName == DEV
}

// IsTestEnvironment to check if the current env is test environment
func IsTestEnvironment() bool {
	return envName == TEST
}
