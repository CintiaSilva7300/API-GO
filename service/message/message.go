package message

import (
	"database/sql"
	"encoding/json"
	"log"

	"API-GO/config"
	"API-GO/domain"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func saveUserToPostgres(db *sql.DB, user domain.USER) {
	err := domain.InsertUser(db, user)
	if err != nil {
		log.Printf("Failed to insert user: %s", err)
		return
	}
	log.Printf("User saved to PostgreSQL: %v", user)
}

func handleMessage(db *sql.DB, d amqp.Delivery) {
	log.Printf("Received a message: %s", d.Body)

	var user domain.USER
	err := json.Unmarshal(d.Body, &user)
	if err != nil {
		log.Printf("Failed to unmarshal JSON: %s", err)
		return
	}

	saveUserToPostgres(db, user)
}

func StartMessageListener(db *sql.DB) {
	conn, err := amqp.Dial(config.GetRabbitMQConnectionString())
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"user", // name
		false,  // durable
		false,  // delete when unused
		false,  // exclusive
		false,  // no-wait
		nil,    // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			go handleMessage(db, d) // Process each message in a new goroutine
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
