package broker

import (
	"log/slog"
	"os"
	"strings"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/devmegablaster/SheetBridge/internal/config"
	"github.com/devmegablaster/SheetBridge/pb"
	"google.golang.org/protobuf/proto"
)

type WriteConsumer struct {
	consumer   *kafka.Consumer
	protoWrite chan *pb.Write
}

func NewWriteConsumer(cfg *config.KafkaConfig, protoWrite chan *pb.Write) *WriteConsumer {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": strings.Split(cfg.Broker, ":")[0],
		"group.id":          cfg.WriteGroup,
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		slog.Error("Unable to init write consumer")
		os.Exit(1)
	}

	err = c.Subscribe(cfg.WriteTopic, nil)
	if err != nil {
		slog.Error("Unable to subscribe to write topic")
		os.Exit(1)
	}

	slog.Info("✅Write consumer listening...")

	return &WriteConsumer{
		consumer:   c,
		protoWrite: protoWrite,
	}
}

func (s *WriteConsumer) Consume() {
	for {
		msg, err := s.consumer.ReadMessage(-1)
		if err != nil {
			slog.Error("Unable to process kafka message", slog.String("error", err.Error()))
			continue
		}

		protoWrite := &pb.Write{}
		if err := proto.Unmarshal(msg.Value, protoWrite); err != nil {
			slog.Error("Unable to process kafka message", slog.String("message", string(msg.Value)))
			continue
		}

		s.protoWrite <- protoWrite
	}
}
