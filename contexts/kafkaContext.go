package contexts

import (
	"os"

	confluentKafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

func KafkaConnection() *confluentKafka.Producer {
	var p *confluentKafka.Producer
	var err error = nil

	if os.Getenv("BROKER_HOST") == "" {
		os.Setenv("BROKER_HOST", "localhost")
	}

	p, err = confluentKafka.NewProducer(&confluentKafka.ConfigMap{"bootstrap.servers": os.Getenv("BROKER_HOST")})

	if err != nil {
		panic(err)
	}

	return p
}
