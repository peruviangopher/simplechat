package consumer

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"simplechat/chat"
)

func Consume(room *chat.Room) {
	// 1. Connect to RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@host.docker.internal:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	// 2. Open a channel
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	// 3. Declare the same queue (must match producer)
	q, err := ch.QueueDeclare(
		room.GetID(), // name
		false,        // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	// 4. Consume messages
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	// 5. Create a channel to block main
	forever := make(chan bool)

	go func() {
		for msg := range msgs {
			room.SendExternalMsg(msg.Body)
		}
	}()

	log.Println(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
