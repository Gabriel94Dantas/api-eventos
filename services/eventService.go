package services

import (
	"fmt"

	confluentKafka "github.com/confluentinc/confluent-kafka-go/kafka"

	contex "facade-golang/api-eventos/contexts"
	model "facade-golang/api-eventos/models"
)

func InsertEventOnTopic(event model.Event) {
	model.Events = append(model.Events, event)
	produceMessage(event)
}

func produceMessage(event model.Event) {
	p := contex.KafkaConnection()

	defer p.Close()

	// Delivery report handler for produced messages
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *confluentKafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	p.Produce(&confluentKafka.Message{
		TopicPartition: confluentKafka.TopicPartition{Topic: &event.Type, Partition: confluentKafka.PartitionAny},
		Key:            []byte(event.Id),
		Value:          model.EventToJson(event),
	}, nil)

	// Wait for message deliveries before shutting down
	p.Flush(15 * 1000)
}
