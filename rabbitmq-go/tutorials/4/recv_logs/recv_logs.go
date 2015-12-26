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

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")

	q, err := ch.QueueDeclare(
		"",    // queue name
		false, // durable
		false, // auto-delete
		true,  // exclusive
		false, // no-wait
		nil,   // args
	)
	failOnError(err, "Failed to declare a queue")
	ss := severitiesFrom(os.Args)
	if len(ss) == 0 {
		log.Printf("should at least given one severity")
		os.Exit(0)
	}

	for _, severity := range ss {
		log.Printf("[*] bind with severity=%s", severity)
		err := ch.QueueBind(
			q.Name,        // queue name
			severity,      // routing key
			"logs_direct", // exchange name
			false,         // no-wait
			nil,
		)
		failOnError(err, "Failed to declare a queue")
	}

	msgs, err := ch.Consume(
		q.Name, // queue name
		"",     // consumer name
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("[*] %s", d.Body)
		}
	}()
	log.Printf("[*] Waiting for logs. To exit press CTRL+C")
	<-forever
}

func severitiesFrom(args []string) []string {
	ss := make([]string, 0, 10)
	for _, s := range args[1:] {
		if len(ss) == cap(ss) {
			log.Println("[*] to many args, stop parsing")
			break
		}
		severity, ok := SEVERITIES[strings.ToUpper(strings.TrimSpace(s))]
		if ok == true {
			log.Println("add s")
			ss = ss[0 : len(ss)+1]
			ss[len(ss)-1] = severity
		}
	}
	return ss
}
