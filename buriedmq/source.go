package buriedmq

import "github.com/streadway/amqp"

// 队列
type MqQueue struct {
	Name       string     // 队列名称，默认default
	Durable    bool       // 是否持久化
	AutoDelete bool       // 是否自动删除
	Exclusive  bool       // 是否设置排他，true队列仅对首次声明他的连接可见，并在连接断开时自动删除
	NoWait     bool       // 是否非阻塞，true非阻塞 不等待返回，false阻塞等待返回信息
	Args       amqp.Table // 参数
	exist      bool       // 是否声明过
}

// 交换机
type MqExchange struct {
	Name       string
	Kind       string
	Durable    bool
	AutoDelete bool
	Internal   bool
	NoWait     bool
	Args       amqp.Table
	exist      bool
}

type MqPublishData struct {
	Exchange  string // 交换机名称，默认空
	Key       string // route_key
	Mandatory bool
	Immediate bool
}

type MqConsumeData struct {
	Queue     string //队列名称，默认default
	Consumer  string
	AutoAck   bool
	Exclusive bool
	NoLocal   bool
	NoWait    bool
	Args      amqp.Table
}

// 接收通道
type MqResult <-chan amqp.Delivery

type WorkerOption func(w *mqWorker)