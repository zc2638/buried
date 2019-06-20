package buriedmq

func NewPublish(d MqPublishData) WorkerOption {
	return func(w *mqWorker) {
		w.publish = d
	}
}

func NewConsume(d MqConsumeData) WorkerOption {
	return func(w *mqWorker) {
		w.consume = d
	}
}