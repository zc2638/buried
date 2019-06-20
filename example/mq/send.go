package main

import (
	"buried/buriedmq"
	"strconv"
)

func main() {

	mq := buriedmq.NewMqSystem("localhost:5672", "guest", "guest")
	defer mq.Close()

	worker := mq.NewWorker()
	for i := 0; i < 50; i ++ {
		_ = worker.Publish([]byte("hello " + strconv.Itoa(i)))
	}
}
