package cache

import (
	"log"

	"github.com/streadway/amqp"
)

// This package level variable will hold the connection to our RabbitMQ instance
var conn *amqp.Connection

func init() {
	var err error

	conn, err = amqp.Dial("amqp://guest:guest@192.168.99.100:5672/")
	if err != nil {
		log.Fatalf("could not connect to rabbitmq: %v", err)
	}
}

func subscribe(qName string) (<-chan amqp.Delivery, func(), error) {
	// create a channel through which we publish
	ch, err := conn.Channel()
	if err != nil {
		return nil, nil, err
	}
	// assert that the queue exists (creates a queue if it doesn't)
	q, err := ch.QueueDeclare(qName, false, false, false, false, nil)
	// create a channel in go, through which incoming messages will be received
	c, err := ch.Consume(q.Name, "", false, false, false, false, nil)
	// return the created channel
	return c, func() { ch.Close() }, err
}
