package main

// 接收者

import (
	"github.com/streadway/amqp"
	"log"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}


func main() {
	// 建立连接
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// 获取channel
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// 声明队列
	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")
	// 请注意，我们也在这里声明队列。因为我们可能在发布者之前启动使用者，所以我们希望在尝试使用队列中的消息之前确保队列存在。

	// 我们将告诉服务器将队列中的消息传递给我们。由于它将异步地向我们发送消息，因此我们将在goroutine中从通道（由amqp::Consume返回）中读取消息。
	msgs, err := ch.Consume(
		q.Name,      // queue
		"",
		true,
		false,
		false,
		false,
		nil,
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
}