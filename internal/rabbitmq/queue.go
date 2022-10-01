package rabbitmq

import (
	"github.com/streadway/amqp"
	"log"
	"test/internal/config"
)

type mq struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func Init() {
	mq := new(mq)
	mqConn, err := amqp.Dial(config.RABBITMQ_HOST)

	if err != nil {
		panic(err)
	}
	mq.conn = mqConn

	mq.OpenChannel()
	mq.Consume()
}

func (mq *mq) OpenChannel() {
	log.Println("RABBITMQ_CONNECT", mq.conn)
	channelRabbitMQ, err := mq.conn.Channel()
	if err != nil {
		panic(err)
	}
	mq.channel = channelRabbitMQ

	log.Println("Successfully channel opened")
}

func (mq *mq) CloseChannel() {
	defer mq.channel.Close()
	defer mq.conn.Close()
}

func (mq *mq) Consume() {
	messages, err := mq.channel.Consume(
		config.RABBITMQ_CONSUME_QUEUE, // queue name
		"",                            // consumer
		true,                          // auto-ack
		false,                         // exclusive
		false,                         // no local
		false,                         // no wait
		nil,                           // arguments
	)
	if err != nil {
		log.Println(err)
	}
	log.Println("Successfully connected to RabbitMQ")

	consumer := make(chan bool)

	go func() {
		for message := range messages {
			// For example, show received message in a console.
			log.Printf(" > Received message: %s\n", message.Body)
		}
	}()

	<-consumer
}
