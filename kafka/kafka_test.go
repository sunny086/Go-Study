package kafka

import (
	"testing"
	"time"

	"github.com/IBM/sarama"
)

const (
	broker = "10.25.16.130:9092"
	topic  = "test-topic"
)

func TestProducer(t *testing.T) {
	producer, err := sarama.NewSyncProducer([]string{broker}, nil)
	if err != nil {
		t.Fatalf("Error creating producer: %s", err)
	}
	defer func() {
		if err := producer.Close(); err != nil {
			t.Fatalf("Error closing producer: %s", err)
		}
	}()

	message := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder("Hello, Kafka!"),
	}

	_, _, err = producer.SendMessage(message)
	if err != nil {
		t.Fatalf("Failed to send message: %s", err)
	}
}

func TestConsumer(t *testing.T) {
	consumer, err := sarama.NewConsumer([]string{broker}, nil)
	if err != nil {
		t.Fatalf("Error creating consumer: %s", err)
	}
	defer func() {
		if err := consumer.Close(); err != nil {
			t.Fatalf("Error closing consumer: %s", err)
		}
	}()

	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		t.Fatalf("Error creating partition consumer: %s", err)
	}
	defer func() {
		if err := partitionConsumer.Close(); err != nil {
			t.Fatalf("Error closing partition consumer: %s", err)
		}
	}()

	select {
	case <-partitionConsumer.Messages():
		// Successfully consumed a message
	case err := <-partitionConsumer.Errors():
		t.Fatalf("Error consuming message: %s", err)
	case <-time.After(5 * time.Second):
		t.Fatal("Timeout waiting for message")
	}
}

func TestSendMsg(t *testing.T) {
	config := sarama.NewConfig()
	producer, err := sarama.NewSyncProducer([]string{broker}, config)
	if err != nil {
		t.Fatalf("Error creating producer: %s", err)
	}
	defer func() {
		if err := producer.Close(); err != nil {
			t.Fatalf("Error closing producer: %s", err)
		}
	}()

	message := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder("Test message"),
	}

	_, _, err = producer.SendMessage(message)
	if err != nil {
		t.Fatalf("Failed to send message: %s", err)
	}
}

func TestSendReceivedData(t *testing.T) {
	config := sarama.NewConfig()
	consumer, err := sarama.NewConsumer([]string{broker}, config)
	if err != nil {
		t.Fatalf("Error creating consumer: %s", err)
	}
	defer func() {
		if err := consumer.Close(); err != nil {
			t.Fatalf("Error closing consumer: %s", err)
		}
	}()

	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		t.Fatalf("Error creating partition consumer: %s", err)
	}
	defer func() {
		if err := partitionConsumer.Close(); err != nil {
			t.Fatalf("Error closing partition consumer: %s", err)
		}
	}()

	message := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder("Test message"),
	}

	producer, err := sarama.NewSyncProducer([]string{broker}, nil)
	if err != nil {
		t.Fatalf("Error creating producer: %s", err)
	}
	defer func() {
		if err := producer.Close(); err != nil {
			t.Fatalf("Error closing producer: %s", err)
		}
	}()

	_, _, err = producer.SendMessage(message)
	if err != nil {
		t.Fatalf("Failed to send message: %s", err)
	}

	select {
	case msg := <-partitionConsumer.Messages():
		receivedData := string(msg.Value)
		expectedData := "Test message"
		if receivedData != expectedData {
			t.Fatalf("Received data doesn't match, expected: %s, received: %s", expectedData, receivedData)
		}
	case err := <-partitionConsumer.Errors():
		t.Fatalf("Error consuming message: %s", err)
	case <-time.After(5 * time.Second):
		t.Fatal("Timeout waiting for message")
	}
}
