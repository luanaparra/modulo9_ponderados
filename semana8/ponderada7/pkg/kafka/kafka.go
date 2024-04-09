package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	godotenv "github.com/joho/godotenv"
	"log"
	"os"
	mongo "pkg/mongo"
)

func main() {

	err := godotenv.Load("../../.env")
	if err != nil {
		fmt.Printf("Error loading .env file: %s", err)
	}

	configMap := &kafka.ConfigMap{
		"bootstrap.servers":  os.Getenv("CONFLUENT_BOOTSTRAP_SERVER_SASL"),
		"security.protocol":  "SASL_SSL",
		"sasl.mechanisms":    "PLAIN",
		"sasl.username":      os.Getenv("CLUSTER_API_KEY"),
		"sasl.password":      os.Getenv("CLUSTER_API_SECRET"),
		"session.timeout.ms": 45000,
		"group.id":           "go-group-1",
		"auto.offset.reset":  "latest",
	}

	consumer, err := kafka.NewConsumer(configMap)

	if err != nil {
		log.Printf("Error creating kafka consumer: %v", err)
	}

	consumer.SubscribeTopics([]string{"nicola"}, nil)

	fmt.Println("Kafka consumer up and running!")

	run := true
	for run {
		e := consumer.Poll(1000)
		switch ev := e.(type) {
		case *kafka.Message:
			// application-specific processing
			fmt.Printf("Consumed event from topic %s: key = %-10s value = %s\n",
				*ev.TopicPartition.Topic, string(ev.Key), string(ev.Value))

		case kafka.Error:
			fmt.Fprintf(os.Stderr, "%% Error: %v\n", ev)
			run = false
		}
	}

	consumer.Close()
}