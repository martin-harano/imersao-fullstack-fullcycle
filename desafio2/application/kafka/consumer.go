package kafka

import (
	"fmt"
	"os"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/jinzhu/gorm"
	"github.com/martin-harano/imersao-fullstack-fullcycle/desafio2/application/factory"
	appmodel "github.com/martin-harano/imersao-fullstack-fullcycle/desafio2/application/model"
)

type KafkaProcessor struct {
	Database     *gorm.DB
	Producer     *ckafka.Producer
	DeliveryChan chan ckafka.Event
}

func NewKafkaProcessor(database *gorm.DB, producer *ckafka.Producer, deliveryChan chan ckafka.Event) *KafkaProcessor {
	return &KafkaProcessor{
		Database:     database,
		Producer:     producer,
		DeliveryChan: deliveryChan,
	}
}

func (k *KafkaProcessor) Consume() {
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": os.Getenv("kafkaBootstrapServers"),
		"group.id":          os.Getenv("kafkaConsumerGroupId"),
		"auto.offset.reset": "earliest",
	}
	c, err := ckafka.NewConsumer(configMap)

	if err != nil {
		panic(err)
	}

	topics := []string{os.Getenv("kafkaUserResgistrationTopic")}
	c.SubscribeTopics(topics, nil)

	fmt.Println("kafka consumer has been started")
	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Println(string(msg.Value))
			k.processMessage(msg)
		}
	}
}

func (k *KafkaProcessor) processMessage(msg *ckafka.Message) {
	kafkaUserResgistrationTopic := os.Getenv("kafkaUserResgistrationTopic")

	switch topic := *msg.TopicPartition.Topic; topic {
	case kafkaUserResgistrationTopic:
		k.processRegistration(msg)
	default:
		fmt.Println("not a valid topic", string(msg.Value))
	}
}

func (k *KafkaProcessor) processRegistration(msg *ckafka.Message) error {
	user := appmodel.NewUser()
	err := user.ParseJson(msg.Value)
	if err != nil {
		return err
	}

	userUseCase := factory.UserUseCaseFactory(k.Database)

	_, err = userUseCase.AddNewUser(
		user.Name,
		user.Email,
	)
	if err != nil {
		fmt.Println("error registering User ->", err)
		return err
	}

	return nil
}
