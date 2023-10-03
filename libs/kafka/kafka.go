package kafka

import (
	"context"
	json2 "encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

const (
	Topic     string = "test-topic1"
	Partition int    = 0
)

var conn *kafka.Conn

type Queue struct{}

func InitializeKafka() {
	connection, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", Topic, Partition)
	if err != nil {
		log.Fatal("Failed to connect with Kafka Broker: ", err)
	}

	if err := connection.SetWriteDeadline(time.Now().Add(10 * time.Second)); err != nil {
		log.Fatal("Failed to write health check to Kafka Broker: ", err)
	}

	//_, err = connection.WriteMessages(
	//	kafka.Message{Value: []byte("Kafka broker is running well..")},
	//)
	//
	//if err != nil {
	//	log.Fatal("failed to write messages:", err)
	//}

	conn = connection
}

func (queue Queue) Consume() error {

	return nil
}

func (queue Queue) Produce(message string) error {
	msg := make(map[string]interface{})
	type Messages struct {
		Message string
	}

	a := "message"

	msg[a] = message
	msg1 := Messages{Message: "잉구잉구다. 엥구엥구당"}

	json, err := json2.Marshal(msg1)
	if err != nil {
		fmt.Println("Failed to transform to json", err)
	}

	if _, err := conn.WriteMessages(kafka.Message{Value: json}); err != nil {
		log.Fatal("failed to write messages:", err)
	}

	return nil
}
