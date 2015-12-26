package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	// "time"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		err_msg := fmt.Sprintf("%s: %s", msg, err)
		log.Fatalln(err_msg)
		panic(err_msg)
	}
}

func bodyFrom(args []string) (s string) {
	if len(args) < 1 || args[0] == "" {
		s = "hello"
	} else {
		s = strings.Join(args, " ")
	}
	return
}

func main() {
	conn, err := amqp.Dial("amqp://localhost:5672")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"logs",   // name
		"fanout", // kind
		true,     // durable
		false,    // auto-delete
		false,    // internal
		false,    // no-wait
		nil,      // args
	)
	failOnError(err, "Failed to declare an exchange")

	body := bodyFrom(os.Args[1:])
	err = ch.Publish(
		"logs", // exchange
		"",     // key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publishing message")

	// time.Sleep(10 * time.Second)
}
