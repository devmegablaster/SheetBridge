package broker

import (
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/devmegablaster/SheetBridge/internal/config"
)

type KafkaProducer struct {
	producer  *kafka.Producer
	topic     string
	partition int32
}

func NewKafkaProducer(topic string, partition int32, cfg config.KafkaConfig) *KafkaProducer {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": strings.Split(cfg.Broker, ":")[0],
	})
	if err != nil {
		slog.Error("Unable to create kafka producer")
		os.Exit(1)
	}

	// INFO: Delivery report handler for produced messages
	go func() {
		for e := range producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					slog.Info(fmt.Sprintf("Failed Produce: %v\n", ev.TopicPartition))
				} else {
					slog.Info(fmt.Sprintf("Produced Successfully: %v\n", ev.TopicPartition))
				}
			}
		}
	}()

	return &KafkaProducer{
		producer:  producer,
		topic:     topic,
		partition: partition,
	}
}

func (p *KafkaProducer) Produce(data []byte) {
	if err := p.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &p.topic,
			Partition: p.partition,
		},
		Value: data,
	}, nil); err != nil {
		slog.Error("Unable to produce message")
	}
}
