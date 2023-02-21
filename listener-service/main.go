package main

import (
	"fmt"
	"listener-service/event"
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	// Connect to Rabbit MQ
	rabbitConn, err := connect()
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
	}

	defer rabbitConn.Close()

	// Listen for message
	log.Println("Listening for and consuming RabbitMQ messages...")

	// Create consumer
	consumer, err := event.NewConsumer(rabbitConn)
	if err != nil {
		log.Fatalf("Failed to create consumer: %s", err)
	}

	// Watch queue and consume events
	err = consumer.Listen([]string{"log.INFO", "log.WARNING", "log.ERROR"})
	if err != nil {
		log.Println(err)
	}

}

func connect() (*amqp.Connection, error) {
	connString := fmt.Sprintf("amqp://%s:%s@%s", os.Getenv("RABBITMQ_USER"), os.Getenv("RABBITMQ_PASSWORD"), os.Getenv("RABBITMQ_HOST"))
	c, err := amqp.Dial(connString)
	if err != nil {
		return nil, err
	}

	return c, nil
}
