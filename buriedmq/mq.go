package buriedmq

import (
	"errors"
	"github.com/streadway/amqp"
)

const QUEUE_DEFAULT = "default"
const (
	EXCHANGE_FANOUT = "fanout"
	EXCHANGE_DIRECT = "direct"
	EXCHANGE_TOPIC  = "topic"
)

type mqSystem struct {
	host         string
	user         string
	pwd          string
	conn         *amqp.Connection
	ch           *amqp.Channel
	exchangePool map[string]*MqExchange
	queuePool    map[string]*MqQueue
}

func NewMqSystem(host string, user string, pwd string) *mqSystem {
	return &mqSystem{host: host, user: user, pwd: pwd}
}

func (s *mqSystem) Conn() error {

	if s.conn == nil || s.conn.IsClosed() {
		conn, err := amqp.Dial("amqp://" + s.user + ":" + s.pwd + "@" + s.host + "/")
		if err != nil {
			return err
		}
		s.conn = conn

		ch, err := s.conn.Channel()
		if err != nil {
			if err := conn.Close(); err != nil {
				return err
			}
			return err
		}
		s.ch = ch
	}

	if s.queuePool == nil {
		s.QueueDeclare(&MqQueue{})
	}

	return nil
}

func (s *mqSystem) CloseCh() error {

	if s.ch != nil {
		return s.ch.Close()
	}
	return nil
}

func (s *mqSystem) Close() error {

	if s.conn == nil {
		return nil
	}
	if err := s.conn.Close(); err != nil && err != amqp.ErrClosed {
		return err
	}
	return nil
}

// 声明队列池
func (s *mqSystem) QueueDeclare(qds ...*MqQueue) {

	if s.queuePool == nil {
		s.queuePool = make(map[string]*MqQueue)
	}
	for _, qd := range qds {
		if qd.Name == "" {
			qd.Name = QUEUE_DEFAULT
		}

		_, ok := s.queuePool[qd.Name]
		if !ok || qd.Name == QUEUE_DEFAULT {
			s.queuePool[qd.Name] = qd
		}
	}
}

// 检查队列
func (s *mqSystem) checkQueue(name string) error {

	queue, ok := s.queuePool[name]
	if !ok {
		return errors.New(name + " queue is not exist in queuePool")
	}

	if !queue.exist {
		_, err := s.ch.QueueDeclare(queue.Name, queue.Durable, queue.AutoDelete, queue.Exclusive, queue.NoWait, queue.Args)
		if err != nil {
			return err
		}
		s.queuePool[name].exist = true
	}
	return nil
}

// 声明交换机池
func (s *mqSystem) ExchangeDeclare(eds ...*MqExchange) {

	if s.exchangePool == nil {
		s.exchangePool = make(map[string]*MqExchange)
	}
	for _, ed := range eds {
		if ed.Name == "" {
			continue
		}
		_, ok := s.exchangePool[ed.Name]
		if !ok {
			s.exchangePool[ed.Name] = ed
		}
	}
}

// 检查交换机
func (s *mqSystem) checkExchange(name string) error {

	if name == "" {
		return nil
	}
	exchange, ok := s.exchangePool[name]
	if !ok {
		return errors.New(name + " exchange is not exist in exchange pool")
	}

	if !exchange.exist {
		if err := s.ch.ExchangeDeclare(
			exchange.Name,
			exchange.Kind,
			exchange.Durable,
			exchange.AutoDelete,
			exchange.Internal,
			exchange.NoWait,
			exchange.Args,
		); err != nil {
			return err
		}
		exchange.exist = true
	}

	return nil
}

func (s *mqSystem) NewWorker(wps ...WorkerOption) *mqWorker {

	worker := &mqWorker{system: s}
	for _, wp := range wps {
		wp(worker)
	}
	return worker
}