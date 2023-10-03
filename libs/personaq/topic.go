package personaq

type Topic struct {
	Topic   string
	Message interface{}
}

func (tp Topic) CreateTopic(topic string, message interface{}) *Topic {
	return &Topic{
		Topic:   "",
		Message: "",
	}
}

func (tp Topic) Connect() {
	ch := make(chan Topic)
	broker := NewBroker()

}

func (tp Topic) Disconnect() {}
