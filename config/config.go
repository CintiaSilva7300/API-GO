package config

import (
	"fmt"
	"os"
)

func GetDBConnectionString() string {
	envUser := os.Getenv("POSTGRES_USER")
	envPass := os.Getenv("POSTGRES_PASS")
	envDB := os.Getenv("POSTGRES_DB")
	envHost := os.Getenv("POSTGRES_HOST")
	envPort := os.Getenv("POSTGRES_PORT")
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", envUser, envPass, envHost, envPort, envDB)
}

func GetRabbitMQConnectionString() string {
	envUser := os.Getenv("RABBITMQ_USER")
	envPass := os.Getenv("RABBITMQ_PASSWORD")
	envHost := os.Getenv("RABBITMQ_URL") // Use RABBITMQ_URL para pegar a URL completa
	return fmt.Sprintf("amqp://%s:%s@%s/", envUser, envPass, envHost)

}
