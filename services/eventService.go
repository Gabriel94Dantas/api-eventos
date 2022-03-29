package services

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	confluentKafka "github.com/confluentinc/confluent-kafka-go/kafka"

	contex "facade-golang/api-eventos/contexts"
	model "facade-golang/api-eventos/models"
)

func InsertEventOnTopic(event model.Event) {
	model.Events = append(model.Events, event)
	produceMessage(event)
}

func CreateTopic(topic model.Topic) {
	createTopic(topic)
}

func produceMessage(event model.Event) {
	p := contex.KafkaProducerConnection()

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

func createTopic(topic model.Topic) {
	a := contex.KafkaAdminCliConnectio()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Create topics on cluster.
	// Set Admin options to wait for the operation to finish (or at most 60s)
	maxDur, err := time.ParseDuration("60s")
	if err != nil {
		panic("ParseDuration(60s)")
	}
	results, err := a.CreateTopics(
		ctx,
		// Multiple topics can be created simultaneously
		// by providing more TopicSpecification structs here.
		[]kafka.TopicSpecification{{
			Topic:             topic.TopicName,
			NumPartitions:     topic.NumPartition,
			ReplicationFactor: topic.ReplicationFactor}},
		// Admin options
		kafka.SetAdminOperationTimeout(maxDur))
	if err != nil {
		fmt.Printf("Failed to create topic: %v\n", err)
		os.Exit(1)
	}

	// Print results
	for _, result := range results {
		fmt.Printf("%s\n", result)
	}

	a.Close()
}
