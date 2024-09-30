package main

import (
	"github.com/devmegablaster/SheetBridge/internal/broker"
	"github.com/devmegablaster/SheetBridge/internal/config"
	"github.com/devmegablaster/SheetBridge/internal/database"
	"github.com/devmegablaster/SheetBridge/pb"
	"github.com/devmegablaster/SheetBridge/pkg/logger"
	"github.com/devmegablaster/SheetBridge/synkit/processor"
)

func main() {
	logger.Init()

	cfg := config.NewConfig()
	cfg.Init()

	dbSvc := database.New(&cfg.Database)

	synkCh := make(chan *pb.Synk)

	consumer := broker.NewSynkConsumer(&cfg.Kafka, synkCh)
	go consumer.Consume()

	handler := processor.NewSynkProcessor(dbSvc, cfg)

	for {
		select {
		case val := <-synkCh:
			handler.Handle(val)
		}
	}
}
