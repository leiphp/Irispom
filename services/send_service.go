package services

import (
	"log"
	"strconv"
	"sync"

	"github.com/streadway/amqp"
)

/*
	提供RabbitMQ服务接口
	作者名称：LeiWen 创建时间：20201107
*/
type SendInterfaceService interface {
	SendMsg(num int, wg *sync.WaitGroup) (interface{}, error)           //接收RabbitMQ消息

}

//初始化对象函数
func NewSendService() SendInterfaceService {
	return &sendService{

	}
}

type sendService struct {

}

func (this *sendService) SendMsg(num int, wg *sync.WaitGroup) (interface{}, error) {

	wg.Done()

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	//声明队列
	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	body := "Hello World!,第"+strconv.Itoa(num)+"条消息"

	log.Printf("Send %d message: %s", num, body)
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing {
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")
	return 200, nil
}
