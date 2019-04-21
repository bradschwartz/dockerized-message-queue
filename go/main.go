package main

import (
  "fmt"
  "log"
  "time"
  "github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
  if err != nil {
    log.Fatalf("%s: %s", msg, err)
  }
}

func getConnection(server string) (*amqp.Connection, error) {
	var amqp_uri string = fmt.Sprintf("amqp://guest:guest@%s:5672/", server)
	var max_tries int = 5
	conn, err := amqp.Dial(amqp_uri)
	if err != nil {
		for max_tries > 0 {
			conn, err := amqp.Dial(amqp_uri)
			if err != nil {
				log.Printf("Failed to connect to %s", server)
				time.Sleep(15 * time.Second)
				max_tries -= 1
			} else {
				return conn, err
			}
		}
	}
	failOnError(err, "Exceed max try outs, couldn't connect to RabbitMQ")
	return conn, err
}

func main() {
	const server string = "mqserver"
	conn, err := getConnection(server)
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	const queue string = "worker"
	q, err := ch.QueueDeclare(
	  queue, // name
	  false,   // durable
	  false,   // delete when unused
	  false,   // exclusive
	  false,   // no-wait
	  nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
	  q.Name, // queue, defined above
	  "",     // consumer
	  true,   // auto-ack
	  false,  // exclusive
	  false,  // no-local
	  false,  // no-wait
	  nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)
	go func() { // `go` starts goroutine, managed multithreading
  	for d := range msgs {
    	log.Printf(" [Go] Received a message: %s", d.Body)
  	}
	}()
	log.Printf(" [Go] Waiting for messages. To exit press CTRL+C")
	<-forever
}