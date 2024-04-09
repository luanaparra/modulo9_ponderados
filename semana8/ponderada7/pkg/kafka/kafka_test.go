package main

import (
	"os"
	"testing"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	godotenv "github.com/joho/godotenv"
)

func TestKafkaConnection(t *testing.T) {
	config := setEnvironment(t)

	consumer, err := kafka.NewConsumer(&config)
	if err != nil {
		t.Fatalf("Error creating test consumer: %v", err)
	}
	defer func() {
		if err := consumer.Close(); err != nil {
			t.Fatalf("Error closing test consumer: %v", err)
		}
	}()

	if _, err := consumer.Subscription(); err != nil {
		t.Fatalf("Error fetching topics: %v", err)
	}
}

func TestQueueConsumption(t *testing.T) {

	config := setEnvironment(t)

	consumer, err := kafka.NewConsumer(&config)
	if err != nil {
		t.Fatalf("Error creating test consumer: %v", err)
	}
	defer func() {
		if err := consumer.Close(); err != nil {
			t.Fatalf("Error closing test consumer: %v", err)
		}
	}()

	topic := "teste"

	producer, err := kafka.NewProducer(&config)
	if err != nil {
		t.Fatalf("Error creating test producer: %v", err)
	}
	defer producer.Close()

	deliveryChan := make(chan kafka.Event)
	err = producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte("test_message"),
	}, deliveryChan)
	if err != nil {
		t.Fatalf("Error producing test message: %v", err)
	}

	e := <-deliveryChan
	m := e.(*kafka.Message)
	if m.TopicPartition.Error != nil {
		t.Fatalf("Error producing test message: %v", m.TopicPartition.Error)
	}

	if err := consumer.SubscribeTopics([]string{topic}, nil); err != nil {
		t.Fatalf("Error subscribing to test topic: %v", err)
	}

	msg, err := consumer.ReadMessage(-1)
	if err != nil {
		t.Fatalf("Error consuming test message: %v", err)
	}

	if string(msg.Value) != "test_message" {
		t.Fatalf("Received unexpected message: %s", string(msg.Value))
	}
}

func setEnvironment(t *testing.T) kafka.ConfigMap{

	t.Helper()

	err := godotenv.Load("../../.env")
	if err != nil {
		t.Fatalf("Error loading .env file: %s", err)
	}

	config := kafka.ConfigMap{
		"bootstrap.servers":  os.Getenv("CONFLUENT_BOOTSTRAP_SERVER_SASL"),
		"sasl.mechanisms":    "PLAIN",
		"security.protocol":  "SASL_SSL",
		"sasl.username":      os.Getenv("CLUSTER_API_KEY"),
		"sasl.password":      os.Getenv("CLUSTER_API_SECRET"),
		"session.timeout.ms": 6000,
		"group.id":           "grupo1",
		"auto.offset.reset":  "latest",
	}

	return config
}