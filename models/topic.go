package models

import (
	"encoding/json"
)

type Topic struct {
	TopicName         string `json:"topicName"`
	NumPartition      int    `json:"numPartition"`
	ReplicationFactor int    `json:"replicationFactor"`
}

func TopicToJson(topic Topic) []byte {
	jsonTopic, err := json.Marshal(&topic)

	if err != nil {
		panic("Error Enconding EventToJson")
	}

	return jsonTopic
}
