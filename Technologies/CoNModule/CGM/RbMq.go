package main

import (
	"log"
	"time"

	"github.com/streadway/amqp"
)

// RabbitMQ ... Connections ---------------------------******************---------------------------

// Listen2RabbitMq ...
func Listen2RabbitMq(rabbitChannel chan<- []byte) {
	defer wg.Done()
	conn, err := amqp.Dial("amqp://con:con@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ", rabbitChannel)

	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel", rabbitChannel)
	defer ch.Close()

	msgs, err := ch.Consume(
		CurrentEntity, // queue
		"Sanjay",      // consumer
		true,          // auto-ack
		false,         // exclusive
		false,         // no-local
		false,         // no-wait
		nil,           // args
	)
	if err != nil {
		log.Fatal("Failed to bind a queue")
	}
	go func(rChannel chan<- []byte) {
		for X := range msgs {
			rChannel <- X.Body
		}
	}(rabbitChannel)

	onCloseRabbitMQ := make(chan *amqp.Error)
	go conn.NotifyClose(onCloseRabbitMQ)
	if <-onCloseRabbitMQ != nil {
		time.Sleep(time.Second * 5)
		Listen2RabbitMq(rabbitChannel)
	}

	log.Printf("Listening to RabbitMQ")
	// forever := make(chan bool)
	// <-forever

}
func failOnError(err error, msg string, rChannel chan<- []byte) {
	if err != nil {
		log.Printf("%s: %s", msg, err)
		time.Sleep(time.Second * 5)
		Listen2RabbitMq(rChannel)
	}
}
