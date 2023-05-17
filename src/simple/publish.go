package simple

import (
	"context"
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

type Content struct {
	Message string `json:"message"`
}

func Publish(content *Content) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Panicf("%s: %s", err, "Failed to connect to RabbitMQ")
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Panicf("%s: %s", err, "Failed to open a channel")
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"q.simple.event",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Panicf("%s: %s", err, "Failed to declare a queue")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	bodyContent, err := json.Marshal(content)
	if err != nil {
		log.Panicf("%s: %s", err, "Failed to parser json")
	}

	err = ch.PublishWithContext(
		ctx,
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "application/json",
			Body:         bodyContent,
		},
	)
	if err != nil {
		log.Printf("%s: %s", err, "Failed to publish a message")
	}
}
