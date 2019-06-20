package main

import (
	"github.com/zc2638/buried"
	"github.com/zc2638/buried/buriedmq"
	"strconv"
)

func main() {

	mq := buriedmq.NewMqSystem("localhost:5672", "guest", "guest")
	defer mq.Close()

	worker := mq.NewWorker()

	l := buried.NewLog(
		buried.NewSystem(worker),
	)

	for i := 0; i < 50; i ++ {
		l.Println("hello " + strconv.Itoa(i))
	}
}
