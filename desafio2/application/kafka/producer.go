package kafka

import (
	"fmt"
	"os"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	appmodel "github.com/martin-harano/imersao-fullstack-fullcycle/desafio2/application/model"
)

func NewKafkaProducer() *ckafka.Producer {
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": os.Getenv("kafkaBootstrapServers"),
	}
	p, err := ckafka.NewProducer(configMap)
	if err != nil {
		panic(err)
	}
	return p
}

func Publish(name string, email string, producer *ckafka.Producer, deliveryChan chan ckafka.Event) error {

	user := appmodel.User{
		Name:  name,
		Email: email,
	}

	msg, err := user.ToJson()
	if err != nil {
		return err
	}

	topic := os.Getenv("kafkaUserResgistrationTopic")

	message := &ckafka.Message{
		TopicPartition: ckafka.TopicPartition{Topic: &topic, Partition: ckafka.PartitionAny},
		Value:          []byte(msg),
	}
	err = producer.Produce(message, deliveryChan)
	if err != nil {
		return err
	}
	return nil
}

func DeliveryReport(deliveryChan chan ckafka.Event) {
	for e := range deliveryChan {
		switch ev := e.(type) {
		case *ckafka.Message:
			if ev.TopicPartition.Error != nil {
				fmt.Println("Delivery failed:", ev.TopicPartition)
			} else {
				fmt.Println("Delivered message to:", ev.TopicPartition)
			}
		}
	}
}
