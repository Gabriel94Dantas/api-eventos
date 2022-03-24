package services

import (
	"context"

	kafka "github.com/segmentio/kafka-go"

	model "facade-golang/api-eventos/models"
)

const (
	broker1Address = "localhost:9092"
)

func InsertEventOnTopic(event model.Event) {
	model.Events = append(model.Events, event)
	produceMessage(event, context.Background())
}

func produceMessage(event model.Event, context context.Context) {
	configurantion := kafka.WriterConfig{
		Brokers: []string{broker1Address},
		Topic:   event.Type,
	}

	message := kafka.Message{
		Key:   []byte(event.Id),
		Value: model.EventToJson(event),
	}

	writer := kafka.NewWriter(configurantion)

	error := writer.WriteMessages(context, message)
	if error != nil {
		panic("Error on produce event")
	}
}
