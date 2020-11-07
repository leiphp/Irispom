package services

import (
	"log"

	"github.com/streadway/amqp"
)

/*
	提供RabbitMQ服务接口
	作者名称：LeiWen 创建时间：20201107
*/
type ReceiveInterfaceService interface {
	ReceiveMsg() (interface{}, error)           //接收RabbitMQ消息

}

//初始化对象函数
func NewReceiveService() ReceiveInterfaceService {
	return &receiveService{

	}
}

type receiveService struct {

}

func (this *receiveService) ReceiveMsg() (interface{}, error) {

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

	return 200, nil
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}