package main

// 发送者

import (
	"github.com/streadway/amqp"
	"log"
)


// 辅助函数来检查每个amqp调用的返回值
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}


func main() {
	// 1. 尝试连接RabbitMQ，建立连接
	// 该连接抽象了套接字连接，并为我们处理协议版本协商和认证等。
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// 连接抽象了socket连接，并为我们处理协议版本协商和认证等。接下来，我们创建一个通道，这是大多数用于完成任务的API所在的位置：
	// 2. 接下来，我们创建一个通道，大多数API都是用过该通道操作的。
	ch, err := conn.Channel()
	failOnError(err, "Failed to connect to channel")
	defer ch.Close()

	// 要发送，我们必须声明要发送到的队列。然后我们可以将消息发布到队列：
	// 3. 声明消息要发送到的队列
	q, err := ch.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil,
		)
	failOnError(err, "Failed to declare a queue")

	body := "Hello World!"
	// 4.将消息发布到声明的队列

	err = ch.Publish(
		"",
		q.Name,     // routing key
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:		 []byte(body),
		})
	failOnError(err, "Failed to publish a message")
	// 声明队列是幂等的——仅当队列不存在时才创建。消息内容是一个字节数组，因此你可以在此处编码任何内容。
}