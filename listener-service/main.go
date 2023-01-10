package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func connect() (*amqp.Connection, error) {
	var counts int64
	backoff := 1 * time.Second
	var connection *amqp.Connection

	for {
		c, err := amqp.Dial("amqp://guest:guest@localhost")
		if err != nil {
			fmt.Println("RabbitMQ not yet ready...")
			counts++
		} else {
			connection = c
			break
		}
		if counts > 5 {
			fmt.Println(err)
			return nil, err
		}

		backoff = (time.Duration(math.Pow(float64(backoff), 2)) * time.Second)
		log.Println("backing off..")
		time.Sleep(backoff)
		continue
	}
	return connection, nil
}

func main() {
	//connect to rabbit-mq
	rabbitConn, err := connect()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer rabbitConn.Close()
	log.Println("Connected to RabbitMQ!")

	//start listening for messages
	//create consumer
	//watch queue and consume events
}
