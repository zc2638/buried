package main

import (
	"buried/buriedmq"
	"fmt"
	"log"
)

func main() {

	mq := buriedmq.NewMqSystem("localhost:5672", "guest", "guest")
	consumeData := buriedmq.MqConsumeData{AutoAck: true}
	worker := mq.NewWorker(
		buriedmq.NewConsume(consumeData),
	)

	res, err := worker.Consume()
	if err != nil {
		log.Fatal(err)
	}

	for m := range res.(buriedmq.MqResult) {
		fmt.Print(string(m.Body))
	}
}
