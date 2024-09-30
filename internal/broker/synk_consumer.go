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

type SynkConsumer struct {
	consumer  *kafka.Consumer
	protoSynk chan *pb.Synk
}

func NewSynkConsumer(cfg *config.KafkaConfig, protoSynk chan *pb.Synk) *SynkConsumer {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": strings.Split(cfg.Broker, ":")[0],
		"group.id":          cfg.SynkGroup,
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		slog.Error("Unable to init synk consumer")
		os.Exit(1)
	}

	err = c.Subscribe(cfg.SynkTopic, nil)
	if err != nil {
		slog.Error("Unable to subscribe to synk topic")
		os.Exit(1)
	}

	slog.Info("✅Synk consumer listening...")

	return &SynkConsumer{
		consumer:  c,
		protoSynk: protoSynk,
	}
}

func (s *SynkConsumer) Consume() {
	for {
		msg, err := s.consumer.ReadMessage(-1)
		if err != nil {
			slog.Error("Unable to process kafka message", slog.String("error", err.Error()))
			continue
		}

		protoSynk := &pb.Synk{}
		if err := proto.Unmarshal(msg.Value, protoSynk); err != nil {
			slog.Error("Unable to process kafka message", slog.String("message", string(msg.Value)))
			continue
		}

		s.protoSynk <- protoSynk
	}

}
