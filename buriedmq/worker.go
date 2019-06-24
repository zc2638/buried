package buriedmq

import (
	"github.com/streadway/amqp"
	"github.com/zc2638/buried"
)

type mqWorker struct {
	System  *mqSystem
	publish MqPublishData
	consume MqConsumeData
}

func (p *mqWorker) Publish(b []byte) error {

	if err := p.System.Conn(); err != nil {
		return err
	}

	mqData := p.publish
	publishing := amqp.Publishing{
		ContentType: "application/json",
		Body:        b,
	}

	if err := p.System.checkExchange(p.publish.Exchange); err != nil {
		return err
	}

	if mqData.Key == "" {
		mqData.Key = QUEUE_DEFAULT
	}

	if err := p.System.checkQueue(mqData.Key); err != nil {
		return err
	}

	err := p.System.ch.Publish(mqData.Exchange, mqData.Key, mqData.Mandatory, mqData.Immediate, publishing)
	if err != nil {
		return err
	}
	return nil
}

func (p *mqWorker) Consume() (buried.Result, error) {

	if err := p.System.Conn(); err != nil {
		return nil, err
	}

	d := p.consume
	if d.Queue == "" {
		d.Queue = QUEUE_DEFAULT
	}

	err := p.System.checkQueue(d.Queue)
	if err != nil {
		return nil, err
	}

	if err := p.System.checkQueueBind(d.Queue); err != nil {
		return nil, err
	}

	var res MqResult
	res, err = p.System.ch.Consume(d.Queue, d.Consumer, d.AutoAck, d.Exclusive, d.NoLocal, d.NoWait, d.Args)
	return res, err
}
