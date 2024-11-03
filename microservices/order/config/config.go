package config

import (
	"log"
	"os"
	"strconv"
)

func GetEnv() string {
	return getEnvironmentValue("ENV")
}

func GetDataSourceURL() string {
	return getEnvironmentValue("DATA_SOURCE_URL")
	//return "root:verysecretpass@tcp(172.17.0.3:3306)/order"
}

func GetApplicationPort() int {
	portStr := getEnvironmentValue("APPLICATION_PORT") // order service port
	port, err := strconv.Atoi(portStr)

	if err != nil {
		log.Fatalf("port: %s is invalid", portStr)
	}
	return port
}

func getEnvironmentValue(key string) string {
	if os.Getenv(key) == "" {
		log.Fatalf("%s environment variable is missing.", key)
	}

	return os.Getenv(key)
}
