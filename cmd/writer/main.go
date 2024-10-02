package main

import (
	"github.com/devmegablaster/SheetBridge/internal/broker"
	"github.com/devmegablaster/SheetBridge/internal/config"
	"github.com/devmegablaster/SheetBridge/pb"
	"github.com/devmegablaster/SheetBridge/pkg/logger"
	"github.com/devmegablaster/SheetBridge/writer"
)

func main() {
	logger.Init()

	cfg := config.NewConfig()
	cfg.Init()

	writeCh := make(chan *pb.Write)

	consumer := broker.NewWriteConsumer(&cfg.Kafka, writeCh)
	go consumer.Consume()

	processor := writer.NewWriteProcessor()

	for {
		select {
		case write := <-writeCh:
			processor.Handle(write)
		}
	}
}
