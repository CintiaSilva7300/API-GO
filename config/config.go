package config

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

func GetDBConnectionString() string {
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	pass := os.Getenv("POSTGRES_PASS")
	dbname := os.Getenv("POSTGRES_DB")

	if port == "" {
		panic("POSTGRES_PORT environment variable is not set")
	}

	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, pass, host, port, dbname)
}

func GetDBConnectionPool() (*pgxpool.Pool, error) {
	connStr := GetDBConnectionString()
	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, err
	}

	pool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}

	return pool, nil
}

func GetRabbitMQConnectionString() string {
	envUser := os.Getenv("RABBITMQ_USER")
	envPass := os.Getenv("RABBITMQ_PASSWORD")
	envHost := os.Getenv("RABBITMQ_URL") // Use RABBITMQ_URL para pegar a URL completa
	return fmt.Sprintf("amqp://%s:%s@%s/", envUser, envPass, envHost)
}
