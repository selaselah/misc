package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/streadway/amqp"
)

var SEVERITIES = map[string]string{
	"INFO":    "INFO",
	"WARN":    "WARNING",
	"WARNING": "WARNING",
	"ERROR":   "ERROR",
	"ERR":     "ERROR",
}

func failOnError(err error, msg string) {
	if err != nil {
		err_msg := fmt.Sprintf("%s: %s", msg, err)
		log.Fatalln(err_msg)
		panic(err_msg)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"logs_direct", // name
		"direct",      // type
		true,          // durable
		false,         // auto-delete
		false,         // internal
		false,         // no-wait
		nil,           // args
	)
	failOnError(err, "Failed to declare an exchange")

	severity := severityFrom(os.Args)
	body := bodyFrom(os.Args)
	err = ch.Publish(
		"logs_direct", // exchange name
		severity,      // routing key
		false,         // mandatory
		false,         // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publishing message")

	log.Printf(" [x] %s: %s", severity, body)
}

func bodyFrom(args []string) (s string) {
	if len(args) < 2 {
		s = "<empty>"
		return
	}

	// args[body_pos] is the log body
	body_pos := 2
	_, has_severity := SEVERITIES[strings.ToUpper(args[1])]
	if has_severity == false {
		body_pos = 1
	}

	if len(args) <= body_pos || strings.TrimSpace(args[body_pos]) == "" {
		s = "<empty>"
		return
	} else {
		s = strings.Join(args[body_pos:], " ")
		return
	}
}

func severityFrom(args []string) string {
	if len(args) < 2 {
		return "INFO"
	}

	severity, has_severity := SEVERITIES[strings.ToUpper(args[1])]
	if has_severity == false {
		return "INFO"
	} else {
		return severity
	}
}
