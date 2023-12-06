package kafka

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/IBM/sarama"
)

const (
	broker = "10.25.16.10:9092"
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
	}
	//定时器
	ticker := time.NewTicker(2 * time.Second)

	// 启动一个无限循环，等待定时器事件
	i := 0
	for {
		select {
		case <-ticker.C:
			i++
			message.Value = sarama.StringEncoder(strconv.Itoa(i) + "-Hello, Kafka!")
			_, _, err = producer.SendMessage(message)
			if err != nil {
				t.Fatalf("Failed to send message: %s", err)
			}
			fmt.Println("send success")
		}
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

	// 使用通道来通知测试结束
	done := make(chan bool)

	// 在goroutine中持续消费消息
	go func() {
		for {
			select {
			case msg := <-partitionConsumer.Messages():
				receivedData := string(msg.Value)
				fmt.Println("receivedData:", receivedData)
			case err := <-partitionConsumer.Errors():
				fmt.Println("Error consuming message:", err)
			}
		}
	}()

	// 让测试保持持续运行
	select {
	case <-done:
		return
	case <-time.After(30 * time.Second): // 设定持续消费的时间，这里设置了30秒
		done <- true // 完成后发送信号，关闭消费者
	}
}
