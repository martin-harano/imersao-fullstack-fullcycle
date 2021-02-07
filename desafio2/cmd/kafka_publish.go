package cmd

import (
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/martin-harano/imersao-fullstack-fullcycle/desafio2/application/kafka"
	"github.com/spf13/cobra"
)

var userName string
var userEmail string

var addUserCMD = &cobra.Command{
	Use:   "adduser",
	Short: "Publish request to add user through kafka",
	Run: func(cmd *cobra.Command, args []string) {
		deliveryChan := make(chan ckafka.Event)
		producer := kafka.NewKafkaProducer()

		kafka.Publish(userName, userEmail, producer, deliveryChan)

		kafka.DeliveryReport(deliveryChan)

	},
}

func init() {
	rootCmd.AddCommand(addUserCMD)
	addUserCMD.Flags().StringVarP(&userName, "name", "n", "", "User name")
	addUserCMD.Flags().StringVarP(&userEmail, "email", "e", "", "User email")
}
