package message

import (
	"context"
	"encoding/json"
	"log"

	"API-GO/config"
	"API-GO/domain"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func saveUserToPostgres(ctx context.Context, pool *pgxpool.Pool, user domain.USER) error {
	err := domain.InsertUser(ctx, pool, user)
	if err != nil {
		log.Printf("Failed to insert user: %s", err)
		return err
	}
	log.Printf("User saved to PostgreSQL: %v", user)
	return nil
}

func handleMessage(ctx context.Context, pool *pgxpool.Pool, ch *amqp.Channel, d amqp.Delivery) {
	log.Printf("Received a message: %s", d.Body)

	var user domain.USER
	err := json.Unmarshal(d.Body, &user)
	if err != nil {
		log.Printf("Failed to unmarshal JSON: %s", err) //Se ocorrer um erro, a mensagem é rejeitada e enviada para a fila de dead-letter.
		d.Nack(false, false)                            // Reject and don't requeue the message
		sendToDeadLetterQueue(ch, d)
		return
	}

	err = saveUserToPostgres(ctx, pool, user)
	if err != nil {
		d.Nack(false, false) // Reject and don't requeue the message
		sendToDeadLetterQueue(ch, d)
	} else {
		d.Ack(false) // Acknowledge the message
	}
}

func sendToDeadLetterQueue(ch *amqp.Channel, d amqp.Delivery) { //envia a mensagem para a fila de dead-letter caso ocorra um erro
	err := ch.Publish(
		"dlx_exchange",    // Exchange
		"dlx_routing_key", // Routing key
		false,             // Mandatory
		false,             // Immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        d.Body,
		},
	)
	if err != nil {
		log.Printf("Failed to publish to DLX: %s", err)
	} else {
		log.Printf("Message sent to DLX: %s", d.Body)
	}
}

func StartMessageListener(pool *pgxpool.Pool) {
	conn, err := amqp.Dial(config.GetRabbitMQConnectionString())
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// Declare the DLX exchange
	err = ch.ExchangeDeclare(
		"dlx_exchange", // name
		"direct",       // type
		true,           // durable
		false,          // auto-deleted
		false,          // internal
		false,          // no-wait
		nil,            // arguments
	)
	failOnError(err, "Failed to declare DLX exchange")

	// Declare the DLX queue
	_, err = ch.QueueDeclare(
		"dlx_queue", // name
		true,        // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	failOnError(err, "Failed to declare DLX queue")

	// Bind DLX queue to DLX exchange
	err = ch.QueueBind(
		"dlx_queue",       // queue name
		"dlx_routing_key", // routing key
		"dlx_exchange",    // exchange
		false,
		nil,
	)
	failOnError(err, "Failed to bind DLX queue to DLX exchange")

	// Declare the main queue with DLX parameters
	args := amqp.Table{
		"x-dead-letter-exchange":    "dlx_exchange",
		"x-dead-letter-routing-key": "dlx_routing_key",
	}

	q, err := ch.QueueDeclare(
		"user", // name
		true,   // durable
		false,  // delete when unused
		false,  // exclusive
		false,  // no-wait
		args,   // arguments
	)
	failOnError(err, "Failed to declare main queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			go handleMessage(context.Background(), pool, ch, d) // Process each message in a new goroutine
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
